<!--
order: 2
-->

# Messages

### `MsgClaim`

Claim completed mission amount for airdrop

```protobuf
message MsgClaim {
  string claimer = 1;
  uint64 missionID = 2;
}

message MsgClaimResponse {
  string claimed = 1 [
    (gogoproto.nullable) = false,
    (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Int",
    (cosmos_proto.scalar) = "cosmos.Int"
  ];
}
```

**State transition**

- Complete the claim for the mission and address
- Transfer the claim amount to the claimer balance

**Fails if**

- Mission is not completed
- The mission doesn't exist
- The claimer is not eligible
- The airdrop start time not reached
- The mission has already been claimed
