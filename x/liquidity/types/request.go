package types

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewRequestDeposit returns a new RequestDeposit.
func NewRequestDeposit(msg *MsgDeposit, pool Pool, id uint64, msgHeight int64) RequestDeposit {
	return RequestDeposit{
		Id:            id,
		PoolId:        msg.PoolId,
		MsgHeight:     msgHeight,
		DepositorAddr: msg.Depositor,
		DepositAmt:    msg.DepositCoins,
		AcceptedAmt:   nil,
		PoolCoin:      sdk.NewCoin(pool.Reserve.Denom, sdk.ZeroInt()),
		Status:        RequestStatusPending,
	}
}

func (req RequestDeposit) GetDepositor() sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(req.DepositorAddr)
	if err != nil {
		panic(err)
	}
	return addr
}

// Validate validates RequestDeposit for genesis.
func (req RequestDeposit) Validate() error {
	if req.Id == 0 {
		return fmt.Errorf("id must not be 0")
	}
	if req.PoolId == 0 {
		return fmt.Errorf("pool id must not be 0")
	}
	if req.MsgHeight == 0 {
		return fmt.Errorf("message height must not be 0")
	}
	if _, err := sdk.AccAddressFromBech32(req.DepositorAddr); err != nil {
		return fmt.Errorf("invalid depositor address %s: %w", req.DepositorAddr, err)
	}
	if err := req.DepositAmt.Validate(); err != nil {
		return fmt.Errorf("invalid deposit coins: %w", err)
	}
	if len(req.DepositAmt) == 0 || len(req.DepositAmt) > 2 {
		return fmt.Errorf("wrong number of deposit coins: %d", len(req.DepositAmt))
	}
	if err := req.AcceptedAmt.Validate(); err != nil {
		return fmt.Errorf("invalid accepted coins: %w", err)
	}
	if len(req.AcceptedAmt) > 2 {
		return fmt.Errorf("wrong number of accepted coins: %d", len(req.AcceptedAmt))
	}
	for _, coin := range req.AcceptedAmt {
		if req.DepositAmt.AmountOf(coin.Denom).IsZero() {
			return fmt.Errorf("mismatching denom pair between deposit coins and accepted coins")
		}
	}
	if err := req.PoolCoin.Validate(); err != nil {
		return fmt.Errorf("invalid minted pool coin %s: %w", req.PoolCoin, err)
	}
	if !req.Status.IsValid() {
		return fmt.Errorf("invalid status: %s", req.Status)
	}
	return nil
}

// SetStatus sets the request's status.
// SetStatus is to easily find locations where the status is changed.
func (req *RequestDeposit) SetStatus(status RequestStatus) {
	req.Status = status
}

// NewRequestWithdraw returns a new RequestWithdraw.
func NewRequestWithdraw(msg *MsgWithdraw, id uint64, msgHeight int64) RequestWithdraw {
	return RequestWithdraw{
		Id:           id,
		PoolId:       msg.PoolId,
		MsgHeight:    msgHeight,
		WithdrawAddr: msg.Withdrawer,
		PoolCoin:     msg.PoolCoin,
		WithdrawAmt:  nil,
		Status:       RequestStatusPending,
	}
}

func (req RequestWithdraw) GetWithdrawer() sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(req.WithdrawAddr)
	if err != nil {
		panic(err)
	}
	return addr
}

