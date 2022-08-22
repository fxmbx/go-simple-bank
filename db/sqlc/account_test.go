package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/fxmbx/go-simple-bank/utils"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	owner := createRandomUser(t)
	arg := CreateAccountParams{
		Owner:    owner.Username,
		Balance:  utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Owner, arg.Owner)
	require.Equal(t, arg.Balance, arg.Balance)
	require.Equal(t, arg.Currency, arg.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
	return account
}
func TestCreatAccount(t *testing.T) {
	owner := createRandomUser(t)

	arg := CreateAccountParams{
		Owner:    owner.Username,
		Balance:  utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Owner, arg.Owner)
	require.Equal(t, arg.Balance, arg.Balance)
	require.Equal(t, arg.Currency, arg.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}

func TestGetAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	account2, err := testQueries.GetAccountByID(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	arg := UpdateAccountParams{
		ID:      account1.ID,
		Balance: utils.RandomMoney(),
	}
	account, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, account1.ID, account.ID)
	require.Equal(t, account1.Owner, account.Owner)
	require.Equal(t, account1.Currency, account.Currency)
	require.Equal(t, account.Balance, arg.Balance)
	require.NotEmpty(t, account)
}

func TestDeleteAcco(t *testing.T) {
	account1 := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	account2, err := testQueries.GetAccountByID(context.Background(), account1.ID)
	require.Error(t, err)
	require.Empty(t, account2)
	require.EqualError(t, err, sql.ErrNoRows.Error())

}
func TestListAccount(t *testing.T) {
	var lastAccount Account
	for i := 0; i <= 10; i++ {
		lastAccount = createRandomAccount(t)
	}
	arg := ListAccountsParams{
		Owner:  lastAccount.Owner,
		Limit:  5,
		Offset: 0,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)
	for _, v := range accounts {
		require.NotEmpty(t, v)
		require.Equal(t, v.Owner, lastAccount.Owner)
		// require.NotEmpty(t, v.Balance)
		// require.NotEmpty(t, v.CreatedAt)
		// require.NotEmpty(t, v.Currency)

	}
}
