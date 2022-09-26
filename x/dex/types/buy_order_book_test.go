package types_test

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"

	"ollo/x/dex/types"
)

func OrderListToBuyOrderBook(list []types.Order) types.BuyOrderBook {
	listCopy := make([]*types.Order, len(list))
	for i, order := range list {
		order := order
		listCopy[i] = &order
	}

	book := types.BuyOrderBook{
		AmountDenom: "foo",
		PriceDenom:  "bar",
		Book: &types.OrderBook{
			IdCount: 0,
			Orders:  listCopy,
		},
	}
	return book
}

func TestAppendOrder(t *testing.T) {
	buyBook := types.NewBuyOrderBook(GenPair())

	// Prevent zero amount
	seller, amount, price := GenOrder()
	_, err := buyBook.AppendOrder(seller, 0, price)
	require.ErrorIs(t, err, types.ErrZeroAmount)

	// Prevent big amount
	_, err = buyBook.AppendOrder(seller, types.MaxAmount+1, price)
	require.ErrorIs(t, err, types.ErrMaxAmount)

	// Prevent zero price
	_, err = buyBook.AppendOrder(seller, amount, 0)
	require.ErrorIs(t, err, types.ErrZeroPrice)

	// Prevent big price
	_, err = buyBook.AppendOrder(seller, amount, types.MaxPrice+1)
	require.ErrorIs(t, err, types.ErrMaxPrice)

	// Can append buy orders
	for i := 0; i < 20; i++ {
		// Append a new order
		creator, amount, price := GenOrder()
		newOrder := types.Order{
			Id:      buyBook.Book.IdCount,
			Creator: creator,
			Amount:  amount,
			Price:   price,
		}
		orderID, err := buyBook.AppendOrder(creator, amount, price)

		// Checks
		require.NoError(t, err)
		require.Contains(t, buyBook.Book.Orders, &newOrder)
		require.Equal(t, newOrder.Id, orderID)
	}

	require.Len(t, buyBook.Book.Orders, 20)
	require.True(t, sort.SliceIsSorted(buyBook.Book.Orders, func(i, j int) bool {
		return buyBook.Book.Orders[i].Price < buyBook.Book.Orders[j].Price
	}))
}

type liquidateSellRes struct {
	Book       []types.Order
	Remaining  types.Order
	Liquidated types.Order
	Gain       int32
	Match      bool
	Filled     bool
}

func simulateLiquidateFromSellOrder(
	t *testing.T,
	inputList []types.Order,
	inputOrder types.Order,
	expected liquidateSellRes,
) {
	book := OrderListToBuyOrderBook(inputList)
	expectedBook := OrderListToBuyOrderBook(expected.Book)

	require.True(t, sort.SliceIsSorted(book.Book.Orders, func(i, j int) bool {
		return book.Book.Orders[i].Price < book.Book.Orders[j].Price
	}))
	require.True(t, sort.SliceIsSorted(expectedBook.Book.Orders, func(i, j int) bool {
		return expectedBook.Book.Orders[i].Price < expectedBook.Book.Orders[j].Price
	}))

	remaining, liquidated, gain, match, filled := book.LiquidateFromSellOrder(inputOrder)

	require.Equal(t, expectedBook, book)
	require.Equal(t, expected.Remaining, remaining)
	require.Equal(t, expected.Liquidated, liquidated)
	require.Equal(t, expected.Gain, gain)
	require.Equal(t, expected.Match, match)
	require.Equal(t, expected.Filled, filled)
}

