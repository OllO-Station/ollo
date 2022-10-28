<!--
Guiding Principles:

Changelogs are for humans, not machines.
There should be an entry for every single version.
The same types of changes should be grouped.
Versions and sections should be linkable.
The latest version comes first.
The release date of each version is displayed.
Mention whether you follow Semantic Versioning.

Usage:

Change log entries are to be added to the Unreleased section under the
appropriate stanza (see below). Each entry should ideally include a tag and
the Github issue reference in the following format:

* (<tag>) \#<issue-number> message

The issue numbers will later be link-ified during the release process so you do
not have to worry about including a link manually, but you can if you wish.

Types of changes (Stanzas):

"Features" for new features.
"Improvements" for changes in existing functionality.
"Deprecated" for soon-to-be removed features.
"Bug Fixes" for any bug fixes.
"Client Breaking" for breaking CLI commands and REST routes.
"State Machine Breaking" for breaking the AppState

Ref: https://keepachangelog.com/en/1.0.0/
-->

# Changelog

## [Latest]

## [v0.0.1] -2022-10-28

* (ollod) added [CHANGELOG.md](https://github.com/OllO-Station/ollo/blob/v0.0.1/CHANGELOG.md).
* (sdk) bump [cosmos-sdk](https://github.com/cosmos/cosmos-sdk) to [v0.46.3](https://github.com/cosmos/cosmos-sdk/releases/tag/v0.46.3). See [CHANGELOG.md](https://github.com/OllO-Station/blob/v0.46.3/CHANGELOG.md) for details.
  * (sdk) add [interchain account](https://github.com/cosmos/ibc-go/tree/main/modules/apps/27-interchain-accounts) module (interhchain-account module is part of ibc-go module).
  * (sdk) add [group module](https://github.com/cosmos/cosmos-sdk).
  * (sdk) add [nft module](https://github.com/cosmos/cosmos-sdk).
  

### Bug Fixes

* (ollod) slashing `downtime_jail_duration` reduced to `600000000000ns`

### Breaking Changes

* (ollod) `chain_id` updated to reflect new value `ollo-testnet-1`

### Features

* (ollod) User balances and accounts have been brought over at zero height

### Improvements

* (ollod) More user-friendly CLI formatting & subcommands
* (ollod) Changed `chain.schema.json` to reflect current values

<!-- Release links -->

[Unreleased]: https://github.com/OllO-Station/ollo/compare/v0.0.1...HEAD
[v0.0.1]: https://github.com/OllO-Station/ollo/releases/tag/v0.0.1