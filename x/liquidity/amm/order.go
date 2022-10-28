package amm

import (
	"fmt"
	"sort"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ Order   = (*BaseOrder)(nil)
	_ Orderer = (*BaseOrderer)(nil)

	DefaultOrderer = BaseOrderer{}
)

// OrderDirection specifies an order direction, either buy or sell.
type OrderDirection int

// OrderDirection enumerations.
const (
	Buy OrderDirection = iota + 1
	Sell
)

func (dir OrderDirection) String() string {
	switch dir {
	case Buy:
		return "Buy"
	case Sell:
		return "Sell"
	default:
		return fmt.Sprintf("OrderDirection(%d)", dir)
	}
}

type Orderer interface {
	Order(dir OrderDirection, price sdk.Dec, amt sdk.Int) Order
}

// BaseOrderer creates new BaseOrder with sufficient offer coin amount
// considering price and amount.
type BaseOrderer struct{}

func (orderer BaseOrderer) Order(dir OrderDirection, price sdk.Dec, amt sdk.Int) Order {
	return NewBaseOrder(dir, price, amt, OfferCoinAmount(dir, price, amt))
}

// Order is the universal interface of an order.
type Order interface {
	GetDirection() OrderDirection
	// GetBatchId returns the batch id where the order was created.
	// Batch id of 0 means the current batch.
	GetBatchId() uint64
	GetPrice() sdk.Dec
	GetAmount() sdk.Int // The original order amount
	GetOfferCoinAmount() sdk.Int
	GetPaidOfferCoinAmount() sdk.Int
	SetPaidOfferCoinAmount(amt sdk.Int)
	GetReceivedDemandCoinAmount() sdk.Int
	SetReceivedDemandCoinAmount(amt sdk.Int)
	GetOpenAmount() sdk.Int
	SetOpenAmount(amt sdk.Int)
	IsMatched() bool
	// HasPriority returns true if the order has higher priority
	// than the other order.
	HasPriority(other Order) bool
	String() string
}

// BaseOrder is the base struct for an Order.
type BaseOrder struct {
	Direction       OrderDirection
	Price           sdk.Dec
	Amount          sdk.Int
	OfferCoinAmount sdk.Int

	// Match info
	OpenAmount               sdk.Int
	PaidOfferCoinAmount      sdk.Int
	ReceivedDemandCoinAmount sdk.Int
}

// NewBaseOrder returns a new BaseOrder.
func NewBaseOrder(dir OrderDirection, price sdk.Dec, amt, offerCoinAmt sdk.Int) *BaseOrder {
	return &BaseOrder{
		Direction:                dir,
		Price:                    price,
		Amount:                   amt,
		OfferCoinAmount:          offerCoinAmt,
		OpenAmount:               amt,
		PaidOfferCoinAmount:      sdk.ZeroInt(),
		ReceivedDemandCoinAmount: sdk.ZeroInt(),
	}
}

// GetDirection returns the order direction.
func (order *BaseOrder) GetDirection() OrderDirection {
	return order.Direction
}

func (order *BaseOrder) GetBatchId() uint64 {
	return 0
}

// GetPrice returns the order price.
func (order *BaseOrder) GetPrice() sdk.Dec {
	return order.Price
}

// GetAmount returns the order amount.
func (order *BaseOrder) GetAmount() sdk.Int {
	return order.Amount
}

func (order *BaseOrder) GetOfferCoinAmount() sdk.Int {
	return order.OfferCoinAmount
}

func (order *BaseOrder) GetPaidOfferCoinAmount() sdk.Int {
	return order.PaidOfferCoinAmount
}

func (order *BaseOrder) SetPaidOfferCoinAmount(amt sdk.Int) {
	order.PaidOfferCoinAmount = amt
}

func (order *BaseOrder) GetReceivedDemandCoinAmount() sdk.Int {
	return order.ReceivedDemandCoinAmount
}

func (order *BaseOrder) SetReceivedDemandCoinAmount(amt sdk.Int) {
	order.ReceivedDemandCoinAmount = amt
}

