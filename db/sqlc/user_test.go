package db

import (
	"context"
	"testing"

	"github.com/husky_dusky/simplebank/util"
	"github.com/stretchr/testify/require"
)

// It creates a random account and returns it
func createRandomUser(t *testing.T) User {
	hashedPassword, err := util.HashPassword("secret")
	require.NoError(t, err)

	arg := CreateUserParams{
		Username:       util.RandomOwnerName(),
		HashedPassword: hashedPassword,
		FullName:       util.RandomOwnerName(),
		Email:          util.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)
	require.NotZero(t, user.HashedPassword)

	require.NotZero(t, user.CreatedAt)
	require.True(t, user.PasswordChangeAt.IsZero())

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomAccount(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)

	user2, err := testQueries.GetUser(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.PasswordChangeAt, user2.PasswordChangeAt)
	require.Equal(t, user1.CreatedAt, user2.CreatedAt)
}
