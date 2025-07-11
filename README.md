# ErcNftMetadataIndexerNetwork

## Description



## Features

- Indexes ERC-721 and ERC-1155 token metadata from multiple EVM-compatible chains using a configurable RPC endpoint list.
- Stores indexed metadata in a PostgreSQL database with optimized schema for efficient querying and filtering.
- Exposes a GraphQL API for retrieving NFT metadata, including ownership, attributes, and associated media URLs.
- Implements a caching layer using Redis to minimize database load and improve API response times.
- Supports event-based indexing by subscribing to transfer events on target smart contracts, ensuring near real-time updates.
- Provides configurable retry mechanisms and error handling for robust data ingestion from unreliable blockchain nodes.
- Includes a background process for periodically refreshing cached metadata and re-indexing data from specific blocks.
- Employs rate limiting and authentication mechanisms to protect the API from abuse and unauthorized access.
## Installation

```bash
go get github.com/uhsr/ErcNftMetadataIndexerNetwork
```

## Usage

```go
package main

import (
    "ercnftmetadataindexernetwork/internal/ercnftmetadataindexernetwork"
)

func main() {
    app := ercnftmetadataindexernetwork.NewApp(false)
    app.Run()
}
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
