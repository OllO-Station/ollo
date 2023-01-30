<!--
order: 3
-->

# Parameters

The parameters of the module contain information about inflation, and distribution of minted coins.

- `mint_denom`: the denom of the minted coins
- `inflation_rate_change`: maximum annual change in inflation rate
- `inflation_max`: maximum inflation rate
- `inflation_min`: minimum inflation rate
- `goal_bonded`: goal of percent bonded coins
- `blocks_per_year`: expected blocks per year
- `distribution_proportions`: distribution_proportions defines the proportion for minted coins distribution
- `funded_addresses`: list of funded addresses

```proto
message Params {
  string mint_denom = 1;
  string inflation_rate_change = 2 [
    (gogoproto.nullable)   = false,
    (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Dec",
    (cosmos_proto.scalar)  = "cosmos.Dec"
  ];
  string inflation_max = 3 [
    (gogoproto.nullable)   = false,
    (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Dec",
    (cosmos_proto.scalar)  = "cosmos.Dec"
  ];
  string inflation_min = 4 [
    (gogoproto.nullable)   = false,
    (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Dec",
    (cosmos_proto.scalar)  = "cosmos.Dec"
  ];
  string goal_bonded = 5 [
    (gogoproto.nullable)   = false,
    (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Dec",
    (cosmos_proto.scalar)  = "cosmos.Dec"
  ];
  uint64 blocks_per_year = 6;
  DistributionProportions distribution_proportions = 7 [(gogoproto.nullable) = false];
  repeated WeightedAddress funded_addresses = 8 [(gogoproto.nullable) = false];
}
```

### `DistributionProportions`

`DistributionProportions` contains propotions for the distributions.

```proto
message DistributionProportions {
  string staking = 1 [
    (gogoproto.nullable)   = false,
    (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Dec",
    (cosmos_proto.scalar)  = "cosmos.Dec"
  ];
  string funded_addresses = 2 [
    (gogoproto.nullable)   = false,
    (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Dec",
    (cosmos_proto.scalar)  = "cosmos.Dec"
  ];
  string community_pool = 3 [
    (gogoproto.nullable)   = false,
    (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Dec",
    (cosmos_proto.scalar)  = "cosmos.Dec"
  ];
}
```

### `WeightedAddress`

`WeightedAddress` is an address with an associated weight to receive part the minted coins depending on the `funded_addresses` distribution proportion.

```proto
message WeightedAddress {
  string address = 1 [(cosmos_proto.scalar) = "cosmos.AddressString"];
  string weight  = 2 [
    (gogoproto.nullable)   = false,
    (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Dec",
    (cosmos_proto.scalar)  = "cosmos.Dec"
  ];
}
```
