#!/usr/bin/env bash
# Cross-language benchmark comparison: fin-proto-*-bin-rs (criterion-style
# hand-rolled timing via examples/bench.rs) vs fin-proto-*-bin-zig
# (bench/bench.zig). Both sides implement the identical methodology — 10k
# warmup iterations, then 10 batches of 100k operations, fastest batch
# reported — so ns/op numbers are directly comparable.
#
# Per-language toolchains must be installed (cargo, zig). Repositories are
# cloned shallow from xinchentechnote unless local checkouts are provided:
#   BENCH_RS_ROOT  dir containing fin-proto-{sse,szse,risk}-bin-rs
#   BENCH_ZIG_ROOT dir containing fin-proto-{sse,szse,risk}-bin-zig and
#                  fin-proto-runtime-bin-zig (the path dependency)
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

command -v cargo >/dev/null 2>&1 || die "cargo not found (install the rust toolchain)"
command -v zig >/dev/null 2>&1 || die "zig not found (install zig 0.17+)"

mkdir -p "$RESULTS"

note "== preparing repositories =="
ensure_repo "$ZIG_ROOT/fin-proto-runtime-bin-zig" fin-proto-runtime-bin-zig
for lib in "${LIBS[@]}"; do
	ensure_repo "$RS_ROOT/fin-proto-$lib-bin-rs" "fin-proto-$lib-bin-rs"
	ensure_repo "$ZIG_ROOT/fin-proto-$lib-bin-zig" "fin-proto-$lib-bin-zig"
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
done

note ""
note "== raw results (ns/op, fastest of $((10)) batches) =="
cat "$RESULTS"/*.tsv

emit_table() {
	echo "### Codec benchmark: Rust (fin-proto-*-bin-rs) vs Zig (fin-proto-*-bin-zig)"
	echo ""
	echo "Methodology on both sides: 10k warmup iterations, 10 batches of 100k operations, fastest batch reported as ns/op. GitHub-hosted runners are shared VMs — treat differences under ~20% as noise."
	echo ""
	echo "| Library | Case | Frame bytes | Rust ns/op | Zig ns/op | Zig vs Rust |"
	echo "|---------|------|------------:|-----------:|----------:|------------:|"
	for lib in "${LIBS[@]}"; do
		r="$RESULTS/$lib.rust.tsv"
		z="$RESULTS/$lib.zig.tsv"
		if [ ! -s "$r" ] || [ ! -s "$z" ]; then
			echo "| $lib | — | — | unavailable | unavailable | — |"
			continue
		fi
		paste "$r" "$z" | awk -F'\t' '{
			rns = $4; zns = $9; op = $3; bytes = $5;
			if (zns + 0 > 0) printf "| %s | %s | %d | %.1f | %.1f | %.2fx |\n", $1, op, bytes, rns, zns, rns / zns;
			else printf "| %s | %s | %d | %.1f | %.1f | — |\n", $1, op, bytes, rns, zns;
		}'
	done
	echo ""
	echo "Zig decode is zero-copy (strings borrow from the input buffer); the Rust runtime materializes owned \`String\`s, which dominates its decode cost."
}

echo ""
emit_table
if [ -n "${GITHUB_STEP_SUMMARY:-}" ]; then
	emit_table >>"$GITHUB_STEP_SUMMARY"
fi
