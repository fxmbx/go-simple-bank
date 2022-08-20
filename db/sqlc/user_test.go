package db

import (
	"context"
	"testing"
	"time"

	"github.com/fxmbx/go-simple-bank/utils"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	hashedPassword, err := utils.HashedPassword(utils.RandomString(6))
	require.NoError(t, err)
	arg := CreateUserParams{
		Username:       utils.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       utils.RandomOwner(),
		Email:          utils.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, user.FullName, arg.FullName)
	require.Equal(t, user.Username, arg.Username)
	require.Equal(t, user.Email, arg.Email)
	require.NotZero(t, user.CreatedAt)
	require.True(t, user.PasswordChangedAt.IsZero())
	return user
}
func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUserByUserName(t *testing.T) {
	user1 := createRandomUser(t)

	user, err := testQueries.GetUser(context.Background(), user1.Username)
	require.NoError(t, err)
	require.Equal(t, user.Username, user1.Username)
	require.WithinDuration(t, user.PasswordChangedAt, user1.PasswordChangedAt, time.Second)
}
