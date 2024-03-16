package db

import (
	"context"
	"testing"

	"github.com/rogsiel/quizzer/internal/util"
	"github.com/stretchr/testify/require"
)

func TestSignupTx(t *testing.T) {
	store := NewStore(testDB)
	n := 5

	errs := make(chan error)
	results := make(chan UserTxResult)

	for i := 0; i < n; i++ {
		go func ()  {
			credentials := util.RandUserData()
			account, err := store.SignupTx(context.Background(), SignupTxParams{
				Name: credentials.Name,
				Email: credentials.Email,
			})
			errs <- err
			results <- account
		}()
	}

	for i := 0; i < n; i++ {
		err := <- errs
		require.NoError(t, err)
		
		account := <- results
		require.NotEmpty(t, account)
		require.NotZero(t, account.User.ID)
		require.NotZero(t, account.User.CreatedAt)

		_, err = store.GetUser(context.Background(), account.User.ID)
		require.NoError(t, err)
	}
}

func TestGetUserTx(t *testing.T) {
	store := NewStore(testDB)
	n := 5

	errs := make(chan error)
	results := make(chan UserTxResult)

	for i := 0; i < n; i++ {
		go func () {
			user := createRandomUser(t)
			account, err := store.GetUserTx(context.Background(), GetUserTxParams{
				ID: user.ID,
			})
			errs <- err
			results <- account
		}()
	}

	for i := 0; i < n; i++ {
		err := <- errs
		require.NoError(t, err)

		account := <- results
		require.NotEmpty(t, account)
		require.NotZero(t, account.User.ID)
		require.NotZero(t, account.User.CreatedAt)
	}
}


func TestGetUsersTx(t *testing.T) {
	store := NewStore(testDB)
	n := 20
	

	for i := 0; i < n; i++ {
		_ = createRandomUser(t)
	}
	
	accounts, err := store.GetUsersTx(context.Background(),GetUsersTxParams{
		Limit: 10,
		Offset: 5,
	})
	require.NoError(t, err)
	
	accounts2, err := testQueries.GetUsers(context.Background(), GetUsersParams{
		Limit: 10,
		Offset: 5,
	})
	require.NoError(t, err)
	
	require.NotEmpty(t, accounts)
	require.Equal(t, accounts[len(accounts)-1], accounts2[len(accounts2)-1])

}

func TestUpdateUserTx(t *testing.T) {
	store := NewStore(testDB)
	n := 5

	errs := make(chan error)
	results := make(chan UserTxResult)
	newNames := make(chan string)

	for i := 0; i < n; i++ {
		go func () {
			user := createRandomUser(t)
			newName := util.RandString(4)
			account, err := store.UpdateUserTx(context.Background(), UpdateUserTxParams{
				ID: user.ID,
				NewName: newName,
			})
			errs <- err
			results <- account
			newNames <- newName
		}()
	}

	for i := 0; i < n; i++ {
		err := <- errs
		require.NoError(t, err)

		account := <- results
		newName := <- newNames
		require.NotEmpty(t, account)
		require.NotZero(t, account.User.ID)
		require.NotZero(t, account.User.CreatedAt)
		require.Equal(t, account.User.Name, newName)
	}
}

func TestDeleteUserTx(t *testing.T) {
	store := NewStore(testDB)
	n := 5

	errsFirst := make(chan error)
	errsSec := make(chan error)
	results := make(chan UserTxResult)

	for i := 0; i < n; i++ {
		go func () {
			user := createRandomUser(t)
			err := store.DeleteUserTx(context.Background(), DeleteUserTxParams(user.ID))

			errsFirst <- err

			account, err := store.GetUserTx(context.Background(), GetUserTxParams{
				ID: user.ID,
			})
			
			errsSec <- err
			results <- account
		}()
	}

	for i := 0; i < n; i++ {
		err := <- errsFirst
		require.NoError(t, err)

		err = <- errsSec
		account := <- results
		require.Error(t, err)
		require.Empty(t, account)
	}
}
