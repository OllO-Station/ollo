<!--
order: 4
-->

# End-blocker

The end-blocker of the module verifies if the airdrop supply is non-null, decay is enabled and decay end has been reached.
Under these conditions, the remaining airdrop supply is transferred to the community pool.

### Pseudo-code

```go
airdropSupply = load(AirdropSupply)
decayInfo = load(Params).DecayInformation

if airdropSupply > 0 && decayInfo.Enabled && BlockTime > decayInfo.DecayEnd
    distrKeeper.FundCommunityPool(airdropSupply)
    airdropSupply = 0
    store(AirdropSupply, airdropSupply)
```
