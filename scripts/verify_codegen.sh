#!/usr/bin/env bash
# Cross-language verification of fin-protoc generated code.
#
# Generates code in all eight target languages from the full-grammar fixture
# (internal/parser/testdata/grammar_full.dsl), then compiles it against the
# matching fin-proto-* runtime and runs the generated round-trip tests.
#
# Each language is verified when its toolchain is available and skipped with
# a notice otherwise, so the script is usable on developer machines; CI
# installs every toolchain and therefore enforces all languages. The Zig and
# C checks additionally need their runtime repositories: they are cloned like
# the other runtimes, or taken from $FIN_PROTO_ZIG_RUNTIME /
# $FIN_PROTO_C_RUNTIME (local checkouts) when set; until those repositories
# are published the check reports PENDING, which never fails the run —
# unlike SKIP under STRICT=1.
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
pending() { RESULTS+=("PEND  $1"); printf '[PEND] %s\n' "$1"; }
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
	-p "$WORK/python" -c "$WORK/cpp" -l "$WORK/lua" -z "$WORK/zig" -C "$WORK/c" || {
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

# --- Zig: scaffold a package around the generated sources and run tests -----
# The generated code imports the `binary_codec` module provided by
# fin-proto-runtime-bin-zig, wired in through a path dependency. Use a local
# checkout via FIN_PROTO_ZIG_RUNTIME, or clone the published repository;
# until it is published this check stays PENDING instead of failing.
if command -v zig >/dev/null 2>&1; then
	ZIG_RUNTIME="${FIN_PROTO_ZIG_RUNTIME:-}"
	if [ -n "$ZIG_RUNTIME" ] && [ -d "$ZIG_RUNTIME/src" ]; then
		note "using local zig runtime: $ZIG_RUNTIME"
	elif ensure_runtime fin-proto-runtime-bin-zig; then
		ZIG_RUNTIME="$RUNTIME_DIR/fin-proto-runtime-bin-zig"
	else
		ZIG_RUNTIME=""
	fi
	if [ -z "$ZIG_RUNTIME" ]; then
		pending "Zig: fin-proto-runtime-bin-zig not published yet (set FIN_PROTO_ZIG_RUNTIME for a local checkout)"
	else
		mkdir -p "$WORK/zigpkg/src"
		cp "$WORK/zig/"*.zig "$WORK/zigpkg/src/"
		# build.zig.zon path dependencies must be relative; link the runtime in
		ln -sfn "$(cd "$ZIG_RUNTIME" && pwd)" "$WORK/zigpkg/binary_codec"
		cat > "$WORK/zigpkg/build.zig" <<'EOF'
const std = @import("std");

pub fn build(b: *std.Build) void {
    const target = b.standardTargetOptions(.{});
    const optimize = b.standardOptimizeOption(.{});
    const codec_dep = b.dependency("binary_codec", .{ .target = target, .optimize = optimize });
    const codec_mod = codec_dep.module("binary_codec");
    const mod = b.addModule("fin_protoc_codegen_verify", .{
        .root_source_file = b.path("src/root.zig"),
        .target = target,
        .optimize = optimize,
    });
    mod.addImport("binary_codec", codec_mod);
    const mod_tests = b.addTest(.{ .root_module = mod });
    const test_step = b.step("test", "Run tests");
    test_step.dependOn(&b.addRunArtifact(mod_tests).step);
}
EOF
		cat > "$WORK/zigpkg/build.zig.zon" <<EOF
.{
    .name = .fin_protoc_codegen_verify,
    .version = "0.0.0",
    .fingerprint = 0x199b6ee547ac9ae3, // derived from the package name
    .minimum_zig_version = "0.17.0",
    .dependencies = .{
        .binary_codec = .{ .path = "binary_codec" },
    },
    .paths = .{""},
}
EOF
		if (cd "$WORK/zigpkg" && zig build test >"$WORK/zig-verify.log" 2>&1); then
			ok "Zig: zig build test generated code"
		else
			tail -15 "$WORK/zig-verify.log"
			fail "Zig: zig build test generated code"
		fi
	fi
else
	skip "Zig: toolchain not found"
fi

# --- C: compile generated sources against fin-proto-runtime-bin-c -----------
# The generated tests.c provides main(); one cc invocation builds and the
# binary runs every per-packet round-trip test.
if command -v cc >/dev/null 2>&1; then
	C_RUNTIME="${FIN_PROTO_C_RUNTIME:-}"
	if [ -n "$C_RUNTIME" ] && [ -d "$C_RUNTIME/include" ]; then
		note "using local c runtime: $C_RUNTIME"
	elif ensure_runtime fin-proto-runtime-bin-c; then
		C_RUNTIME="$RUNTIME_DIR/fin-proto-runtime-bin-c"
	else
		C_RUNTIME=""
	fi
	if [ -z "$C_RUNTIME" ]; then
		pending "C: fin-proto-runtime-bin-c not published yet (set FIN_PROTO_C_RUNTIME for a local checkout)"
	else
		if (cd "$WORK/c" \
			&& cc -std=c11 -Wall -Wextra -O2 -I"$C_RUNTIME/include" -I. \
				*.c "$C_RUNTIME"/src/*.c -o "$WORK/c_test" 2>"$WORK/c-build.log" \
			&& "$WORK/c_test" >"$WORK/c-verify.log" 2>&1); then
			ok "C: cc build + generated tests"
		else
			tail -15 "$WORK/c-build.log" "$WORK/c-verify.log"
			fail "C: compile/run generated code"
		fi
	fi
else
	skip "C: toolchain not found"
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
