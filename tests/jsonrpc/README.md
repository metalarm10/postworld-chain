# JSON-RPC Compatibility Testing

This directory contains tools and scripts for testing JSON-RPC API compatibility between Cosmos EVM and Ethereum clients.

## Quick Start

```bash
# From project root
make test-rpc-compat
```

## Test Guide

### 1. Build EVMD Docker Image

```bash
# From project root
make localnet-build-env
```

### 2. Start Nodes

```bash
# Start postworldd with JSON-RPC enabled
./tests/jsonrpc/scripts/postworldd/start-postworldd.sh

# Start geth for comparison
./tests/jsonrpc/scripts/geth/start-geth.sh

# Or start both at once
./tests/jsonrpc/scripts/start-networks.sh
```

### 3. Run Compatibility Tests

```bash
# Use the simulator for comprehensive testing
cd tests/jsonrpc/simulator
go build .
./simulator
```

### 4. Stop Nodes

```bash
# Stop postworldd
./tests/jsonrpc/scripts/postworldd/stop-postworldd.sh

# Stop geth
./tests/jsonrpc/scripts/geth/stop-geth.sh

# Or stop both at once
./tests/jsonrpc/scripts/stop-networks.sh
```

## Available Endpoints

### postworldd Endpoints

- **JSON-RPC**: http://localhost:8545
- **WebSocket**: ws://localhost:8546  
- **Cosmos REST**: http://localhost:1317
- **Tendermint RPC**: http://localhost:26657
- **gRPC**: localhost:9090

### geth Endpoints

- **JSON-RPC**: http://localhost:8547
- **WebSocket**: ws://localhost:8548

## Scripts Structure

### `scripts/postworldd/`

- `start-postworldd.sh` - Initialize and start single-node postworldd for testing
- `stop-postworldd.sh` - Stop the postworldd testing node

### `scripts/geth/`

- `start-geth.sh` - Start geth node using ethereum/client-go:v1.15.10
- `stop-geth.sh` - Stop the geth testing node

### `scripts/`

- `start-both.sh` - Start both postworldd and geth nodes
- `stop-both.sh` - Stop both nodes

## Testing with Simulator

The simulator in `./simulator/` is the primary tool for comprehensive compatibility testing:

```bash
cd tests/jsonrpc/simulator
go build .
./simulator
```

## Configuration

The scripts use the following defaults:

### postworldd Configuration

- Container name: `postworldd-jsonrpc-test`
- Chain ID: `local-4221`
- Validator count: 1
- Data directory: `tests/jsonrpc/.postworldd`

### geth Configuration

- Container name: `geth-jsonrpc-test`
- Chain ID: 1337 (dev mode)
- Data directory: `tests/jsonrpc/.geth-data`

## Troubleshooting

### Container fails to start

- Check if the Docker image was built: `docker images | grep cosmos/postworldd`
- Check container logs: `docker logs postworldd-jsonrpc-test`

### JSON-RPC not responding

- Verify the container is running: `docker ps | grep postworldd-jsonrpc-test`
- Check if ports are bound: `docker port postworldd-jsonrpc-test`
- Test with curl: `curl -X POST -H "Content-Type: application/json" --data '{"jsonrpc":"2.0","method":"eth_chainId","params":[],"id":1}' http://localhost:8545`
