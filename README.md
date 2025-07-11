# ErcNftMetadataIndexerNetwork

## Description

A curated collection of Solidity smart contracts implementing ERC-721A and ERC-1155 NFT standards with gas-efficient minting via merkle tree proofs and on-chain metadata compression using optimized string packing.

## Features

- Indexes ERC-721 and ERC-1155 metadata from multiple EVM-compatible chains, including Ethereum, Polygon, and Binance Smart Chain.
- Stores NFT metadata in a normalized, relational database schema optimized for efficient querying and filtering.
- Implements a robust event listener that subscribes to transfer events and token URI updates on monitored smart contracts.
- Utilizes a rate-limiting mechanism to prevent overloading RPC endpoints and ensure data consistency.
- Provides a GraphQL API for querying NFT metadata, including token IDs, owner addresses, and associated media URLs.
- Supports customizable metadata refresh policies, allowing users to configure how frequently the indexer checks for updates to token URIs.
- Includes a caching layer to minimize database reads and improve API response times for frequently accessed NFT metadata.
- Deploys a distributed architecture with horizontally scalable components for handling high volumes of NFT transactions.
## Installation

```bash
pip install ercnftmetadataindexernetwork
```

## Usage

```python
from ercnftmetadataindexernetwork import ErcNftMetadataIndexerNetwork

# Initialize
app = ErcNftMetadataIndexerNetwork()

# Run
app.run()
```

## Contributing

We welcome contributions! Here's how to get started:

1. Fork this repository
2. Create a new branch for your feature (`git checkout -b feature/your-feature`)
3. Commit your changes (`git commit -am 'Add some awesome feature'`)
4. Push to the branch (`git push origin feature/your-feature`)
5. Open a Pull Request

## License

Distributed under the MIT License. See `LICENSE` for more information.
