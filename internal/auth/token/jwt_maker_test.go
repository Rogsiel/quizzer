package token

import (
	"math/rand"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/require"
)

func randStr (length int) string {
    alphnum := "abcdefghijklmnopqrstuvwxyz"
    k := len(alphnum)
    rand_str := make([]byte, length)
    for i := 0; i < length; i++ {
	rand_str[i] = alphnum[rand.Intn(k)]
    }
    return string(rand_str)
}

func TestJWTMaker(t *testing.T) {
    maker, err := NewJWTMaker(randStr(32))
    require.NoError(t, err)
    username := randStr(rand.Intn(10))
    duration := time.Minute

    issuedAt := time.Now()
    expiredAt := issuedAt.Add(duration)

    token, err := maker.CreateToken(username, duration)
    require.NoError(t, err)
    require.NotEmpty(t, token)

    payload, err := maker.VerifyToken(token)
    require.NoError(t, err)
    require.NotEmpty(t, payload)

    require.NotZero(t, payload.ID)
    require.Equal(t, username, payload.UserName)
    require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
    require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestExpiredJWTToken(t *testing.T) {
    maker, err := NewJWTMaker(randStr(32))
    require.NoError(t, err)

    token, err := maker.CreateToken(randStr(rand.Intn(10)),-time.Minute)
    require.NoError(t, err)
    require.NotEmpty(t, token)

    payload, err := maker.VerifyToken(token)
    require.Error(t, err)
    require.Nil(t, payload)
    require.EqualError(t, err, ErrorExpiredToken.Error())
}

func TestInvalidJWTToken(t *testing.T) {
    payload, err := NewPayload(randStr(rand.Intn(10)), time.Minute)
    require.NoError(t, err)

    jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
    token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
    require.NoError(t, err)

    maker, err := NewJWTMaker(randStr(32))
    require.NoError(t, err)

    payload, err = maker.VerifyToken(token)
    require.Error(t, err)
    require.EqualError(t, err, ErrorInvalidToken.Error())
    require.Nil(t, payload)
}
