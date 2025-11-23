#!/bin/bash
set -euo pipefail

VERSION="${1:-}"
PLATFORM="${2:-}"

if [ -z "$VERSION" ] || [ -z "$PLATFORM" ]; then
  echo "Error: Missing required arguments"
  echo "Usage: $0 <version> <platform>"
  exit 1
fi

ROOT_DIR="$(git rev-parse --show-toplevel)"
TMP_DIR="$ROOT_DIR/.tmp"
YOGA_DIR="$TMP_DIR/yoga"
BUILD_DIR="$YOGA_DIR/build"

clone_yoga() {
  mkdir -p "$TMP_DIR"
  rm -rf "$YOGA_DIR"

  git clone https://github.com/facebook/yoga.git "$YOGA_DIR"
  cd "$YOGA_DIR"

  if git rev-parse --verify "$VERSION" >/dev/null 2>&1; then
    git checkout "$VERSION"
  else
    echo "Error: version '$VERSION' not found."
    return 1
  fi
}

build_yoga() {
  cd "$YOGA_DIR"

  if [[ "$PLATFORM" == windows_* ]]; then
    cmake -S "$YOGA_DIR" -B "$BUILD_DIR" \
      -G "MinGW Makefiles" \
      -DCMAKE_BUILD_TYPE=Release \
      -DCMAKE_POSITION_INDEPENDENT_CODE=ON \
      -DCMAKE_C_COMPILER=gcc \
      -DCMAKE_CXX_COMPILER=g++
  else
    cmake -S "$YOGA_DIR" -B "$BUILD_DIR" \
      -DCMAKE_BUILD_TYPE=Release \
      -DCMAKE_POSITION_INDEPENDENT_CODE=ON
  fi

  cmake --build "$BUILD_DIR" --target yogacore --config Release -j "$(nproc 2>/dev/null || echo 2)"
}

install_artifacts() {
  local lib_dir="$ROOT_DIR/etc/lib/$PLATFORM"
  local include_dir="$ROOT_DIR/etc/include"

  mkdir -p "$lib_dir"
  cp "$BUILD_DIR/yoga/libyogacore.a" "$lib_dir/"

  rm -rf "$include_dir/yoga"
  mkdir -p "$include_dir"
  cp -r "$YOGA_DIR/yoga" "$include_dir/"
}

cleanup() {
  rm -rf "$TMP_DIR"
}

set_outputs() {
  if [ -n "${GITHUB_OUTPUT:-}" ]; then
    echo "lib-path=etc/lib/$PLATFORM" >> "$GITHUB_OUTPUT"
    echo "headers-path=etc/include" >> "$GITHUB_OUTPUT"
  fi
}

main() {
  echo "Building Yoga $VERSION for $PLATFORM"

  clone_yoga
  build_yoga
  install_artifacts
  cleanup
  set_outputs

  echo "Build successful"
}

main