func TestLiquidateFromSellOrder(t *testing.T) {
	// No match for empty book
	inputOrder := types.Order{Id: 10, Creator: MockAccount("1"), Amount: 100, Price: 30}
	book := OrderListToBuyOrderBook([]types.Order{})
	_, _, _, match, _ := book.LiquidateFromSellOrder(inputOrder)
	require.False(t, match)

	// Buy book
	inputBook := []types.Order{
		{Id: 2, Creator: MockAccount("2"), Amount: 30, Price: 15},
		{Id: 1, Creator: MockAccount("1"), Amount: 200, Price: 20},
		{Id: 0, Creator: MockAccount("0"), Amount: 50, Price: 25},
	}

	// Test no match if highest bid too low (25 < 30)
	book = OrderListToBuyOrderBook(inputBook)
	_, _, _, match, _ = book.LiquidateFromSellOrder(inputOrder)
	require.False(t, match)

	// Entirely filled (30 < 50)
	inputOrder = types.Order{Id: 10, Creator: MockAccount("1"), Amount: 30, Price: 22}
	expected := liquidateSellRes{
		Book: []types.Order{
			{Id: 2, Creator: MockAccount("2"), Amount: 30, Price: 15},
			{Id: 1, Creator: MockAccount("1"), Amount: 200, Price: 20},
			{Id: 0, Creator: MockAccount("0"), Amount: 20, Price: 25},
		},
		Remaining:  types.Order{Id: 10, Creator: MockAccount("1"), Amount: 0, Price: 22},
		Liquidated: types.Order{Id: 0, Creator: MockAccount("0"), Amount: 30, Price: 25},
		Gain:       int32(30 * 25),
		Match:      true,
		Filled:     true,
	}
	simulateLiquidateFromSellOrder(t, inputBook, inputOrder, expected)

	// Entirely filled and liquidated ( 50 = 50)
	inputOrder = types.Order{Id: 10, Creator: MockAccount("1"), Amount: 50, Price: 15}
	expected = liquidateSellRes{
		Book: []types.Order{
			{Id: 2, Creator: MockAccount("2"), Amount: 30, Price: 15},
			{Id: 1, Creator: MockAccount("1"), Amount: 200, Price: 20},
		},
		Remaining:  types.Order{Id: 10, Creator: MockAccount("1"), Amount: 0, Price: 15},
		Liquidated: types.Order{Id: 0, Creator: MockAccount("0"), Amount: 50, Price: 25},
		Gain:       int32(50 * 25),
		Match:      true,
		Filled:     true,
	}
	simulateLiquidateFromSellOrder(t, inputBook, inputOrder, expected)

	// Not filled and entirely liquidated (60 > 50)
	inputOrder = types.Order{Id: 10, Creator: MockAccount("1"), Amount: 60, Price: 10}
	expected = liquidateSellRes{
		Book: []types.Order{
			{Id: 2, Creator: MockAccount("2"), Amount: 30, Price: 15},
			{Id: 1, Creator: MockAccount("1"), Amount: 200, Price: 20},
		},
		Remaining:  types.Order{Id: 10, Creator: MockAccount("1"), Amount: 10, Price: 10},
		Liquidated: types.Order{Id: 0, Creator: MockAccount("0"), Amount: 50, Price: 25},
		Gain:       int32(50 * 25),
		Match:      true,
		Filled:     false,
	}
	simulateLiquidateFromSellOrder(t, inputBook, inputOrder, expected)
}

type fillSellRes struct {
	Book       []types.Order
	Remaining  types.Order
	Liquidated []types.Order
	Gain       int32
	Filled     bool
}

func simulateFillSellOrder(
	t *testing.T,
	inputList []types.Order,
	inputOrder types.Order,
	expected fillSellRes,
) {
	book := OrderListToBuyOrderBook(inputList)
	expectedBook := OrderListToBuyOrderBook(expected.Book)

	require.True(t, sort.SliceIsSorted(book.Book.Orders, func(i, j int) bool {
		return book.Book.Orders[i].Price < book.Book.Orders[j].Price
	}))
	require.True(t, sort.SliceIsSorted(expectedBook.Book.Orders, func(i, j int) bool {
		return expectedBook.Book.Orders[i].Price < expectedBook.Book.Orders[j].Price
	}))

	remaining, liquidated, gain, filled := book.FillSellOrder(inputOrder)

	require.Equal(t, expectedBook, book)
	require.Equal(t, expected.Remaining, remaining)
	require.Equal(t, expected.Liquidated, liquidated)
	require.Equal(t, expected.Gain, gain)
	require.Equal(t, expected.Filled, filled)
}

