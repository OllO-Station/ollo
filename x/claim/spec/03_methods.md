<!--
order: 3
-->

# Methods

### `CompleteMission`

Complete a mission for an eligible address.
This method can be used by an external chain importing `claim` in order to define customized mission for the chain.

```go
CompleteMission(
    ctx sdk.Context,
    missionID uint64,
    address string,
) error
```

**State transition**

- Complete the mission `missionID` in the claim record `address`

**Fails if**

- The mission doesn't exist
- The address has no claim record
- The mission has already been completed for the address

### `ClaimMission`

Claim mission for an eligible claim record and mission id.
This method can be used by an external module importing `claim` in order to define customized mission claims for the
chain.

```go
ClaimMission(
    ctx sdk.Context,
    claimRecord types.ClaimRecord,
    missionID uint64,
) error
```

**State transition**

- Transfer the claim amount related to the mission

**Fails if**

- The mission doesn't exist
- The address has no claim record
- The airdrop start time not reached
- The mission has not been completed for the address
- The mission has already been claimed for the address