func (order *BaseOrder) GetOpenAmount() sdk.Int {
	return order.OpenAmount
}

func (order *BaseOrder) SetOpenAmount(amt sdk.Int) {
	order.OpenAmount = amt
}

func (order *BaseOrder) IsMatched() bool {
	return order.OpenAmount.LT(order.Amount)
}

// HasPriority returns whether the order has higher priority than
// the other order.
func (order *BaseOrder) HasPriority(other Order) bool {
	return order.Amount.GT(other.GetAmount())
}

func (order *BaseOrder) String() string {
	return fmt.Sprintf("BaseOrder(%s,%s,%s)", order.Direction, order.Price, order.Amount)
}

// OrderBook is an order book.
type OrderBook struct {
	buys, sells *orderBookTicks
}

// NewOrderBook returns a new OrderBook.
func NewOrderBook(orders ...Order) *OrderBook {
	ob := &OrderBook{
		buys:  newOrderBookBuyTicks(),
		sells: newOrderBookSellTicks(),
	}
	ob.AddOrder(orders...)
	return ob
}

// AddOrder adds orders to the order book.
func (ob *OrderBook) AddOrder(orders ...Order) {
	for _, order := range orders {
		if MatchableAmount(order, order.GetPrice()).IsPositive() {
			switch order.GetDirection() {
			case Buy:
				ob.buys.addOrder(order)
			case Sell:
				ob.sells.addOrder(order)
			}
		}
	}
}

// Orders returns all orders in the order book.
func (ob *OrderBook) Orders() []Order {
	var orders []Order
	for _, tick := range append(ob.buys.ticks, ob.sells.ticks...) {
		orders = append(orders, tick.orders...)
	}
	return orders
}

// BuyOrdersAt returns buy orders at given price in the order book.
// Note that the orders are not sorted.
func (ob *OrderBook) BuyOrdersAt(price sdk.Dec) []Order {
	return ob.buys.ordersAt(price)
}

// SellOrdersAt returns sell orders at given price in the order book.
// Note that the orders are not sorted.
func (ob *OrderBook) SellOrdersAt(price sdk.Dec) []Order {
	return ob.sells.ordersAt(price)
}

func (ob *OrderBook) HighestPrice() (sdk.Dec, bool) {
	highestBuyPrice, _, foundBuy := ob.buys.highestPrice()
	highestSellPrice, _, foundSell := ob.sells.highestPrice()
	switch {
	case foundBuy && foundSell:
		return sdk.MaxDec(highestBuyPrice, highestSellPrice), true
	case foundBuy:
		return highestBuyPrice, true
	case foundSell:
		return highestSellPrice, true
	default:
		return sdk.Dec{}, false
	}
}

func (ob *OrderBook) LowestPrice() (sdk.Dec, bool) {
	lowestBuyPrice, _, foundBuy := ob.buys.lowestPrice()
	lowestSellPrice, _, foundSell := ob.sells.lowestPrice()
	switch {
	case foundBuy && foundSell:
		return sdk.MinDec(lowestBuyPrice, lowestSellPrice), true
	case foundBuy:
		return lowestBuyPrice, true
	case foundSell:
		return lowestSellPrice, true
	default:
		return sdk.Dec{}, false
	}
}

func (ob *OrderBook) stringRepresentation(prices []sdk.Dec) string {
	if len(prices) == 0 {
		return "<nil>"
	}
	sort.Slice(prices, func(i, j int) bool {
		return prices[i].GT(prices[j])
	})
	var b strings.Builder
	b.WriteString("+--------sell--------+------------price-------------+--------buy---------+\n")
	for _, price := range prices {
		buyAmt := TotalMatchableAmount(ob.BuyOrdersAt(price), price)
		sellAmt := TotalMatchableAmount(ob.SellOrdersAt(price), price)
		_, _ = fmt.Fprintf(&b, "| %18s | %28s | %-18s |\n", sellAmt, price.String(), buyAmt)
	}
	b.WriteString("+--------------------+------------------------------+--------------------+")
	return b.String()
}

