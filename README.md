# Postworld Chain

**Postworld Chain** is an EVM-compatible Cosmos SDK blockchain built on [Cosmos EVM](https://github.com/cosmos/evm).

## Overview

Postworld Chain combines the power of Ethereum's EVM with Cosmos SDK's modularity and IBC interoperability.

- **Chain ID**: `postworld_9000-1`
- **Binary**: `postworldd`
- **Bech32 Prefix**: `postworld`
- **Native Denom**: `asurvive` (1 SURVIVE = 10^18 asurvive)

## Features

- Full EVM compatibility (Solidity smart contracts, JSON-RPC, Web3 support)
- Ethereum tooling support (MetaMask, Remix, Hardhat, etc.)
- IBC connectivity for cross-chain asset transfers
- Native ERC-20 token support
- Customizable precompiles and extensions
- EIP-1559 fee market mechanism

## Getting Started

### Running a Local Node

To run a local development node:

```bash
./local_node.sh
```

This will:
- Build and install the `postworldd` binary
- Initialize a local testnet
- Start the chain with JSON-RPC enabled on `localhost:8545`

### Connecting with MetaMask

- **Network Name**: Postworld Chain
- **RPC URL**: `http://localhost:8545`
- **Chain ID**: `9000` (decimal)
- **Currency Symbol**: `SURVIVE`

## Development

### Building

```bash
make install
```

### Testing

```bash
# Unit tests
make test-unit

# Integration tests
make test-integration

# Solidity contract tests
make test-solidity
```

## Built on Cosmos EVM

This chain is built using [Cosmos EVM](https://github.com/cosmos/evm), a plug-and-play solution for adding EVM compatibility to Cosmos SDK chains.

### Credits

- **Cosmos EVM** - Base implementation by [Cosmos Labs](https://cosmoslabs.io/)
- **evmOS** - Original work by [Tharsis](https://github.com/evmos/OS)

## License

Apache 2.0

## Support

For issues and feature requests, please use the [GitHub issue tracker](https://github.com/metalarm10/postworld-chain/issues).
