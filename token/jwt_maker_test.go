package token

import (
	"testing"
	"time"

	"github.com/husky_dusky/simplebank/util"
	"github.com/stretchr/testify/require"
)

func TestJWTMaker(t *testing.T) {
	maker, err := NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)

	username := util.RandomOwnerName()
	duration := time.Minute

	issue := time.Now()
	expitedAt := issue.Add(duration)

	token, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, username, payload.Username)
	require.WithinDuration(t, issue, time.Unix(payload.IssuedAt, 0), time.Second)
	require.WithinDuration(t, expitedAt, time.Unix(payload.ExpiredAt, 0), time.Second)
}

func TestExpiredJWTToken(t *testing.T) {
	maker, err := NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)

	token, err := maker.CreateToken(util.RandomOwnerName(), -time.Minute)
	require.NoError(t, err)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.Nil(t, payload)

	require.EqualError(t, err, ErrExpiredToken.Error())
}

func TestInvalidJWTToken(t *testing.T) {
	maker, err := NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)

	token, err := maker.CreateToken(util.RandomOwnerName(), time.Minute)
	require.NoError(t, err)

	payload, err := maker.VerifyToken(token + "invalid")
	require.Error(t, err)
	require.Nil(t, payload)

	require.EqualError(t, err, ErrInvalidToken.Error())
}
