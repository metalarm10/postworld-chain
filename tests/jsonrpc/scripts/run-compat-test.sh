#!/bin/bash

# JSON-RPC Compatibility Test Runner with Docker Image Optimization
# This script handles Docker image building with content-based caching

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(cd "$SCRIPT_DIR/../../.." && pwd)"
JSONRPC_DIR="$PROJECT_ROOT/tests/jsonrpc"

echo "🔍 Checking Docker image requirements..."

# Check postworldd image and build if needed
if ! docker image inspect cosmos/postworldd >/dev/null 2>&1; then
    echo "📦 Building cosmos/postworldd image..."
    make -C "$PROJECT_ROOT" localnet-build-env
else
    echo "✓ cosmos/postworldd image already exists, skipping build"
fi

# Check if simulator image already exists
if docker image inspect jsonrpc_simulator >/dev/null 2>&1; then
    echo "✓ Simulator image already exists"
else
    echo "📦 Will build simulator image..."
fi

# Initialize postworldd data directory
echo "🔧 Preparing postworldd data directory..."

# Clear existing directory to avoid key conflicts
if [ -d "$JSONRPC_DIR/.postworldd" ]; then
    echo "🧹 Removing existing .postworldd directory..."
    rm -rf "$JSONRPC_DIR/.postworldd"
fi

# Create fresh directory with correct permissions  
mkdir -p "$JSONRPC_DIR/.postworldd"
chmod 777 "$JSONRPC_DIR/.postworldd"

echo "🔧 postworldd will auto-initialize when container starts..."

# Run the compatibility tests
echo "🚀 Running JSON-RPC compatibility tests..."
cd "$JSONRPC_DIR" && docker compose up --build --abort-on-container-exit


echo "✅ JSON-RPC compatibility test completed!"
