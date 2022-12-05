package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"github.com/techschool/simplebank/util"
	"testing"
)

func CreateEntry(t *testing.T, args CreateEntryParams) Entry {
	entry, err := testQueries.CreateEntry(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.Equal(t, entry.AccountID, args.AccountID)
	require.Equal(t, entry.Amount, args.Amount)
	return entry
}

func CreateRandomEntry(t *testing.T) Entry {
	account := createRandomAccount(t)
	money := util.RandomMoney()
	args := CreateEntryParams{
		Amount:    money,
		AccountID: account.ID,
	}
	return CreateEntry(t, args)
}

func CreateMultipleRandomEntries(t *testing.T) int64 {
	account := createRandomAccount(t)
	money := util.RandomMoney()
	args := CreateEntryParams{
		Amount:    money,
		AccountID: account.ID,
	}

	for i := 0; i < 10; i++ {
		CreateEntry(t, args)
	}
	return account.ID
}

func TestCreateEntry(t *testing.T) {
	CreateRandomEntry(t)
}

func TestGetEntry(t *testing.T) {
	entry := CreateRandomEntry(t)
	res, err := testQueries.GetEntry(context.Background(), entry.ID)
	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.Equal(t, entry.AccountID, res.AccountID)
	require.Equal(t, entry.Amount, res.Amount)
}

func TestListEntries(t *testing.T) {
	account_id := CreateMultipleRandomEntries(t)

	args := ListEntriesParams{
		AccountID: account_id,
		Offset:    4,
		Limit:     5,
	}

	entries, err := testQueries.ListEntries(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, entries)

	for _, entry := range entries {
		require.NotEmpty(t, entry.ID)
		require.NotEmpty(t, entry.AccountID)
		require.NotEmpty(t, entry.Amount)
		require.NotEmpty(t, entry.CreatedAt)
	}
}
