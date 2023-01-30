<!--
order: 6
-->

# Events

### `EventMissionCompleted`

This event is emitted when a mission `missionID` is completed for a specific eligible `address`.

```protobuf
message EventMissionCompleted {
  uint64 missionID = 1;
  string address = 2;
}
```

### `EventMissionClaimed`

This event is emitted when a mission `missionID` is claimed for a specific eligible address `claimer`.

```protobuf
message EventMissionClaimed {
  uint64 missionID = 1;
  string claimer = 2;
}
```
