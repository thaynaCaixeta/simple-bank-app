package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/thaynaCaixeta/simple-bank-app/db/util"
)

func createRandomAccount(t *testing.T) Account {
	args := CreateAccountParams{
		Owner:    util.GenerateRandomOwner(),
		Balance:  util.GenerateRandomMoney(),
		Currency: util.GenerateRandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, args.Owner, account.Owner)
	require.Equal(t, args.Balance, account.Balance)
	require.Equal(t, args.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	randomAccount := createRandomAccount(t)
	retrievedAccount, err := testQueries.GetAccount(context.Background(), randomAccount.ID)
	require.NoError(t, err)
	require.NotEmpty(t, retrievedAccount)

	require.Equal(t, randomAccount.ID, retrievedAccount.ID)
	require.Equal(t, randomAccount.Owner, retrievedAccount.Owner)
	require.Equal(t, randomAccount.Balance, retrievedAccount.Balance)
	require.Equal(t, randomAccount.Currency, retrievedAccount.Currency)
	require.WithinDuration(t, randomAccount.CreatedAt.Time, retrievedAccount.CreatedAt.Time, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	randomAccount := createRandomAccount(t)

	args := UpdateAccountParams{
		ID:      randomAccount.ID,
		Balance: util.GenerateRandomMoney(),
	}

	updatedAccount, err := testQueries.UpdateAccount(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, updatedAccount)

	require.Equal(t, randomAccount.ID, updatedAccount.ID)
	require.Equal(t, randomAccount.Owner, updatedAccount.Owner)
	require.Equal(t, args.Balance, updatedAccount.Balance)
	require.Equal(t, randomAccount.Currency, updatedAccount.Currency)
	require.WithinDuration(t, randomAccount.CreatedAt.Time, updatedAccount.CreatedAt.Time, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	randomAccount := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), randomAccount.ID)
	require.NoError(t, err)

	deletedAccount, err := testQueries.GetAccount(context.Background(), randomAccount.ID)
	require.Error(t, err)
	require.Empty(t, deletedAccount)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	args := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), args)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