// FullString returns a full string representation of the order book.
// FullString includes all possible price ticks from the order book's
// highest price to the lowest price.
func (ob *OrderBook) FullString(tickPrec int) string {
	var prices []sdk.Dec
	highest, found := ob.HighestPrice()
	if !found {
		return "<nil>"
	}
	lowest, _ := ob.LowestPrice()
	for ; lowest.LTE(highest); lowest = UpTick(lowest, tickPrec) {
		prices = append(prices, lowest)
	}
	return ob.stringRepresentation(prices)
}

// String returns a compact string representation of the order book.
// String includes a tick only when there is at least one order on it.
func (ob *OrderBook) String() string {
	priceSet := map[string]sdk.Dec{}
	for _, tick := range append(ob.buys.ticks, ob.sells.ticks...) {
		priceSet[tick.price.String()] = tick.price
	}
	prices := make([]sdk.Dec, 0, len(priceSet))
	for _, price := range priceSet {
		prices = append(prices, price)
	}
	return ob.stringRepresentation(prices)
}

// orderBookTicks represents a list of orderBookTick.
// This type is used for both buy/sell sides of OrderBook.
type orderBookTicks struct {
	ticks           []*orderBookTick
	priceIncreasing bool
}

func newOrderBookBuyTicks() *orderBookTicks {
	return &orderBookTicks{
		priceIncreasing: false,
	}
}

func newOrderBookSellTicks() *orderBookTicks {
	return &orderBookTicks{
		priceIncreasing: true,
	}
}

func (ticks *orderBookTicks) findPrice(price sdk.Dec) (i int, exact bool) {
	i = sort.Search(len(ticks.ticks), func(i int) bool {
		if ticks.priceIncreasing {
			return ticks.ticks[i].price.GTE(price)
		} else {
			return ticks.ticks[i].price.LTE(price)
		}
	})
	if i < len(ticks.ticks) && ticks.ticks[i].price.Equal(price) {
		exact = true
	}
	return
}

func (ticks *orderBookTicks) addOrder(order Order) {
	i, exact := ticks.findPrice(order.GetPrice())
	if exact {
		ticks.ticks[i].addOrder(order)
	} else {
		if i < len(ticks.ticks) {
			// Insert a new order book tick at index i.
			ticks.ticks = append(ticks.ticks[:i], append([]*orderBookTick{newOrderBookTick(order)}, ticks.ticks[i:]...)...)
		} else {
			// Append a new order book tick at the end.
			ticks.ticks = append(ticks.ticks, newOrderBookTick(order))
		}
	}
}

func (ticks *orderBookTicks) ordersAt(price sdk.Dec) []Order {
	i, exact := ticks.findPrice(price)
	if !exact {
		return nil
	}
	return ticks.ticks[i].orders
}

func (ticks *orderBookTicks) highestPrice() (sdk.Dec, int, bool) {
	if len(ticks.ticks) == 0 {
		return sdk.Dec{}, 0, false
	}
	if ticks.priceIncreasing {
		return ticks.ticks[len(ticks.ticks)-1].price, len(ticks.ticks) - 1, true
	} else {
		return ticks.ticks[0].price, 0, true
	}
}

func (ticks *orderBookTicks) lowestPrice() (sdk.Dec, int, bool) {
	if len(ticks.ticks) == 0 {
		return sdk.Dec{}, 0, false
	}
	if ticks.priceIncreasing {
		return ticks.ticks[0].price, 0, true
	} else {
		return ticks.ticks[len(ticks.ticks)-1].price, len(ticks.ticks) - 1, true
	}
}

// orderBookTick represents a tick in OrderBook.
type orderBookTick struct {
	price  sdk.Dec
	orders []Order
}

func newOrderBookTick(order Order) *orderBookTick {
	return &orderBookTick{
		price:  order.GetPrice(),
		orders: []Order{order},
	}
}

