package types_test

import (
	"math/rand"
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"ollo/x/dex/types"
)

func GenString(n int) string {
	alpha := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	buf := make([]rune, n)
	for i := range buf {
		buf[i] = alpha[rand.Intn(len(alpha))]
	}

	return string(buf)
}

func GenAddress() string {
	pk := ed25519.GenPrivKey().PubKey()
	addr := pk.Address()
	return sdk.AccAddress(addr).String()
}

func GenAmount() int32 {
	return int32(rand.Intn(int(types.MaxAmount)) + 1)
}

func GenPrice() int32 {
	return int32(rand.Intn(int(types.MaxPrice)) + 1)
}

func GenPair() (string, string) {
	return GenString(10), GenString(10)
}

func GenOrder() (string, int32, int32) {
	return GenLocalAccount(), GenAmount(), GenPrice()
}

func GenLocalAccount() string {
	return GenAddress()
}

func MockAccount(str string) string {
	return str
}

func OrderListToOrderBook(list []types.Order) types.OrderBook {
	listCopy := make([]*types.Order, len(list))
	for i, order := range list {
		order := order
		listCopy[i] = &order
	}

	return types.OrderBook{
		IdCount: 0,
		Orders:  listCopy,
	}
}

func TestRemoveOrderFromID(t *testing.T) {
	inputList := []types.Order{
		{Id: 3, Creator: MockAccount("3"), Amount: 2, Price: 10},
		{Id: 2, Creator: MockAccount("2"), Amount: 30, Price: 15},
		{Id: 1, Creator: MockAccount("1"), Amount: 200, Price: 20},
		{Id: 0, Creator: MockAccount("0"), Amount: 50, Price: 25},
	}

	book := OrderListToOrderBook(inputList)
	expectedList := []types.Order{
		{Id: 3, Creator: MockAccount("3"), Amount: 2, Price: 10},
		{Id: 1, Creator: MockAccount("1"), Amount: 200, Price: 20},
		{Id: 0, Creator: MockAccount("0"), Amount: 50, Price: 25},
	}
	expectedBook := OrderListToOrderBook(expectedList)
	err := book.RemoveOrderFromID(2)
	require.NoError(t, err)
	require.Equal(t, expectedBook, book)

	book = OrderListToOrderBook(inputList)
	expectedList = []types.Order{
		{Id: 3, Creator: MockAccount("3"), Amount: 2, Price: 10},
		{Id: 2, Creator: MockAccount("2"), Amount: 30, Price: 15},
		{Id: 1, Creator: MockAccount("1"), Amount: 200, Price: 20},
	}
	expectedBook = OrderListToOrderBook(expectedList)
	err = book.RemoveOrderFromID(0)
	require.NoError(t, err)
	require.Equal(t, expectedBook, book)

	book = OrderListToOrderBook(inputList)
	expectedList = []types.Order{
		{Id: 2, Creator: MockAccount("2"), Amount: 30, Price: 15},
		{Id: 1, Creator: MockAccount("1"), Amount: 200, Price: 20},
		{Id: 0, Creator: MockAccount("0"), Amount: 50, Price: 25},
	}
	expectedBook = OrderListToOrderBook(expectedList)
	err = book.RemoveOrderFromID(3)
	require.NoError(t, err)
	require.Equal(t, expectedBook, book)

	book = OrderListToOrderBook(inputList)
	err = book.RemoveOrderFromID(4)
	require.ErrorIs(t, err, types.ErrOrderNotFound)
}
