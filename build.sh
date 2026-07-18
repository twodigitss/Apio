#!/bin/bash

# Exit on error
set -e

# Source and target paths
SRC_FILE="cmd/tea/main.go"
OUT_BIN="cmd/tea/main"

echo "=== Building apio ==="

# Verify source file exists
if [ ! -f "$SRC_FILE" ]; then
    echo "Error: Source file $SRC_FILE not found."
    exit 1
fi

# Compile stripped binary (removes debug info and symbols)
echo "Compiling optimized binary..."
go build -ldflags="-s -w" -o "$OUT_BIN" "$SRC_FILE"

# Measure initial size
SIZE_STRIPPED=$(du -h "$OUT_BIN" | cut -f1)
echo "✓ Compiled size (stripped): $SIZE_STRIPPED"

# Check for upx availability
# if command -v upx >/dev/null 2>&1; then
#     echo "UPX found! Compressing binary..."
#     if [ "$(uname)" = "Darwin" ]; then
#         upx --force-macos "$OUT_BIN"
#         echo "Re-signing binary..."
#         codesign -f -s - "$OUT_BIN"
#     else
#         upx "$OUT_BIN"
#     fi
#     SIZE_COMPRESSED=$(du -h "$OUT_BIN" | cut -f1)
#     echo "✓ Compressed size: $SIZE_COMPRESSED"
# else
#     echo "ℹ UPX is not installed."
#     echo "  You can compress the binary even more (~60-70% extra reduction) by installing UPX."
#     echo "  On macOS, you can install it running: brew install upx"
# fi

# Install if 'install' or '--install' argument is provided
if [ "$1" = "install" ] || [ "$1" = "--install" ]; then
    echo "Installing binary to ~/.local/bin/apio..."
    mkdir -p "$HOME/.local/bin"
    cp "$OUT_BIN" "$HOME/.local/bin/apio"
    echo "✓ Installed successfully! Make sure ~/.local/bin is in your PATH."
fi

echo "=== Build finished successfully! ==="
echo "Binary location: $OUT_BIN"
