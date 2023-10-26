<!--
order: 0
title: Exchange Overview
parent:
  title: "exchange"
-->
# Atomic Marketplace

## Overview

The Atomic Marketplace is a decentralized exchange platform built on the Cosmos SDK, allowing users to trade digital assets in a secure and efficient manner. The platform enables traditional trading functionalities along with Atomic Swaps, providing a more flexible and inclusive trading environment.

## Features

- **Atomic Swaps**: Secure, trustless trading between two parties without the need for a central custodian.
- **Traditional Trading**: Limit and market orders for various trading strategies.
- **Multi-Asset Support**: Trade a variety of assets, including native tokens and wrapped tokens from other blockchains.
- **Decentralized**: Built on top of the Cosmos SDK, ensuring full decentralization and interoperability.
- **Efficient Matching Engine**: High-performance trading with a low-latency matching engine.
  
## Installation

### Requirements

- Go 1.16+
- Cosmos SDK
- Docker (Optional but recommended)

### Build and Install

```bash
# Clone the repository
git clone https://github.com/your-repo/atomic-marketplace.git

# Navigate to the project directory
cd atomic-marketplace

# Install dependencies
go mod tidy

# Build the application
go build -o atomic-marketplace .

# Initialize the node (replace `your-moniker` with your preferred name)
./atomic-marketplace init your-moniker
```

## Usage

### Running a Node

```bash
./atomic-marketplace start
```

### Creating a Market

```bash
atomic-marketplace tx atomic-marketplace create-market --base-denom="uatom" --quote-denom="uusd" --from=mykey
```

### Placing an Order

```bash
atomic-marketplace tx atomic-marketplace place-order --market-id=1 --is-buy=true --price="1.0uusd" --quantity="100uatom" --from=mykey
```

### Executing Atomic Swap

```bash
atomic-marketplace tx atomic-marketplace execute-atomic-swap --from=address1 --to=address2 --amount=100uatom --secret=secret123 --secret-hash=hash123 --expire-height=10000
```

## Testing

Run the test suite to ensure the application is working as expected.

```bash
go test ./...
```

## Contribution

Contributions are welcome! Please read the [Contribution Guide](CONTRIBUTING.md) for more information.

## License

This project is licensed under the MIT License. See [LICENSE](LICENSE) for more details.

## Support and Community

For support, you can join our [Discord Channel](#) or [Telegram Group](#). You can also create an issue in the GitHub repository.

---

For more details, check out the [full documentation](#).

This README is meant to be a starting point. As your project grows, consider adding more in-depth guides in a `/docs` directory or using a documentation tool like MkDocs or Docusaurus.

# `exchange`

## Abstract

## Contents

1. [Concepts](01_concepts.md)
2. [State](02_state.md)
    * [Market](02_state.md#market)
    * [Order](02_state.md#order)
3. [Messages](03_messages.md)
    * [MsgCreateMarket](03_messages.md#msgcreatemarket)
    * [MsgPlaceLimitOrder](03_messages.md#msgplacelimitorder)
    * [MsgPlaceBatchLimitOrder](03_messages.md#msgplacebatchlimitorder)
    * [MsgPlaceMMLimitOrder](03_messages.md#msgplacemmlimitorder)
    * [MsgPlaceMMBatchLimitOrder](03_messages.md#msgplacemmbatchlimitorder)
    * [MsgPlaceMarketOrder](03_messages.md#msgplacemarketorder)
    * [MsgCancelOrder](03_messages.md#msgcancelorder)
    * [MsgCancelAllOrders](03_messages.md#msgcancelallorders)
    * [MsgSwapExactAmountIn](03_messages.md#msgswapexactamountin)
4. [Begin-Block](04_begin_block.md)
5. [Events](05_events.md)
    * [Begin-Block](05_events.md#begin-block)
    * [Handlers](05_events.md#handlers)
6. [Parameters](06_params.md)
