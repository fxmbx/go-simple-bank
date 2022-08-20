package token

import (
	"testing"
	"time"

	"github.com/fxmbx/go-simple-bank/utils"
	"github.com/stretchr/testify/require"
)

func TestCreatePasetoToken(t *testing.T) {
	symmetricKey := utils.RandomString(32)
	maker, err := NewPasetoMaker(symmetricKey)
	require.NoError(t, err)
	require.NotEmpty(t, maker)
	issuedAt := time.Now()
	expiredAt := issuedAt.Add(time.Minute)
	username := utils.RandomOwner()
	token, err := maker.CreateToken(username, time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.WithinDuration(t, payload.ExpiredAt, expiredAt, time.Second)
	require.Equal(t, payload.Username, username)
	require.NotZero(t, payload.ID)

}
