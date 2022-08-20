package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {

	password := RandomString(6)
	hashed_password1, err := HashedPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashed_password1)
	require.NotEqual(t, hashed_password1, password)

	checkPassword := MatchPassword(password, hashed_password1)
	require.NoError(t, checkPassword)

	wrongPassword := RandomString(6)
	errPassword := MatchPassword(wrongPassword, hashed_password1)
	require.EqualError(t, errPassword, bcrypt.ErrMismatchedHashAndPassword.Error())

	hashed_password2, err := HashedPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashed_password2)
	require.NotEqual(t, hashed_password2, hashed_password1)

}