func TestFillSellOrder(t *testing.T) {
	var inputBook []types.Order

	// Empty book
	inputOrder := types.Order{Id: 10, Creator: MockAccount("1"), Amount: 30, Price: 30}
	expected := fillSellRes{
		Book:       []types.Order{},
		Remaining:  inputOrder,
		Liquidated: []types.Order(nil),
		Gain:       int32(0),
		Filled:     false,
	}
	simulateFillSellOrder(t, inputBook, inputOrder, expected)

	// No match
	inputBook = []types.Order{
		{Id: 2, Creator: MockAccount("2"), Amount: 30, Price: 15},
		{Id: 1, Creator: MockAccount("1"), Amount: 200, Price: 20},
		{Id: 0, Creator: MockAccount("0"), Amount: 50, Price: 25},
	}
	expected = fillSellRes{
		Book:       inputBook,
		Remaining:  inputOrder,
		Liquidated: []types.Order(nil),
		Gain:       int32(0),
		Filled:     false,
	}
	simulateFillSellOrder(t, inputBook, inputOrder, expected)

	// First order liquidated, not filled
	inputOrder = types.Order{Id: 10, Creator: MockAccount("1"), Amount: 60, Price: 22}
	expected = fillSellRes{
		Book: []types.Order{
			{Id: 2, Creator: MockAccount("2"), Amount: 30, Price: 15},
			{Id: 1, Creator: MockAccount("1"), Amount: 200, Price: 20},
		},
		Remaining: types.Order{Id: 10, Creator: MockAccount("1"), Amount: 10, Price: 22},
		Liquidated: []types.Order{
			{Id: 0, Creator: MockAccount("0"), Amount: 50, Price: 25},
		},
		Gain:   int32(50 * 25),
		Filled: false,
	}
	simulateFillSellOrder(t, inputBook, inputOrder, expected)

	// Filled with two order
	inputOrder = types.Order{Id: 10, Creator: MockAccount("1"), Amount: 60, Price: 18}
	expected = fillSellRes{
		Book: []types.Order{
			{Id: 2, Creator: MockAccount("2"), Amount: 30, Price: 15},
			{Id: 1, Creator: MockAccount("1"), Amount: 190, Price: 20},
		},
		Remaining: types.Order{Id: 10, Creator: MockAccount("1"), Amount: 0, Price: 18},
		Liquidated: []types.Order{
			{Id: 0, Creator: MockAccount("0"), Amount: 50, Price: 25},
			{Id: 1, Creator: MockAccount("1"), Amount: 10, Price: 20},
		},
		Gain:   int32(50*25 + 10*20),
		Filled: true,
	}
	simulateFillSellOrder(t, inputBook, inputOrder, expected)

	// Not filled, buy order book liquidated
	inputOrder = types.Order{Id: 10, Creator: MockAccount("1"), Amount: 300, Price: 10}
	expected = fillSellRes{
		Book:      []types.Order{},
		Remaining: types.Order{Id: 10, Creator: MockAccount("1"), Amount: 20, Price: 10},
		Liquidated: []types.Order{
			{Id: 0, Creator: MockAccount("0"), Amount: 50, Price: 25},
			{Id: 1, Creator: MockAccount("1"), Amount: 200, Price: 20},
			{Id: 2, Creator: MockAccount("2"), Amount: 30, Price: 15},
		},
		Gain:   int32(50*25 + 200*20 + 30*15),
		Filled: false,
	}
	simulateFillSellOrder(t, inputBook, inputOrder, expected)
}