// Validate validates RequestWithdraw for genesis.
func (req RequestWithdraw) Validate() error {
	if req.Id == 0 {
		return fmt.Errorf("id must not be 0")
	}
	if req.PoolId == 0 {
		return fmt.Errorf("pool id must not be 0")
	}
	if req.MsgHeight == 0 {
		return fmt.Errorf("message height must not be 0")
	}
	if _, err := sdk.AccAddressFromBech32(req.WithdrawAddr); err != nil {
		return fmt.Errorf("invalid withdrawer address %s: %w", req.WithdrawAddr, err)
	}
	if err := req.PoolCoin.Validate(); err != nil {
		return fmt.Errorf("invalid pool coin %s: %w", req.PoolCoin, err)
	}
	if req.PoolCoin.IsZero() {
		return fmt.Errorf("pool coin must not be 0")
	}
	if err := req.WithdrawAmt.Validate(); err != nil {
		return fmt.Errorf("invalid withdrawn coins: %w", err)
	}
	if len(req.WithdrawAmt) > 2 {
		return fmt.Errorf("wrong number of withdrawn coins: %d", len(req.WithdrawAmt))
	}
	if !req.Status.IsValid() {
		return fmt.Errorf("invalid status: %s", req.Status)
	}
	return nil
}

// SetStatus sets the request's status.
// SetStatus is to easily find locations where the status is changed.
func (req *RequestWithdraw) SetStatus(status RequestStatus) {
	req.Status = status
}

// NewOrderForOrderLimit returns a new Order from MsgOrderLimit.
func NewOrderForOrderLimit(msg *MsgOrderLimit, id uint64, pair Pair, offerCoin sdk.Coin, price sdk.Dec, expires time.Time, msgHeight int64) Order {
	return Order{
		Type:        OrderTypeLimit,
		Id:          id,
		PairId:      pair.Id,
		MsgHeight:   msgHeight,
		CreatorAddr: msg.Orderer,
		Direction:   msg.Direction,
		Offer:       offerCoin,
		Remaining:   offerCoin,
		Received:    sdk.NewCoin(msg.DemandCoinDenom, sdk.ZeroInt()),
		Price:       price,
		Amt:         msg.Amount,
		OpenAmt:     msg.Amount,
		BatchId:     pair.CurrentBatchId,
		Expires:     expires,
		Status:      OrderStatusMatching,
	}
}

// NewOrderForOrderMarket returns a new Order from MsgOrderMarket.
func NewOrderForOrderMarket(msg *MsgOrderMarket, id uint64, pair Pair, offerCoin sdk.Coin, price sdk.Dec, expires time.Time, msgHeight int64) Order {
	return Order{
		Type:        OrderTypeMarket,
		Id:          id,
		PairId:      pair.Id,
		MsgHeight:   msgHeight,
		CreatorAddr: msg.Orderer,
		Direction:   msg.Direction,
		Offer:       offerCoin,
		Remaining:   offerCoin,
		Received:    sdk.NewCoin(msg.DemandCoinDenom, sdk.ZeroInt()),
		Price:       price,
		Amt:         msg.Amount,
		OpenAmt:     msg.Amount,
		BatchId:     pair.CurrentBatchId,
		Expires:     expires,
		Status:      OrderStatusMatching,
	}
}

func NewOrder(
	typ OrderType, id uint64, pair Pair, orderer sdk.AccAddress,
	offerCoin sdk.Coin, price sdk.Dec, amt sdk.Int, expires time.Time, msgHeight int64) Order {
	var (
		dir         OrderDirection
		demandDenom string
	)
	if offerCoin.Denom == pair.BaseDenom {
		dir = OrderDirectionSell
		demandDenom = pair.QuoteDenom
	} else {
		dir = OrderDirectionBuy
		demandDenom = pair.BaseDenom
	}
	return Order{
		Type:        typ,
		Id:          id,
		PairId:      pair.Id,
		MsgHeight:   msgHeight,
		CreatorAddr: orderer.String(),
		Direction:   dir,
		Offer:       offerCoin,
		Remaining:   offerCoin,
		Received:    sdk.NewCoin(demandDenom, sdk.ZeroInt()),
		Price:       price,
		Amt:         amt,
		OpenAmt:     amt,
		BatchId:     pair.CurrentBatchId,
		Expires:     expires,
		Status:      OrderStatusMatching,
	}
}

