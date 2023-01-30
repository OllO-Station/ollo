<!--
order: 1
-->

# State

The state of the module indexes the following values:

- `Minter`: the minter is a space for holding current inflation information
- `Params`: parameter of the module

```
Minter: [] -> Minter
Params: [] -> Params
```

### `Minter`

`Minter` holds current inflation information, it contains the annual inflation rate, and the annual expected provisions

```proto
message Minter {
  string inflation = 1 [
    (gogoproto.nullable)   = false,
    (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Dec",
    (cosmos_proto.scalar)  = "cosmos.Dec"
  ];
  string annual_provisions = 2 [
    (gogoproto.nullable)   = false,
    (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Dec",
    (cosmos_proto.scalar)  = "cosmos.Dec"
  ];
}
```

### `Params`

Described in **[Parameters](03_params.md)**