func (tick *orderBookTick) addOrder(order Order) {
	tick.orders = append(tick.orders, order)
}

var (
	_ OrderView = (*OrderBookView)(nil)
	_ OrderView = Pool(nil)
	_ OrderView = MultipleOrderViews(nil)
)

type OrderView interface {
	HighestBuyPrice() (sdk.Dec, bool)
	LowestSellPrice() (sdk.Dec, bool)
	BuyAmountOver(price sdk.Dec, inclusive bool) sdk.Int
	SellAmountUnder(price sdk.Dec, inclusive bool) sdk.Int
}

type OrderBookView struct {
	buyAmtAccSums, sellAmtAccSums []amtAccSum
}

func (ob *OrderBook) MakeView() *OrderBookView {
	view := &OrderBookView{
		buyAmtAccSums:  make([]amtAccSum, len(ob.buys.ticks)),
		sellAmtAccSums: make([]amtAccSum, len(ob.sells.ticks)),
	}
	for i, tick := range ob.buys.ticks {
		var prevSum sdk.Int
		if i == 0 {
			prevSum = sdk.ZeroInt()
		} else {
			prevSum = view.buyAmtAccSums[i-1].sum
		}
		view.buyAmtAccSums[i] = amtAccSum{
			price: tick.price,
			sum:   prevSum.Add(TotalMatchableAmount(tick.orders, tick.price)),
		}
	}
	for i, tick := range ob.sells.ticks {
		var prevSum sdk.Int
		if i == 0 {
			prevSum = sdk.ZeroInt()
		} else {
			prevSum = view.sellAmtAccSums[i-1].sum
		}
		view.sellAmtAccSums[i] = amtAccSum{
			price: tick.price,
			sum:   prevSum.Add(TotalMatchableAmount(tick.orders, tick.price)),
		}
	}
	return view
}

func (view *OrderBookView) Match() {
	if len(view.buyAmtAccSums) == 0 || len(view.sellAmtAccSums) == 0 {
		return
	}
	buyIdx := sort.Search(len(view.buyAmtAccSums), func(i int) bool {
		return view.BuyAmountOver(view.buyAmtAccSums[i].price, true).GT(
			view.SellAmountUnder(view.buyAmtAccSums[i].price, true))
	})
	sellIdx := sort.Search(len(view.sellAmtAccSums), func(i int) bool {
		return view.SellAmountUnder(view.sellAmtAccSums[i].price, true).GT(
			view.BuyAmountOver(view.sellAmtAccSums[i].price, true))
	})
	if buyIdx == len(view.buyAmtAccSums) && sellIdx == len(view.sellAmtAccSums) {
		return
	}
	matchAmt := sdk.ZeroInt()
	if buyIdx > 0 {
		matchAmt = view.buyAmtAccSums[buyIdx-1].sum
	}
	if sellIdx > 0 {
		matchAmt = sdk.MaxInt(matchAmt, view.sellAmtAccSums[sellIdx-1].sum)
	}
	for i, accSum := range view.buyAmtAccSums {
		if i < buyIdx {
			view.buyAmtAccSums[i].sum = zeroInt
		} else {
			view.buyAmtAccSums[i].sum = accSum.sum.Sub(matchAmt)
		}
	}
	for i, accSum := range view.sellAmtAccSums {
		if i < sellIdx {
			view.sellAmtAccSums[i].sum = zeroInt
		} else {
			view.sellAmtAccSums[i].sum = accSum.sum.Sub(matchAmt)
		}
	}
}

func (view *OrderBookView) HighestBuyPrice() (sdk.Dec, bool) {
	if len(view.buyAmtAccSums) == 0 {
		return sdk.Dec{}, false
	}
	i := sort.Search(len(view.buyAmtAccSums), func(i int) bool {
		return view.buyAmtAccSums[i].sum.IsPositive()
	})
	if i >= len(view.buyAmtAccSums) {
		return sdk.Dec{}, false
	}
	return view.buyAmtAccSums[i].price, true
}

