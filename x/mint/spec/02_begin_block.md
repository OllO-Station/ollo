<!--
order: 2
-->

# Begin-block

Begin-block contains the logic to:

- recalculate minter parameters
- mint new coins
- distribute new coins depending on distribution proportions

### Pseudo-code

```go
minter = load(Minter)
params = load(Params)
minter = calculateInflationAndAnnualProvision(params)
store(Minter, minter)

mintedCoins = minter.BlockProvision(params)
Mint(mintedCoins)

DistributeMintedCoins(mintedCoin)
```

The inflation rate calculation follows the same logic as the [Cosmos SDK `mint` module](https://github.com/cosmos/cosmos-sdk/tree/main/x/mint#inflation-rate-calculation)
