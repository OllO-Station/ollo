package types

import (
	"strconv"
	"strings"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"ollo/x/liquidity/amm"
	"sort"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	"github.com/tendermint/tendermint/crypto"
)

// AlphabeticalDenomPair returns denom pairs that are alphabetically sorted.
func AlphabeticalDenomPair(denom1, denom2 string) (resDenom1, resDenom2 string) {
	if denom1 > denom2 {
		return denom2, denom1
	}
	return denom1, denom2
}

// SortDenoms sorts denoms in alphabetical order.
func SortDenoms(denoms []string) []string {
	sort.Strings(denoms)
	return denoms
}

// GetCoinsTotalAmount returns total amount of all coins in sdk.Coins.
func GetCoinsTotalAmount(coins sdk.Coins) sdk.Int {
	totalAmount := sdk.ZeroInt()
	for _, coin := range coins {
		totalAmount = totalAmount.Add(coin.Amount)
	}
	return totalAmount
}

// ValidateReserveCoinLimit checks if total amounts of depositCoins exceed maxReserveCoinAmount.
func ValidateReserveCoinLimit(maxReserveCoinAmount sdk.Int, depositCoins sdk.Coins) error {
	totalAmount := GetCoinsTotalAmount(depositCoins)
	if maxReserveCoinAmount.IsZero() {
		return nil
	} else if totalAmount.GT(maxReserveCoinAmount) {
		return ErrInsufficientOfferCoin
	} else {
		return nil
	}
}

func GetOfferCoinFee(offerCoin sdk.Coin, swapFeeRate sdk.Dec) sdk.Coin {
	if swapFeeRate.IsZero() {
		return sdk.NewCoin(offerCoin.Denom, sdk.ZeroInt())
	}
	// apply half-ratio swap fee rate and ceiling
	// see https://ollo/issues/41 for details
	return sdk.NewCoin(offerCoin.Denom, sdk.NewDecFromInt(offerCoin.Amount).Mul(swapFeeRate.QuoInt64(2)).Ceil().TruncateInt()) // Ceil(offerCoin.Amount * (swapFeeRate/2))
}

func MustParseCoinsNormalized(coinStr string) sdk.Coins {
	coins, err := sdk.ParseCoinsNormalized(coinStr)
	if err != nil {
		panic(err)
	}
	return coins
}

func CheckOverflow(a, b sdk.Int) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = ErrIntOverflowLiquidity
		}
	}()
	a.Mul(b)
	a.Quo(b)
	b.Quo(a)
	return nil
}

func CheckOverflowWithDec(a, b sdk.Dec) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = ErrIntOverflowLiquidity
		}
	}()
	a.Mul(b)
	a.Quo(b)
	b.Quo(a)
	return nil
}

type AddressType int32

const (
	// the 32 bytes length address type of ADR 028.
	AddressType32Bytes AddressType = 0
	// the default 20 bytes length address type.
	AddressType20Bytes AddressType = 1
)

func DeriveAddress(addressType AddressType, moduleName, name string) sdk.AccAddress {
	switch addressType {
	case AddressType32Bytes:
		return sdk.AccAddress(address.Module(moduleName, []byte(name)))
	case AddressType20Bytes:
		return sdk.AccAddress(crypto.AddressHash([]byte(moduleName + name)))
	default:
		return sdk.AccAddress{}
	}
}

type sendCoinsTxKey struct {
	from, to string
}

type sendCoinsTx struct {
	from, to sdk.AccAddress
	amt      sdk.Coins
}

// BulkSendCoinsOperation holds a list of SendCoins operations for bulk execution.
type BulkSendCoinsOperation struct {
	txSet map[sendCoinsTxKey]*sendCoinsTx
	txs   []*sendCoinsTx
}

// NewBulkSendCoinsOperation returns an empty BulkSendCoinsOperation.
func NewBulkSendCoinsOperation() *BulkSendCoinsOperation {
	return &BulkSendCoinsOperation{
		txSet: map[sendCoinsTxKey]*sendCoinsTx{},
	}
}

// QueueSendCoins queues a BankKeeper.SendCoins operation for later execution.
func (op *BulkSendCoinsOperation) QueueSendCoins(fromAddr, toAddr sdk.AccAddress, amt sdk.Coins) {
	if amt.IsValid() && !amt.IsZero() {
		txKey := sendCoinsTxKey{fromAddr.String(), toAddr.String()}
		tx, ok := op.txSet[txKey]
		if !ok {
			tx = &sendCoinsTx{fromAddr, toAddr, sdk.Coins{}}
			op.txSet[txKey] = tx
			op.txs = append(op.txs, tx)
		}
		tx.amt = tx.amt.Add(amt...)
	}
}

// Run runs BankKeeper.InputOutputCoins once for queued operations.
func (op *BulkSendCoinsOperation) Run(ctx sdk.Context, bankKeeper BankKeeper) error {
	if len(op.txs) > 0 {
		var (
			inputs  []banktypes.Input
			outputs []banktypes.Output
		)
		for _, tx := range op.txs {
			inputs = append(inputs, banktypes.NewInput(tx.from, tx.amt))
			outputs = append(outputs, banktypes.NewOutput(tx.to, tx.amt))
		}
		return bankKeeper.InputOutputCoins(ctx, inputs, outputs)
	}
	return nil
}

