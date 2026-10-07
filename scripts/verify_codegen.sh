#!/usr/bin/env bash
# Cross-language verification of fin-protoc generated code.
#
# Generates code in all six target languages from the full-grammar fixture
# (internal/parser/testdata/grammar_full.dsl), then compiles it against the
# matching fin-proto-* runtime and runs the generated round-trip tests.
#
# Each language is verified when its toolchain is available and skipped with
# a notice otherwise, so the script is usable on developer machines; CI
# installs every toolchain and therefore enforces all six.
#
# Usage: scripts/verify_codegen.sh [path-to-fin-protoc-binary]
set -uo pipefail

ROOT="$(cd "$(dirname "$0")/.." && pwd)"
BIN="${1:-$ROOT/bin/fin-protoc}"
DSL="$ROOT/internal/parser/testdata/grammar_full.dsl"
RUNTIME_DIR="${FIN_RUNTIME_DIR:-$ROOT/.verify-runtime}"
RUNTIME_BASE="https://github.com/xinchentechnote"
# STRICT=1 (set by CI) turns SKIP into failure so a missing toolchain or a
# failed runtime clone can never silently green the pipeline.
STRICT="${STRICT:-0}"

FAILED=0
declare -a RESULTS

summary_file="${GITHUB_STEP_SUMMARY:-}"

note() { printf '%s\n' "$*"; }
ok() { RESULTS+=("PASS  $1"); printf '[PASS] %s\n' "$1"; }
fail() { RESULTS+=("FAIL  $1"); printf '[FAIL] %s\n' "$1"; FAILED=$((FAILED + 1)); }
skip() {
	RESULTS+=("SKIP  $1"); printf '[SKIP] %s\n' "$1"
	if [ "$STRICT" = 1 ]; then FAILED=$((FAILED + 1)); fi
}

ensure_runtime() {
	local repo="$1"
	local dest="$RUNTIME_DIR/$repo"
	[ -d "$dest/.git" ] && return 0
	mkdir -p "$RUNTIME_DIR"
	# submodules may point at SSH URLs a CI runner cannot reach; retry
	# without them before giving up
	if ! git clone -q --depth 1 --recurse-submodules --shallow-submodules \
		"$RUNTIME_BASE/$repo.git" "$dest" 2>"$WORK/clone-$repo.log"; then
		rm -rf "$dest"
		if ! git clone -q --depth 1 "$RUNTIME_BASE/$repo.git" "$dest" \
			2>>"$WORK/clone-$repo.log"; then
			note "clone $repo failed:" && tail -3 "$WORK/clone-$repo.log"
			return 1
		fi
	fi
}

WORK="$(mktemp -d)"
trap 'rm -rf "$WORK"' EXIT

if [ ! -x "$BIN" ]; then
	fail "fin-protoc binary not found at $BIN (run make main-build first)"
	print_summary() { :; }
	exit 1
fi

note "== generating all targets from $DSL =="
"$BIN" compile -f "$DSL" \
	-g "$WORK/go" -r "$WORK/rust" -j "$WORK/java" \
	-p "$WORK/python" -c "$WORK/cpp" -l "$WORK/lua" || {
	fail "code generation"
	exit 1
}

# --- Go: build + vet + run generated round-trip tests -----------------------
if command -v go >/dev/null 2>&1; then
	if ensure_runtime fin-proto-go; then
		mkdir -p "$WORK/gomod"
		cp "$WORK/go/"*.go "$WORK/gomod/"
		cat > "$WORK/gomod/go.mod" <<EOF
module fin-protoc-codegen-verify

go 1.24

require (
	github.com/stretchr/testify v1.11.1
	github.com/xinchentechnote/fin-proto-go v0.0.0-00010101000000-000000000000
)

replace github.com/xinchentechnote/fin-proto-go => "$RUNTIME_DIR/fin-proto-go"
EOF
		if (cd "$WORK/gomod" \
			&& { go mod tidy || { sleep 2; go mod tidy; } } \
			&& go build ./... && go vet ./... && go test ./...) \
			>"$WORK/go-verify.log" 2>&1; then
			ok "Go: build + vet + generated tests"
		else
			tail -15 "$WORK/go-verify.log"
			fail "Go: build/vet/test generated code"
		fi
	else
		skip "Go: could not clone fin-proto-go"
	fi
else
	skip "Go: toolchain not found"
fi

# --- Rust: build + run generated tests inside fin-proto-rs workspace --------
if command -v cargo >/dev/null 2>&1; then
	if ensure_runtime fin-proto-rs; then
		rm -f "$RUNTIME_DIR/fin-proto-rs/sample-binary/src/"*.rs
		cp "$WORK/rust/"*.rs "$RUNTIME_DIR/fin-proto-rs/sample-binary/src/"
		if (cd "$RUNTIME_DIR/fin-proto-rs" && cargo test -q -p sample-binary >/dev/null 2>&1); then
			ok "Rust: cargo build + generated tests"
		else
			fail "Rust: cargo build/test generated code"
		fi
	else
		skip "Rust: could not clone fin-proto-rs"
	fi
else
	skip "Rust: toolchain not found"
fi

