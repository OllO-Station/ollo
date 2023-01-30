<!--
order: 5
-->

# Parameters

The parameters of the module contain information about time-based decay for the airdrop supply.

```protobuf
message Params {
  DecayInformation decayInformation = 1 [(gogoproto.nullable) = false];
  google.protobuf.Timestamp airdropStart = 2
  [(gogoproto.nullable) = false, (gogoproto.stdtime) = true];
}
```

### `DecayInformation`

This parameter determines if the airdrop starts to decay at a specific time.

```protobuf
message DecayInformation {
  bool enabled = 1;
  google.protobuf.Timestamp decayStart = 2 [(gogoproto.nullable) = false, (gogoproto.stdtime) = true];
  google.protobuf.Timestamp decayEnd = 3 [(gogoproto.nullable) = false, (gogoproto.stdtime) = true];
}
```

When enabled, the claimable amount for each eligible address will start to decrease starting from `decayStart` till `decayEnd` where the airdrop is ended and the remaining fund is transferred to the community pool.

The decrease is linear.

If `decayStart == decayEnd`, there is no decay for the airdrop but the airdrop ends at `decayEnd` and the remaining fund is transferred to the community pool.

### `AirdropStart`

This parameter determines the airdrop start time.
When set, the user cannot claim the airdrop after completing the mission. The airdrop will be available only after the block time reaches the airdrop start time.
If the mission was completed, the user could call the `MsgClaim` to claim an airdrop from a completed mission.
The claim will be called automatically if the mission is completed and the airdrop start time is already reached.