// NewPoolResponse returns a new PoolResponse from given information.
func NewPoolResponse(pool Pool, rx, ry sdk.Coin, poolCoinSupply sdk.Int) PoolResponse {
	var price *sdk.Dec
	if !pool.Disabled {
		p := pool.AMMPool(rx.Amount, ry.Amount, sdk.Int{}).Price()
		price = &p
	}
	return PoolResponse{
		Id:             pool.Id,
		TypeId:         pool.TypeId,
		PairId:         pool.PairId,
		Creator:        pool.CreatorAddr,
		ReserveAddress: pool.Reserve.Addr,
		PoolCoinDenom:  pool.Reserve.Denom,
		PoolCoinSupply: poolCoinSupply,
		MinPrice:       pool.PriceRange.Min,
		MaxPrice:       pool.PriceRange.Max,
		Price:          price,
		Balances: PoolBalances{
			BaseCoin:  ry,
			QuoteCoin: rx,
		},
		PrevDepositRequestId:  pool.PrevDepositReqId,
		PrevWithdrawRequestId: pool.PrevWithdrawReqId,
		Inactive:              pool.Disabled,
	}
}

// IsTooSmallOrderAmount returns whether the order amount is too small for
// matching, based on the order price.
func IsTooSmallOrderAmount(amt sdk.Int, price sdk.Dec) bool {
	return amt.LT(amm.MinCoinAmount) || price.MulInt(amt).LT(sdk.NewDecFromInt(amm.MinCoinAmount))
}

// PriceLimits returns the lowest and the highest price limits with given last price
// and price limit ratio.
func PriceLimits(lastPrice, priceLimitRatio sdk.Dec, tickPrec int) (lowestPrice, highestPrice sdk.Dec) {
	lowestPrice = amm.PriceToUpTick(lastPrice.Mul(sdk.OneDec().Sub(priceLimitRatio)), tickPrec)
	highestPrice = amm.PriceToDownTick(lastPrice.Mul(sdk.OneDec().Add(priceLimitRatio)), tickPrec)
	return
}

func NewMMOrderIndex(orderer sdk.AccAddress, pairId uint64, orderIds []uint64) MarketMakingOrderId {
	return MarketMakingOrderId{
		CreatorAddr: orderer.String(),
		PairId:      pairId,
		OrderIds:    orderIds,
	}
}

func (index MarketMakingOrderId) GetOrderer() sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(index.CreatorAddr)
	if err != nil {
		panic(err)
	}
	return addr
}

// MMOrderTick holds information about each tick's price and amount of an MMOrder.
type MMOrderTick struct {
	OfferCoinAmount sdk.Int
	Price           sdk.Dec
	Amount          sdk.Int
}

// MMOrderTicks returns fairly distributed tick information with given parameters.
func MMOrderTicks(dir OrderDirection, minPrice, maxPrice sdk.Dec, amt sdk.Int, maxNumTicks, tickPrec int) (ticks []MMOrderTick) {
	ammDir := amm.OrderDirection(dir)
	if minPrice.Equal(maxPrice) {
		return []MMOrderTick{{OfferCoinAmount: amm.OfferCoinAmount(ammDir, minPrice, amt), Price: minPrice, Amount: amt}}
	}
	gap := maxPrice.Sub(minPrice).QuoInt64(int64(maxNumTicks - 1))
	switch dir {
	case OrderDirectionBuy:
		var prevP sdk.Dec
		for i := 0; i < maxNumTicks-1; i++ {
			p := amm.PriceToDownTick(minPrice.Add(gap.MulInt64(int64(i))), tickPrec)
			if prevP.IsNil() || !p.Equal(prevP) {
				ticks = append(ticks, MMOrderTick{
					Price: p,
				})
				prevP = p
			}
		}
		tickAmt := amt.QuoRaw(int64(len(ticks) + 1))
		for i := range ticks {
			ticks[i].Amount = tickAmt
			ticks[i].OfferCoinAmount = amm.OfferCoinAmount(ammDir, ticks[i].Price, ticks[i].Amount)
			amt = amt.Sub(tickAmt)
		}
		ticks = append(ticks, MMOrderTick{
			OfferCoinAmount: amm.OfferCoinAmount(ammDir, maxPrice, amt),
			Price:           maxPrice,
			Amount:          amt,
		})
	case OrderDirectionSell:
		var prevP sdk.Dec
		for i := 0; i < maxNumTicks-1; i++ {
			p := amm.PriceToUpTick(maxPrice.Sub(gap.MulInt64(int64(i))), tickPrec)
			if prevP.IsNil() || !p.Equal(prevP) {
				ticks = append(ticks, MMOrderTick{
					Price: p,
				})
				prevP = p
			}
		}
		tickAmt := amt.QuoRaw(int64(len(ticks) + 1))
		for i := range ticks {
			ticks[i].Amount = tickAmt
			ticks[i].OfferCoinAmount = amm.OfferCoinAmount(ammDir, ticks[i].Price, ticks[i].Amount)
			amt = amt.Sub(tickAmt)
		}
		ticks = append(ticks, MMOrderTick{
			OfferCoinAmount: amm.OfferCoinAmount(ammDir, minPrice, amt),
			Price:           minPrice,
			Amount:          amt,
		})
	}
	return
}

// FormatUint64s returns comma-separated string representation of
// a slice of uint64.
func FormatUint64s(us []uint64) (s string) {
	ss := make([]string, 0, len(us))
	for _, u := range us {
		ss = append(ss, strconv.FormatUint(u, 10))
	}
	return strings.Join(ss, ",")
}