func (order Order) GetOrderer() sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(order.CreatorAddr)
	if err != nil {
		panic(err)
	}
	return addr
}

// Validate validates Order for genesis.
func (order Order) Validate() error {
	if order.Id == 0 {
		return fmt.Errorf("id must not be 0")
	}
	if order.PairId == 0 {
		return fmt.Errorf("pair id must not be 0")
	}
	if order.MsgHeight == 0 {
		return fmt.Errorf("message height must not be 0")
	}
	if _, err := sdk.AccAddressFromBech32(order.CreatorAddr); err != nil {
		return fmt.Errorf("invalid orderer address %s: %w", order.CreatorAddr, err)
	}
	if order.Direction != OrderDirectionBuy && order.Direction != OrderDirectionSell {
		return fmt.Errorf("invalid direction: %s", order.Direction)
	}
	if err := order.Offer.Validate(); err != nil {
		return fmt.Errorf("invalid offer coin %s: %w", order.Offer, err)
	}
	if order.Offer.IsZero() {
		return fmt.Errorf("offer coin must not be 0")
	}
	if err := order.Remaining.Validate(); err != nil {
		return fmt.Errorf("invalid remaining offer coin %s: %w", order.Remaining, err)
	}
	if order.Offer.Denom != order.Remaining.Denom {
		return fmt.Errorf("offer coin denom %s != remaining offer coin denom %s", order.Offer.Denom, order.Remaining.Denom)
	}
	if err := order.Received.Validate(); err != nil {
		return fmt.Errorf("invalid received coin %s: %w", order.Received, err)
	}
	if !order.Price.IsPositive() {
		return fmt.Errorf("price must be positive: %s", order.Price)
	}
	if !order.Amt.IsPositive() {
		return fmt.Errorf("amount must be positive: %s", order.Amt)
	}
	if order.OpenAmt.IsNegative() {
		return fmt.Errorf("open amount must not be negative: %s", order.OpenAmt)
	}
	if order.BatchId == 0 {
		return fmt.Errorf("batch id must not be 0")
	}
	if order.Expires.IsZero() {
		return fmt.Errorf("no expiration info")
	}
	if !order.Status.IsValid() {
		return fmt.Errorf("invalid status: %s", order.Status)
	}
	return nil
}

// ExpiredAt returns whether the order should be deleted at given time.
func (order Order) ExpiredAt(t time.Time) bool {
	return !order.Expires.After(t)
}

// SetStatus sets the order's status.
// SetStatus is to easily find locations where the status is changed.
func (order *Order) SetStatus(status OrderStatus) {
	order.Status = status
}

// IsValid returns true if the RequestStatus is one of:
// RequestStatusMatching, RequestStatusSucceeded, RequestStatusFailed.
func (status RequestStatus) IsValid() bool {
	switch status {
	case RequestStatusPending, RequestStatusSuccess, RequestStatusFail:
		return true
	default:
		return false
	}
}

// ShouldBeDeleted returns true if the RequestStatus is one of:
// RequestStatusSucceeded, RequestStatusFailed.
func (status RequestStatus) ShouldBeDeleted() bool {
	switch status {
	case RequestStatusSuccess, RequestStatusFail:
		return true
	default:
		return false
	}
}

// IsValid returns true if the OrderStatus is one of:
// OrderStatusMatching, OrderStatusNotMatched, OrderStatusPartiallyMatched,
// OrderStatusCompleted, OrderStatusCanceled, OrderStatusExpired.
func (status OrderStatus) IsValid() bool {
	switch status {
	case OrderStatusMatching, OrderStatusNoMatch, OrderStatusPartialMatch,
		OrderStatusMatched, OrderStatusCanceled, OrderStatusExpired:
		return true
	default:
		return false
	}
}

