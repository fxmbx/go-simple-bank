package db

import (
	"context"
	"testing"
	"time"

	"github.com/fxmbx/go-simple-bank/utils"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T) Entry {

	account := createRandomAccount(t)
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    utils.RandomMoney(),
	}
	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	return entry
}

func TestCreatEntry(t *testing.T) {
	createRandomEntry(t)
}

func TestGetEntryById(t *testing.T) {
	entry1 := createRandomEntry(t)
	entry2, err := testQueries.GetEntryByID(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestListEntries(t *testing.T) {
	for i := 0; i < 5; i++ {
		createRandomEntry(t)
	}
	arg := ListEntriesParams{
		Limit:  2,
		Offset: 2,
	}
	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 2)
	for _, v := range entries {
		require.NotEmpty(t, v)
	}
}

func TestListEntriesByAccountID(t *testing.T) {
	entry1 := createRandomEntry(t)

	arg := ListEntriesyByAccountIDParams{
		AccountID: entry1.AccountID,
		Limit:     1,
		Offset:    1,
	}
	entries, err := testQueries.ListEntriesyByAccountID(context.Background(), arg)
	require.NoError(t, err)
	for _, v := range entries {
		require.NotEmpty(t, v)
	}
}
