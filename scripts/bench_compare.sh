#!/usr/bin/env bash
# Cross-language benchmark comparison: fin-proto-*-bin-rs (examples/bench.rs),
# fin-proto-*-bin-zig (bench/bench.zig) and fin-proto-*-bin-c (bench/bench.c).
# All three implement the identical methodology — 10k warmup iterations, then
# 10 batches of 100k operations, fastest batch reported — so ns/op numbers are
# directly comparable.
#
# Toolchains: cargo, zig and a C compiler (cc) must be installed. Repositories
# are cloned shallow from xinchentechnote unless local checkouts are provided:
#   BENCH_RS_ROOT  dir containing fin-proto-{sse,szse,risk}-bin-rs
#   BENCH_ZIG_ROOT dir containing fin-proto-{sse,szse,risk}-bin-zig and
#                  fin-proto-runtime-bin-zig (the path dependency)
#   BENCH_C_ROOT   dir containing fin-proto-{sse,szse,risk}-bin-c and
#                  fin-proto-runtime-bin-c
#
# Raw TSV output lands in ./bench-results/, and a markdown comparison table
# is printed and appended to $GITHUB_STEP_SUMMARY when set.
#
# Usage: scripts/bench_compare.sh
set -uo pipefail

ROOT="$(cd "$(dirname "$0")/.." && pwd)"
RUNTIME_BASE="https://github.com/xinchentechnote"
RESULTS="$ROOT/bench-results"
LIBS=(sse szse risk)

WORK="$(mktemp -d)"
trap 'rm -rf "$WORK"' EXIT

note() { printf '%s\n' "$*"; }
die() { printf '[FAIL] %s\n' "$*" >&2; exit 1; }

ensure_repo() {
	local dest="$1" repo="$2"
	if [ -d "$dest/.git" ]; then
		note "using local checkout: $dest"
		return 0
	fi
	mkdir -p "$(dirname "$dest")"
	git clone -q --depth 1 "$RUNTIME_BASE/$repo.git" "$dest" 2>"$WORK/clone-$repo.log" ||
		die "could not clone $repo: $(tail -2 "$WORK/clone-$repo.log")"
}

RS_ROOT="${BENCH_RS_ROOT:-$WORK/rs}"
ZIG_ROOT="${BENCH_ZIG_ROOT:-$WORK/zig}"
C_ROOT="${BENCH_C_ROOT:-$WORK/c}"

command -v cargo >/dev/null 2>&1 || die "cargo not found (install the rust toolchain)"
command -v zig >/dev/null 2>&1 || die "zig not found (install zig 0.17+)"
command -v cc >/dev/null 2>&1 || die "cc not found (install a C toolchain)"

mkdir -p "$RESULTS"

note "== preparing repositories =="
ensure_repo "$ZIG_ROOT/fin-proto-runtime-bin-zig" fin-proto-runtime-bin-zig
ensure_repo "$C_ROOT/fin-proto-runtime-bin-c" fin-proto-runtime-bin-c
for lib in "${LIBS[@]}"; do
	ensure_repo "$RS_ROOT/fin-proto-$lib-bin-rs" "fin-proto-$lib-bin-rs"
	ensure_repo "$ZIG_ROOT/fin-proto-$lib-bin-zig" "fin-proto-$lib-bin-zig"
	ensure_repo "$C_ROOT/fin-proto-$lib-bin-c" "fin-proto-$lib-bin-c"
done

for lib in "${LIBS[@]}"; do
	note "== benchmarking $lib =="
	(cd "$RS_ROOT/fin-proto-$lib-bin-rs" && cargo run --release -q --example bench) \
		>"$RESULTS/$lib.rust.tsv" || die "rust bench failed for $lib"

	# the generated code resolves `binary_codec` through a path dependency on
	# the runtime checkout beside it; the bench binary prints its TSV on stderr
	(cd "$ZIG_ROOT/fin-proto-$lib-bin-zig" \
		&& zig build -Doptimize=ReleaseFast >/dev/null 2>"$WORK/zig-build-$lib.log" \
		&& ./zig-out/bin/bench 2>"$RESULTS/$lib.zig.tsv") ||
		{ tail -5 "$WORK/zig-build-$lib.log"; die "zig bench failed for $lib"; }

	# make -s regenerates src/ from the DSL, compiles with -O2 and runs the
	# bench binary, whose stdout carries only the two TSV lines
	(cd "$C_ROOT/fin-proto-$lib-bin-c" && make -s bench) \
		>"$RESULTS/$lib.c.tsv" || die "c bench failed for $lib"
done

note ""
note "== raw results (ns/op, fastest batch) =="
cat "$RESULTS"/*.tsv

emit_table() {
	echo "### Codec benchmark: Rust (fin-proto-*-bin-rs) vs Zig (fin-proto-*-bin-zig) vs C (fin-proto-*-bin-c)"
	echo ""
	echo "Methodology on all sides: 10k warmup iterations, 10 batches of 100k operations, fastest batch reported as ns/op. GitHub-hosted runners are shared VMs — treat differences under ~20% as noise."
	echo ""
	echo "| Library | Case | Frame bytes | Rust ns/op | Zig ns/op | C ns/op | Zig vs Rust | C vs Rust |"
	echo "|---------|------|------------:|-----------:|----------:|--------:|------------:|----------:|"
	awk -F'\t' '
		{
			lib = $1; lang = $2; op = $3; ns = $4 + 0; bytes = $5 + 0;
			k = lib SUBSEP op;
			if (!(k in idx)) { idx[k] = ++n; ordlib[n] = lib; ordop[n] = op }
			val[k, lang] = ns; byt[k] = bytes;
		}
		function fmt(v) { return v > 0 ? sprintf("%.1f", v) : "—" }
		function ratio(a, b) { return (a > 0 && b > 0) ? sprintf("%.2fx", a / b) : "—" }
		END {
			for (i = 1; i <= n; i++) {
				lib = ordlib[i]; op = ordop[i]; k = lib SUBSEP op;
				printf "| %s | %s | %d | %s | %s | %s | %s | %s |\n",
					lib, op, byt[k],
					fmt(val[k, "rust"]), fmt(val[k, "zig"]), fmt(val[k, "c"]),
					ratio(val[k, "rust"], val[k, "zig"]), ratio(val[k, "rust"], val[k, "c"]);
			}
		}' "$RESULTS"/*.tsv
	echo ""
	echo "Zig and C decode are zero-copy (strings borrow from the input buffer); the Rust runtime materializes owned \`String\`s, which dominates its decode cost."
}

echo ""
emit_table
if [ -n "${GITHUB_STEP_SUMMARY:-}" ]; then
	emit_table >>"$GITHUB_STEP_SUMMARY"
fi
