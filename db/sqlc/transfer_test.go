package db

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"github.com/techschool/simplebank/util"
	"testing"
)

func createRandomTransfer(t *testing.T, account1 Account, account2 Account) Transfer {
	args := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), args)
	require.NoError(t, err)
	require.Equal(t, transfer.FromAccountID, account1.ID)
	require.Equal(t, transfer.ToAccountID, account2.ID)
	require.NotEmpty(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	createRandomTransfer(t, account1, account2)
}

func TestGetTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	transfer1 := createRandomTransfer(t, account1, account2)

	res, err := testQueries.GetTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)
	require.Equal(t, res.ID, transfer1.ID)
	require.Equal(t, res.FromAccountID, transfer1.FromAccountID)
	require.Equal(t, res.ToAccountID, transfer1.ToAccountID)
	require.Equal(t, res.Amount, transfer1.Amount)
	require.Equal(t, res.CreatedAt, transfer1.CreatedAt)

	transfer2 := createRandomTransfer(t, account1, account2)
	res2, err := testQueries.GetTransfer(context.Background(), transfer2.ID)
	require.NoError(t, err)
	require.Equal(t, res2.ID, transfer2.ID)
	require.NotEqual(t, transfer1.ID, transfer2.ID)
}

func TestListTransfers(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	for i := 0; i < 5; i++ {
		createRandomTransfer(t, account1, account2)
	}

	args := ListTransfersParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Limit:         5,
		Offset:        2,
	}

	res, err := testQueries.ListTransfers(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, res)

	for _, trans := range res {
		fmt.Print("meow")
		require.NotEmpty(t, trans.ID)
		require.NotEmpty(t, trans.FromAccountID)
		require.NotEmpty(t, trans.ToAccountID)
		require.NotEmpty(t, trans.Amount)
		require.NotEmpty(t, trans.CreatedAt)
	}
}