# --- Java: compile + run generated JUnit tests in fin-proto-java ------------
JAVA_OK=0
if command -v java >/dev/null 2>&1; then
	# the sample-bin module declares a Java 11 toolchain
	if /usr/libexec/java_home -v 11 >/dev/null 2>&1 \
		|| ls /usr/lib/jvm 2>/dev/null | grep -q 11; then
		JAVA_OK=1
	elif [ -n "${JAVA_HOME:-}" ]; then
		JAVA_OK=1 # CI provides a matching JDK; accept the default one
	fi
fi
if [ "$JAVA_OK" = 1 ] && command -v gradle >/dev/null 2>&1; then
	if ensure_runtime fin-proto-java; then
		SRC="$RUNTIME_DIR/fin-proto-java/sample-bin/src"
		PKG_DIR="com/finproto/sample/messages"
		rm -f "$SRC/main/java/$PKG_DIR/"*.java "$SRC/test/java/$PKG_DIR/"*Test.java
		cp "$WORK/java/main/java/com/grammar/full/messages/"*.java "$SRC/main/java/$PKG_DIR/"
		cp "$WORK/java/test/java/com/grammar/full/messages/"*Test.java "$SRC/test/java/$PKG_DIR/"
		if (cd "$RUNTIME_DIR/fin-proto-java" \
			&& ./gradlew -q --no-daemon :sample-bin:test >/dev/null 2>&1); then
			ok "Java: gradle build + generated tests"
		else
			fail "Java: gradle build/test generated code"
		fi
	else
		skip "Java: could not clone fin-proto-java"
	fi
else
	skip "Java: toolchain (JDK 11 + gradle) not found"
fi

# --- C++: compile + run generated gtest suite (runtime is Linux-only) -------
LUAC_BIN="$(command -v luac || command -v luac5.4 || command -v luac5.3 || true)"
GTEST_OK=0
if [ "$(uname -s)" = "Linux" ] && command -v g++ >/dev/null 2>&1; then
	if echo '#include <gtest/gtest.h>' | g++ -std=c++17 -x c++ - -fsyntax-only >/dev/null 2>&1; then
		GTEST_OK=1
	fi
fi
if [ "$GTEST_OK" = 1 ]; then
	if ensure_runtime fin-proto-cpp; then
		cp "$WORK/cpp/include/grammar_full.hpp" "$RUNTIME_DIR/fin-proto-cpp/include/"
		if g++ -std=c++17 -I"$RUNTIME_DIR/fin-proto-cpp" \
			"$WORK/cpp/test/grammar_full_test.cpp" \
			-lgtest -lgtest_main -lz -pthread -o "$WORK/cpp_test" \
			&& "$WORK/cpp_test"; then
			ok "C++: g++ build + generated gtest suite"
		else
			fail "C++: compile/run generated code"
		fi
	else
		skip "C++: could not clone fin-proto-cpp"
	fi
else
	skip "C++: Linux + g++ + libgtest required (runtime uses <endian.h>)"
fi

# --- Python: syntax check, then run generated unittest against fin-proto-py -
if command -v python3 >/dev/null 2>&1; then
	if python3 -m py_compile "$WORK/python/"*.py 2>/dev/null; then
		# the fin-proto-py runtime needs numpy; prefer a reusable venv, fall
		# back to a throwaway one, otherwise only the syntax check remains
		PYBIN=""
		if [ -x "$ROOT/.verify-venv/bin/python" ]; then
			PYBIN="$ROOT/.verify-venv/bin/python"
		elif python3 -m venv "$WORK/venv" 2>/dev/null && "$WORK/venv/bin/pip" install -q numpy >/dev/null 2>&1; then
			PYBIN="$WORK/venv/bin/python"
		fi
		if [ -n "$PYBIN" ] && ensure_runtime fin-proto-py; then
			if (cd "$WORK/python" \
				&& PYTHONPATH="$RUNTIME_DIR/fin-proto-py/lib:$WORK/python" \
				"$PYBIN" -m unittest "grammar_full_test" >/dev/null 2>&1); then
				ok "Python: syntax + generated unittest"
			else
				fail "Python: generated unittest failed"
			fi
		else
			skip "Python: unittest run (venv/numpy or fin-proto-py unavailable)"
		fi
	else
		fail "Python: generated code has syntax errors"
	fi
else
	skip "Python: toolchain not found"
fi

# --- Lua: Wireshark dissectors are plain Lua; syntax check only -------------
if [ -n "$LUAC_BIN" ]; then
	LUA_FAIL=0
	for f in "$WORK/lua/"*.lua; do
		"$LUAC_BIN" -p "$f" || LUA_FAIL=1
	done
	if [ "$LUA_FAIL" = 0 ]; then
		ok "Lua: luac syntax check"
	else
		fail "Lua: generated dissector has syntax errors"
	fi
else
	skip "Lua: luac not found"
fi

note ""
note "== verification summary =="
for r in "${RESULTS[@]}"; do note "  $r"; done

# publish the summary on the GitHub run page (visible without log access)
if [ -n "$summary_file" ]; then
	{
		echo "### Generated-code verification (grammar_full.dsl)"
		echo ""
		echo "| Result | Check |"
		echo "|--------|-------|"
		for r in "${RESULTS[@]}"; do
			echo "| ${r%% *} | ${r#*  } |"
		done
	} >>"$summary_file"
fi

if [ "$FAILED" -gt 0 ]; then
	note "$FAILED language(s) FAILED"
	exit 1
fi
note "all available toolchains verified"
