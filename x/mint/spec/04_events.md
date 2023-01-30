<!--
order: 4
-->

# Events

### `EventMint`

This event is emitted when new coins are minted. The event contains the amount of coins minted with the parameters of the minter at the current block.

```protobuf
message EventMint {
  string bondedRatio = 1 [
    (gogoproto.nullable) = false,
    (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Dec",
    (cosmos_proto.scalar) = "cosmos.Dec"
  ];
  string inflation = 2 [
    (gogoproto.nullable) = false,
    (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Dec",
    (cosmos_proto.scalar) = "cosmos.Dec"
  ];
  string annualProvisions = 3 [
    (gogoproto.nullable) = false,
    (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Dec",
    (cosmos_proto.scalar) = "cosmos.Dec"
  ];
  string amount = 4 [
    (gogoproto.nullable) = false,
    (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Int",
    (cosmos_proto.scalar) = "cosmos.Int"
  ];
}
```
