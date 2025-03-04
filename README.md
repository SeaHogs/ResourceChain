# ResourceChain

ResourceChain is a blockchain application built using the Cosmos SDK and Ignite CLI. It provides a basic CRUD interface for managing "resources" on the blockchain.

## Features
- Create a resource
- List all resources
- Get details of a specific resource
- Update resource details
- Delete a resource

## Prerequisites
Before running ResourceChain, ensure you have the following installed:
- [Go](https://go.dev/doc/install) (>=1.23)
- [Ignite CLI](https://docs.ignite.com/guide/install)
- [Cosmos SDK](https://docs.cosmos.network)

## Installation & Setup

### 1. Clone the Repository
```sh
git clone https://github.com/your-repo/resourcechain.git
cd resourcechain
```

### 2. Build & Start the Blockchain
```sh
ignite chain serve
```
This will compile the blockchain, initialize a local network, and start a node.

### 3. Interacting with the Blockchain

#### Create a Resource
```sh
resourcechaind tx resource create-resource "Resource Name" "Description" \
  --from alice --chain-id resourcechain --keyring-backend test --yes
```

#### List Resources
```sh
resourcechaind query resource list-resource
```

#### Get a Specific Resource
```sh
resourcechaind query resource show-resource <resource_id>
```
Example:
```sh
resourcechaind query resource show-resource 1
```

#### Update a Resource
```sh
resourcechaind tx resource update-resource <resource_id> "New Name" "New Description" \
  --from alice --chain-id resourcechain --keyring-backend test --yes
```

#### Delete a Resource
```sh
resourcechaind tx resource delete-resource <resource_id> \
  --from alice --chain-id resourcechain --keyring-backend test --yes
```

## Running in a Local Network
You can configure multiple nodes and validators using Ignite. To start a local multi-node network:
```sh
ignite chain start
```

## API & gRPC
Once the blockchain is running, you can interact using REST APIs:
```sh
curl http://localhost:1317/resource/resource/1
```
Or use gRPC clients to interact with the chain.

---