// IsMatchable returns true if the OrderStatus is one of:
// OrderStatusMatching, OrderStatusNotMatched, OrderStatusPartiallyMatched.
func (status OrderStatus) IsMatchable() bool {
	switch status {
	case OrderStatusMatching, OrderStatusNoMatch, OrderStatusPartialMatch:
		return true
	default:
		return false
	}
}

// CanBeExpired has the same condition as IsMatchable.
func (status OrderStatus) CanBeExpired() bool {
	return status.IsMatchable()
}

// CanBeCanceled returns true if the OrderStatus is one of:
// OrderStatusMatching, OrderStatusNotMatched, OrderStatusPartiallyMatched.
func (status OrderStatus) CanBeCanceled() bool {
	switch status {
	case OrderStatusMatching, OrderStatusNoMatch, OrderStatusPartialMatch:
		return true
	default:
		return false
	}
}

// IsCanceledOrExpired returns true if the OrderStatus is one of:
// OrderStatusCanceled, OrderStatusExpired.
func (status OrderStatus) IsCanceledOrExpired() bool {
	switch status {
	case OrderStatusCanceled, OrderStatusExpired:
		return true
	default:
		return false
	}
}

// ShouldBeDeleted returns true if the OrderStatus is one of:
// OrderStatusCompleted, OrderStatusCanceled, OrderStatusExpired.
func (status OrderStatus) ShouldBeDeleted() bool {
	return status == OrderStatusMatched || status.IsCanceledOrExpired()
}

// MustMarshalRequestDeposit returns the RequestDeposit bytes. Panics if fails.
func MustMarshalRequestDeposit(cdc codec.BinaryCodec, msg RequestDeposit) []byte {
	return cdc.MustMarshal(&msg)
}

// UnmarshalRequestDeposit returns the RequestDeposit from bytes.
func UnmarshalRequestDeposit(cdc codec.BinaryCodec, value []byte) (msg RequestDeposit, err error) {
	err = cdc.Unmarshal(value, &msg)
	return msg, err
}

// MustUnmarshalRequestDeposit returns the RequestDeposit from bytes.
// It throws panic if it fails.
func MustUnmarshalRequestDeposit(cdc codec.BinaryCodec, value []byte) RequestDeposit {
	msg, err := UnmarshalRequestDeposit(cdc, value)
	if err != nil {
		panic(err)
	}
	return msg
}

// MustMarshaRequestWithdraw returns the RequestWithdraw bytes.
// It throws panic if it fails.
func MustMarshaRequestWithdraw(cdc codec.BinaryCodec, msg RequestWithdraw) []byte {
	return cdc.MustMarshal(&msg)
}

// UnmarshalRequestWithdraw returns the RequestWithdraw from bytes.
func UnmarshalRequestWithdraw(cdc codec.BinaryCodec, value []byte) (msg RequestWithdraw, err error) {
	err = cdc.Unmarshal(value, &msg)
	return msg, err
}

// MustUnmarshalRequestWithdraw returns the RequestWithdraw from bytes.
// It throws panic if it fails.
func MustUnmarshalRequestWithdraw(cdc codec.BinaryCodec, value []byte) RequestWithdraw {
	msg, err := UnmarshalRequestWithdraw(cdc, value)
	if err != nil {
		panic(err)
	}
	return msg
}

// MustMarshaOrder returns the Order bytes.
// It throws panic if it fails.
func MustMarshaOrder(cdc codec.BinaryCodec, order Order) []byte {
	return cdc.MustMarshal(&order)
}

// UnmarshalOrder returns the Order from bytes.
func UnmarshalOrder(cdc codec.BinaryCodec, value []byte) (order Order, err error) {
	err = cdc.Unmarshal(value, &order)
	return order, err
}

// MustUnmarshalOrder returns the Order from bytes.
// It throws panic if it fails.
func MustUnmarshalOrder(cdc codec.BinaryCodec, value []byte) Order {
	msg, err := UnmarshalOrder(cdc, value)
	if err != nil {
		panic(err)
	}
	return msg
}
