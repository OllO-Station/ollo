<!--
order: 5
-->

# Client

## CLI

### Query

The `query` commands allow users to query `mint` state.

```sh
testappd q mint
```

#### `params`

Shows the params of the module.

```sh
testappd q mint params
```

Example output:

```yml
blocks_per_year: "6311520"
distribution_proportions:
  community_pool: "0.300000000000000000"
  funded_addresses: "0.400000000000000000"
  staking: "0.300000000000000000"
funded_addresses:
  - address: cosmos1ezptsm3npn54qx9vvpah4nymre59ykr9967vj9
    weight: "0.400000000000000000"
  - address: cosmos1aqn8ynvr3jmq67879qulzrwhchq5dtrvh6h4er
    weight: "0.300000000000000000"
  - address: cosmos1pkdk6m2nh77nlaep84cylmkhjder3areczme3w
    weight: "0.300000000000000000"
goal_bonded: "0.670000000000000000"
inflation_max: "0.200000000000000000"
inflation_min: "0.070000000000000000"
inflation_rate_change: "0.130000000000000000"
mint_denom: stake
```

#### `annual-provisions`

Shows the current minting annual provisions valu

```sh
testappd q mint annual-provisions
```

Example output:

```yml
52000470.516851147993560400
```

#### `inflation`

Shows the current minting inflation value

```sh
testappd q mint inflation
```

Example output:

```yml
0.130001213701730800
```
