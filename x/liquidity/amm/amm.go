package amm

import (
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"sort"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// The minimum and maximum coin amount used in the amm package.

var (
	MinCoinAmount = sdk.NewInt(100)
	MaxCoinAmount = sdk.NewInt(10000000000000)
)

var (
	MinPoolPrice                    = sdk.NewDecWithPrec(1, 15) // 10^-15
	MaxPoolPrice                    = sdk.NewDecFromInt(sdk.NewInt(2))
	MinIntelligentPoolPriceGapRatio = sdk.NewDecWithPrec(1, 3) // 0.001, 0.1%
)

// TickPrecision represents a tick precision.
type TickPrecision int

func (prec TickPrecision) PriceToDownTick(price sdk.Dec) sdk.Dec {
	return PriceToDownTick(price, int(prec))
}

func (prec TickPrecision) PriceToUpTick(price sdk.Dec) sdk.Dec {
	return PriceToUpTick(price, int(prec))
}

func (prec TickPrecision) UpTick(price sdk.Dec) sdk.Dec {
	return UpTick(price, int(prec))
}

func (prec TickPrecision) DownTick(price sdk.Dec) sdk.Dec {
	return DownTick(price, int(prec))
}

func (prec TickPrecision) HighestTick() sdk.Dec {
	return HighestTick(int(prec))
}

func (prec TickPrecision) LowestTick() sdk.Dec {
	return LowestTick(int(prec))
}

func (prec TickPrecision) TickToIndex(tick sdk.Dec) int {
	return TickToIndex(tick, int(prec))
}

func (prec TickPrecision) TickFromIndex(i int) sdk.Dec {
	return TickFromIndex(i, int(prec))
}

func (prec TickPrecision) RoundPrice(price sdk.Dec) sdk.Dec {
	return RoundPrice(price, int(prec))
}

func (prec TickPrecision) TickGap(price sdk.Dec) sdk.Dec {
	return TickGap(price, int(prec))
}

func (prec TickPrecision) RandomTick(r *rand.Rand, minPrice, maxPrice sdk.Dec) sdk.Dec {
	return RandomTick(r, minPrice, maxPrice, int(prec))
}

// char returns the characteristic(integral part) of
// log10(x * pow(10, sdk.Precision)).
func char(x sdk.Dec) int {
	if x.IsZero() {
		panic("cannot calculate log10 for 0")
	}
	return len(x.BigInt().Text(10)) - 1
}

// pow10 returns pow(10, n - sdk.Precision).
func pow10(n int) sdk.Dec {
	x := big.NewInt(10)
	x.Exp(x, big.NewInt(int64(n)), nil)
	return sdk.NewDecFromBigIntWithPrec(x, sdk.Precision)
}

// isPow10 returns whether x is a power of 10 or not.
func isPow10(x sdk.Dec) bool {
	b := x.BigInt()
	if b.Sign() <= 0 {
		return false
	}
	ten := big.NewInt(10)
	if b.Cmp(ten) == 0 {
		return true
	}
	zero := big.NewInt(0)
	m := new(big.Int)
	for b.Cmp(ten) >= 0 {
		b.DivMod(b, ten, m)
		if m.Cmp(zero) != 0 {
			return false
		}
	}
	return b.Cmp(big.NewInt(1)) == 0
}

// PriceToDownTick returns the highest price tick under(or equal to) the price.
func PriceToDownTick(price sdk.Dec, prec int) sdk.Dec {
	b := price.BigInt()
	l := char(price)
	d := int64(l - prec)
	if d > 0 {
		p := big.NewInt(10)
		p.Exp(p, big.NewInt(d), nil)
		b.Quo(b, p).Mul(b, p)
	}
	return sdk.NewDecFromBigIntWithPrec(b, sdk.Precision)
}

// PriceToUpTick returns the lowest price tick greater or equal than
// the price.
func PriceToUpTick(price sdk.Dec, prec int) sdk.Dec {
	tick := PriceToDownTick(price, prec)
	if !tick.Equal(price) {
		return UpTick(tick, prec)
	}
	return tick
}

// UpTick returns the next lowest price tick above the price.
func UpTick(price sdk.Dec, prec int) sdk.Dec {
	tick := PriceToDownTick(price, prec)
	if tick.Equal(price) {
		l := char(price)
		return price.Add(pow10(l - prec))
	}
	l := char(tick)
	return tick.Add(pow10(l - prec))
}

// DownTick returns the next highest price tick under the price.
// DownTick doesn't check if the price is the lowest price tick.
func DownTick(price sdk.Dec, prec int) sdk.Dec {
	tick := PriceToDownTick(price, prec)
	if tick.Equal(price) {
		l := char(price)
		var d sdk.Dec
		if isPow10(price) {
			d = pow10(l - prec - 1)
		} else {
			d = pow10(l - prec)
		}
		return price.Sub(d)
	}
	return tick
}

// HighestTick returns the highest possible price tick.
func HighestTick(prec int) sdk.Dec {
	i := new(big.Int).SetBits([]big.Word{0, 0, 0, 0, 0x1000000000000000})
	return PriceToDownTick(sdk.NewDecFromBigIntWithPrec(i, sdk.Precision), prec)
}

// LowestTick returns the lowest possible price tick.
func LowestTick(prec int) sdk.Dec {
	return sdk.NewDecWithPrec(1, int64(sdk.Precision-prec))
}

// TickToIndex returns a tick index for given price.
// Tick index 0 means the lowest possible price fit in ticks.
func TickToIndex(price sdk.Dec, prec int) int {
	b := price.BigInt()
	l := len(b.Text(10)) - 1
	d := int64(l - prec)
	if d > 0 {
		q := big.NewInt(10)
		q.Exp(q, big.NewInt(d), nil)
		b.Quo(b, q)
	}
	p := int(math.Pow10(prec))
	b.Sub(b, big.NewInt(int64(p)))
	return (l-prec)*9*p + int(b.Int64())
}

// TickFromIndex returns a price for given tick index.
// See TickToIndex for more details about tick indices.
func TickFromIndex(i, prec int) sdk.Dec {
	p := int(math.Pow10(prec))
	l := i/(9*p) + prec
	t := big.NewInt(int64(p + i%(p*9)))
	if l > prec {
		m := big.NewInt(10)
		m.Exp(m, big.NewInt(int64(l-prec)), nil)
		t.Mul(t, m)
	}
	return sdk.NewDecFromBigIntWithPrec(t, sdk.Precision)
}

// RoundTickIndex returns rounded tick index using banker's rounding.
func RoundTickIndex(i int) int {
	return (i + 1) / 2 * 2
}

// RoundPrice returns rounded price using banker's rounding.
func RoundPrice(price sdk.Dec, prec int) sdk.Dec {
	tick := PriceToDownTick(price, prec)
	if price.Equal(tick) {
		return price
	}
	return TickFromIndex(RoundTickIndex(TickToIndex(tick, prec)), prec)
}

// TickGap returns tick gap at given price.
func TickGap(price sdk.Dec, prec int) sdk.Dec {
	tick := PriceToDownTick(price, prec)
	l := char(tick)
	return pow10(l - prec)
}

// RandomTick returns a random tick within range [minPrice, maxPrice].
// If prices are not on ticks, then prices are adjusted to the nearest
// ticks.
func RandomTick(r *rand.Rand, minPrice, maxPrice sdk.Dec, prec int) sdk.Dec {
	minPrice = PriceToUpTick(minPrice, prec)
	maxPrice = PriceToDownTick(maxPrice, prec)
	minPriceIdx := TickToIndex(minPrice, prec)
	maxPriceIdx := TickToIndex(maxPrice, prec)
	return TickFromIndex(minPriceIdx+r.Intn(maxPriceIdx-minPriceIdx), prec)
}

var (
	zeroInt = sdk.ZeroInt()
	oneDec  = sdk.OneDec()
	fourDec = sdk.NewDec(4)
)

// OfferCoinAmount returns the minimum offer coin amount for
// given order direction, price and order amount.
func OfferCoinAmount(dir OrderDirection, price sdk.Dec, amt sdk.Int) sdk.Int {
	switch dir {
	case Buy:
		return price.MulInt(amt).Ceil().TruncateInt()
	case Sell:
		return amt
	default:
		panic(fmt.Sprintf("invalid order direction: %s", dir))
	}
}

// MatchableAmount returns matchable amount of an order considering
// remaining offer coin and price.
func MatchableAmount(order Order, price sdk.Dec) (matchableAmt sdk.Int) {
	switch order.GetDirection() {
	case Buy:
		remainingOfferCoinAmt := order.GetOfferCoinAmount().Sub(order.GetPaidOfferCoinAmount())
		matchableAmt = sdk.MinInt(
			order.GetOpenAmount(),
			sdk.NewDecFromInt(remainingOfferCoinAmt).QuoTruncate(price).TruncateInt(),
		)
	case Sell:
		matchableAmt = order.GetOpenAmount()
	}
	if price.MulInt(matchableAmt).TruncateInt().IsZero() {
		matchableAmt = zeroInt
	}
	return
}

// TotalAmount returns total amount of orders.
func TotalAmount(orders []Order) sdk.Int {
	amt := sdk.ZeroInt()
	for _, order := range orders {
		amt = amt.Add(order.GetAmount())
	}
	return amt
}

// TotalMatchableAmount returns total matchable amount of orders.
func TotalMatchableAmount(orders []Order, price sdk.Dec) (amt sdk.Int) {
	amt = sdk.ZeroInt()
	for _, order := range orders {
		amt = amt.Add(MatchableAmount(order, price))
	}
	return
}

// OrderGroup represents a group of orders with same batch id.
type OrderGroup struct {
	BatchId uint64
	Orders  []Order
}

// GroupOrdersByBatchId groups orders by their batch id and returns a
// slice of OrderGroup.
func GroupOrdersByBatchId(orders []Order) (groups []*OrderGroup) {
	groupByBatchId := map[uint64]*OrderGroup{}
	for _, order := range orders {
		group, ok := groupByBatchId[order.GetBatchId()]
		if !ok {
			i := sort.Search(len(groups), func(i int) bool {
				if order.GetBatchId() == 0 {
					return groups[i].BatchId == 0
				}
				if groups[i].BatchId == 0 {
					return true
				}
				return order.GetBatchId() <= groups[i].BatchId
			})
			group = &OrderGroup{BatchId: order.GetBatchId()}
			groupByBatchId[order.GetBatchId()] = group
			groups = append(groups[:i], append([]*OrderGroup{group}, groups[i:]...)...)
		}
		group.Orders = append(group.Orders, order)
	}
	return
}

// SortOrders sorts orders using its HasPriority condition.
func SortOrders(orders []Order) {
	sort.SliceStable(orders, func(i, j int) bool {
		return orders[i].HasPriority(orders[j])
	})
}

// findFirstTrueCondition uses the binary search to find the first index
// where f(i) is true, while searching in range [start, end].
// It assumes that f(j) == false where j < i and f(j) == true where j >= i.
// start can be greater than end.
func findFirstTrueCondition(start, end int, f func(i int) bool) (i int, found bool) {
	if start < end {
		i = start + sort.Search(end-start+1, func(i int) bool {
			return f(start + i)
		})
		if i > end {
			return 0, false
		}
		return i, true
	}
	i = start - sort.Search(start-end+1, func(i int) bool {
		return f(start - i)
	})
	if i < end {
		return 0, false
	}
	return i, true
}

// inv returns the inverse of x.
func inv(x sdk.Dec) (r sdk.Dec) {
	r = oneDec.Quo(x)
	return
}

var (
	// Pool price gap ratio function thresholds
	t1 = sdk.MustNewDecFromStr("0.01")
	t2 = sdk.MustNewDecFromStr("0.02")
	t3 = sdk.MustNewDecFromStr("0.1")

	// Pool price gap ratio function coefficients
	a1, b1 = sdk.MustNewDecFromStr("0.007"), sdk.MustNewDecFromStr("0.00003")
	a2, b2 = sdk.MustNewDecFromStr("0.09"), sdk.MustNewDecFromStr("-0.0008")
	a3     = sdk.MustNewDecFromStr("0.05")
	b4     = sdk.MustNewDecFromStr("0.005")
)

func poolOrderPriceGapRatio(poolPrice, currentPrice sdk.Dec) (r sdk.Dec) {
	if poolPrice.IsZero() {
		poolPrice = sdk.NewDecWithPrec(1, sdk.Precision) // lowest possible sdk.Dec
	}
	x := currentPrice.Sub(poolPrice).Abs().Quo(poolPrice)
	switch {
	case x.LTE(t1):
		return a1.Mul(x).Add(b1)
	case x.LTE(t2):
		return a2.Mul(x).Add(b2)
	case x.LTE(t3):
		return a3.Mul(x)
	default:
		return b4
	}
}
