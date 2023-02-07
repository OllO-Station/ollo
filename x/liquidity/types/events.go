package types

// Event types for the liquidity module.
const (
	EventTypeCreatePool          = TypeMsgCreatePool
	EventTypeDepositWithinBatch  = TypeMsgDepositWithinBatch
	EventTypeWithdrawWithinBatch = TypeMsgWithdrawWithinBatch
	EventTypeSwapWithinBatch     = TypeMsgSwapWithinBatch
	EventTypeDepositToPool       = "deposit_to_pool"
	EventTypeWithdrawFromPool    = "withdraw_from_pool"
	EventTypeSwapTransacted      = "swap_transacted"

	AttributeValuePoolId         = "pool_id"      //nolint:revive
	AttributeValuePoolTypeId     = "pool_type_id" //nolint:revive
	AttributeValuePoolName       = "pool_name"
	AttributeValueReserveAccount = "reserve_account"
	AttributeValuePoolDenom      = "pool_denom"
	AttributeValuePoolCoinAmount = "pool_coin_amount"
	AttributeValueBatchIndex     = "batch_index"
	AttributeValueMsgIndex       = "msg_index"

	AttributeValueDepositCoins = "deposit_coins"

	AttributeValueOfferDenom             = "offer_denom"
	AttributeValueOfferCoinAmount        = "offer_coin_amount"
	AttributeValueOfferCoinFeeAmount     = "offer_coin_fee_amount"
	AttributeValueExchangedCoinFeeAmount = "exchanged_coin_fee_amount"
	AttributeValueDemandDenom            = "demand_denom"
	AttributeValueOrderPrice             = "order_price"

	AttributeValueDepositor        = "depositor"
	AttributeValueRefundedCoins    = "refunded_coins"
	AttributeValueAcceptedCoins    = "accepted_coins"
	AttributeValueSuccess          = "success"
	AttributeValueWithdrawer       = "withdrawer"
	AttributeValueWithdrawCoins    = "withdraw_coins"
	AttributeValueWithdrawFeeCoins = "withdraw_fee_coins"
	AttributeValueSwapRequester    = "swap_requester"
	AttributeValueSwapTypeId       = "swap_type_id" //nolint:revive
	AttributeValueSwapPrice        = "swap_price"

	AttributeValueTransactedCoinAmount       = "transacted_coin_amount"
	AttributeValueRemainingOfferCoinAmount   = "remaining_offer_coin_amount"
	AttributeValueExchangedOfferCoinAmount   = "exchanged_offer_coin_amount"
	AttributeValueExchangedDemandCoinAmount  = "exchanged_demand_coin_amount"
	AttributeValueReservedOfferCoinFeeAmount = "reserved_offer_coin_fee_amount"
	AttributeValueOrderExpiryHeight          = "order_expiry_height"

	AttributeValueCategory = ModuleName

	Success = "success"
	Failure = "failure"

	EventTypeCreatePair       = "create_pair"
	EventTypeCreateRangedPool = "create_ranged_pool"
	EventTypeDeposit          = "deposit"
	EventTypeWithdraw         = "withdraw"
	EventTypeLimitOrder       = "limit_order"
	EventTypeMarketOrder      = "market_order"
	EventTypeMMOrder          = "mm_order"
	EventTypeCancelOrder      = "cancel_order"
	EventTypeCancelAllOrders  = "cancel_all_orders"
	EventTypeCancelMMOrder    = "cancel_mm_order"
	EventTypeDepositResult    = "deposit_result"
	EventTypeWithdrawalResult = "withdrawal_result"
	EventTypeOrderResult      = "order_result"
	EventTypeUserOrderMatched = "user_order_matched"
	EventTypePoolOrderMatched = "pool_order_matched"

	AttributeKeyCreator            = "creator"
	AttributeKeyDepositor          = "depositor"
	AttributeKeyWithdrawer         = "withdrawer"
	AttributeKeyOrderer            = "orderer"
	AttributeKeyBaseDenom          = "base_denom"
	AttributeKeyQuoteDenom         = "quote_denom"
	AttributeKeyDepositCoins       = "deposit_coins"
	AttributeKeyAcceptedCoins      = "accepted_coins"
	AttributeKeyMintedPoolCoin     = "minted_pool_coin"
	AttributeKeyPoolCoin           = "pool_coin"
	AttributeKeyWithdrawnCoins     = "withdrawn_coins"
	AttributeKeyRefundedCoins      = "refunded_coins"
	AttributeKeyReserveAddress     = "reserve_address"
	AttributeKeyEscrowAddress      = "escrow_address"
	AttributeKeyRequestId          = "request_id"
	AttributeKeyPoolId             = "pool_id"
	AttributeKeyPairId             = "pair_id"
	AttributeKeyBatchId            = "batch_id"
	AttributeKeyOrderId            = "order_id"
	AttributeKeyOrderIds           = "order_ids"
	AttributeKeyOrderDirection     = "order_direction"
	AttributeKeyOfferCoin          = "offer_coin"
	AttributeKeyDemandDenom        = "demand_denom"
	AttributeKeyPrice              = "price"
	AttributeKeyAmount             = "amount"
	AttributeKeyOpenAmount         = "open_amount"
	AttributeKeyExpireAt           = "expire_at"
	AttributeKeyRemainingOfferCoin = "remaining_offer_coin"
	AttributeKeyReceivedCoin       = "received_coin"
	AttributeKeyPairIds            = "pair_ids"
	AttributeKeyCanceledOrderIds   = "canceled_order_ids"
	AttributeKeyStatus             = "status"
	AttributeKeyMatchedAmount      = "matched_amount"
	AttributeKeyPaidCoin           = "paid_coin"
)
