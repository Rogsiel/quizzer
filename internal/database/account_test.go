package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/rogsiel/quizzer/internal/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User{
	arg := CreateUserParams{
		Name:	util.RandUsername(),
		Email:	util.RandomEmail(),
	}
	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Name, user.Name)
	require.Equal(t, arg.Email, user.Email)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)

	return user
}
func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)


	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Name, user2.Name)
	require.Equal(t, user2.Email, user2.Email)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func TestUpdateUser(t *testing.T) {
	user1 := createRandomUser(t)

	arg := UpdateUserParams{
		ID: user1.ID,
		Name: util.RandUsername(),
	}
	user2, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	
	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, arg.Name, user2.Name)
	require.Equal(t, user1.Email, user2.Email)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func TestDeleteUser(t *testing.T) {
	user1 := createRandomUser(t)
	err := testQueries.DeleteUser(context.Background(), user1.ID)
	require.NoError(t, err)

	user2, err := testQueries.GetUser(context.Background(), user1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, user2)
}

func TestGetAllUsers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomUser(t)
	}
	arg := GetUsersParams{
		Limit: 5,
		Offset: 5,
	}
	users, err := testQueries.GetUsers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, users, 5)

	for _, user := range users {
		require.NotEmpty(t, user)
	}

}
