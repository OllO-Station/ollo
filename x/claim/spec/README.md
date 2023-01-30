<!--
order: 0
title: Claim Overview
parent:
  title: "claim"
-->

# `claim`

## Abstract

This document specifies the `claim` module developed by Ignite.

This module can be used by blockchains that wish to offer airdrops to eligible addresses upon the completion of specific actions.

Eligible addresses with airdrop allocations are listed in the genesis state of the module.

Initial claim, staking, and voting missions are natively supported. The developer can add custom missions related to their blockchain functionality. The `CompleteMission` method exposed by the module keeper can be used for blockchain specific missions.

## Contents

1. **[State](01_state.md)**
2. **[Messages](02_messages.md)**
3. **[Methods](03_methods.md)**
4. **[End-Blocker](04_end_blocker.md)**
5. **[Parameters](05_params.md)**
6. **[Events](06_events.md)**
7. **[Client](07_client.md)**
