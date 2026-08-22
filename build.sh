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

# Install if 'install' or '--install' argument is provided
if [ "$1" = "install" ] || [ "$1" = "--install" ]; then
    echo "Installing binary to ~/.local/bin/apio..."
    mkdir -p "$HOME/.local/bin"
    cp "$OUT_BIN" "$HOME/.local/bin/apio"
    echo "✓ Installed successfully! Make sure ~/.local/bin is in your PATH."

    CONFIG_DIR="$HOME/.config/apio"
    CONFIG_FILE="$CONFIG_DIR/config.toml"
    mkdir -p "$CONFIG_DIR"
    if [ ! -f "$CONFIG_FILE" ]; then
        cp "cmd/tea/config.toml" "$CONFIG_FILE"
        echo "✓ Config installed to $CONFIG_FILE"
    else
        echo "  Config already exists at $CONFIG_FILE, skipping."
    fi
fi

echo "=== Build finished successfully! ==="
echo "Binary location: $OUT_BIN"
