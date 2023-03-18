package types

const (
	EventTypeLiquidStake                = "liquid_stake"
	EventTypeLiquidUnstake              = "liquid_unstake"
	EventTypeLiquidValidatorAdd         = "liquid_validator_add"
	EventTypeLiquidValidatorRemove      = "liquid_validator_remove"
	EventTypeBeginRestake               = "begin_restake"
	EventTypeBeginRebalance             = "begin_rebalance"
	EventTypeUnbondInactiveLiquidTokens = "unbond_inactive_liquid_tokens"

	AttributeKeyDelegator             = "delegator"
	AttributeKeyNewShares             = "new_shares"
	AttributeLiquidTokenMintedAmount  = "liquid_token_minted_amount"
	AttributeCompletionTime           = "completion_time"
	AttributeUnbondAmount             = "unbond_amount"
	AttributeLiquidValidator          = "liquid_validator"
	AttributeRedelegationSuccessCount = "redelegation_success_count"
	AttributeRedelegationFailureCount = "redelegation_failure_count"
	AttributeValueCategory            = ModuleName
)
