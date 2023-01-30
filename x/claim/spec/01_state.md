<!--
order: 1
-->

# State

The state of the module stores data for the three following properties:

- Who: what is the list of eligible addresses, what is the allocation for each eligible address
- How: what are the missions to claim airdrops
- When: when does decaying for the airdrop start, and when does the airdrop end

The state of the module indexes the following values:

- `AirdropSupply`: the amount of tokens that remain for the airdrop
- `InitialClaim`: information about an initial claim, a portion of the airdrop that can be claimed without completing a specific task
- `ClaimRecords`: list of eligible addresses with allocated airdrop, and the current status of completed missions
- `Missions`: the list of missions to claim airdrop with their associated weight
- `Params`: parameter of the module

```
AirdropSupply:  [] -> sdk.Int
InitialClaim:   [] -> InitialClaim

ClaimRecords:   [address] -> ClaimRecord
Missions:       [id] -> Mission

Params:         [] -> Params
```

### `InitialClaim`

`InitialClaim` determines the rules for the initial claim, a portion of the airdrop that can be directly claimed without completing a specific task. The mission is completed by sending a `MsgClaim` message.

The structure determines if the initial claim is enabled for the chain, and what mission is completed when sending `MsgClaim`.

```protobuf
message InitialClaim {
  bool   enabled   = 1;
  uint64 missionID = 2;
}
```

### `ClaimRecord`

`ClaimRecord` contains information about an address eligible for airdrop, what amount the address is eligible for, and which missions have already been completed and claimed.

```protobuf
message ClaimRecord {
  string address   = 1 [(cosmos_proto.scalar) = "cosmos.AddressString"];
  string claimable = 2 [
    (gogoproto.nullable)   = false,
    (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Int",
    (cosmos_proto.scalar)  = "cosmos.Int"
  ];
  repeated uint64 completedMissions = 3;
  repeated uint64 claimedMissions = 4;
}
```

### `Mission`

`Mission` represents a mission to be completed to claim a percentage of the airdrop supply.

```protobuf
message Mission {
  uint64 missionID   = 1;
  string description = 2;
  string weight      = 3 [
    (gogoproto.nullable)   = false,
    (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Dec",
    (cosmos_proto.scalar)  = "cosmos.Dec"
  ];
}
```

### `Params`

Described in **[Parameters](05_params.md)**