func (view *OrderBookView) LowestSellPrice() (sdk.Dec, bool) {
	if len(view.sellAmtAccSums) == 0 {
		return sdk.Dec{}, false
	}
	i := sort.Search(len(view.sellAmtAccSums), func(i int) bool {
		return view.sellAmtAccSums[i].sum.IsPositive()
	})
	if i >= len(view.sellAmtAccSums) {
		return sdk.Dec{}, false
	}
	return view.sellAmtAccSums[i].price, true
}

func (view *OrderBookView) BuyAmountOver(price sdk.Dec, inclusive bool) sdk.Int {
	i := sort.Search(len(view.buyAmtAccSums), func(i int) bool {
		if inclusive {
			return view.buyAmtAccSums[i].price.LT(price)
		} else {
			return view.buyAmtAccSums[i].price.LTE(price)
		}
	})
	if i == 0 {
		return sdk.ZeroInt()
	}
	return view.buyAmtAccSums[i-1].sum
}

func (view *OrderBookView) BuyAmountUnder(price sdk.Dec, inclusive bool) sdk.Int {
	i := sort.Search(len(view.buyAmtAccSums), func(i int) bool {
		if inclusive {
			return view.buyAmtAccSums[i].price.LTE(price)
		} else {
			return view.buyAmtAccSums[i].price.LT(price)
		}
	})
	if i == 0 {
		return view.buyAmtAccSums[len(view.buyAmtAccSums)-1].sum
	}
	return view.buyAmtAccSums[len(view.buyAmtAccSums)-1].sum.Sub(view.buyAmtAccSums[i-1].sum)
}

func (view *OrderBookView) SellAmountUnder(price sdk.Dec, inclusive bool) sdk.Int {
	i := sort.Search(len(view.sellAmtAccSums), func(i int) bool {
		if inclusive {
			return view.sellAmtAccSums[i].price.GT(price)
		} else {
			return view.sellAmtAccSums[i].price.GTE(price)
		}
	})
	if i == 0 {
		return sdk.ZeroInt()
	}
	return view.sellAmtAccSums[i-1].sum
}

func (view *OrderBookView) SellAmountOver(price sdk.Dec, inclusive bool) sdk.Int {
	i := sort.Search(len(view.sellAmtAccSums), func(i int) bool {
		if inclusive {
			return view.sellAmtAccSums[i].price.GTE(price)
		} else {
			return view.sellAmtAccSums[i].price.GT(price)
		}
	})
	if i == 0 {
		return view.sellAmtAccSums[len(view.sellAmtAccSums)-1].sum
	}
	return view.sellAmtAccSums[len(view.sellAmtAccSums)-1].sum.Sub(view.sellAmtAccSums[i-1].sum)
}

type amtAccSum struct {
	price sdk.Dec
	sum   sdk.Int
}

type MultipleOrderViews []OrderView

func (views MultipleOrderViews) HighestBuyPrice() (price sdk.Dec, found bool) {
	for _, view := range views {
		p, f := view.HighestBuyPrice()
		if f && (price.IsNil() || p.GT(price)) {
			price = p
			found = true
		}
	}
	return
}

func (views MultipleOrderViews) LowestSellPrice() (price sdk.Dec, found bool) {
	for _, view := range views {
		p, f := view.LowestSellPrice()
		if f && (price.IsNil() || p.LT(price)) {
			price = p
			found = true
		}
	}
	return
}

func (views MultipleOrderViews) BuyAmountOver(price sdk.Dec, inclusive bool) sdk.Int {
	amt := sdk.ZeroInt()
	for _, view := range views {
		amt = amt.Add(view.BuyAmountOver(price, inclusive))
	}
	return amt
}

func (views MultipleOrderViews) SellAmountUnder(price sdk.Dec, inclusive bool) sdk.Int {
	amt := sdk.ZeroInt()
	for _, view := range views {
		amt = amt.Add(view.SellAmountUnder(price, inclusive))
	}
	return amt
}
