package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

const minSecretKeySize = 32

type JWTMaker struct {
    secretKey	string
}

func NewJWTMaker(secretKey string) (Maker, error) {
    if len(secretKey) < minSecretKeySize {
	return nil, fmt.Errorf("Invalid key: Key must be at least %d characters", minSecretKeySize)
    }
    return &JWTMaker{secretKey}, nil
}

func (maker *JWTMaker) CreateToken(username string, duration time.Duration) (string, *Payload, error){
    payload, err := NewPayload(username, duration)
    if err != nil {
    	return "", payload, err
    }
    jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
    token, err := jwtToken.SignedString([]byte(maker.secretKey))
    return token, payload, err
}

func (maker *JWTMaker) VerifyToken(token string) (*Payload, error){
    keyFunc := func(token *jwt.Token) (any, error) {
	_, ok := token.Method.(*jwt.SigningMethodHMAC)
	if !ok {
	    return nil, ErrorInvalidToken
	}
	return []byte(maker.secretKey), nil
    }

    jwtToken, err := jwt.ParseWithClaims(token,&Payload{}, keyFunc)
    if err != nil {
	ver, ok := err.(*jwt.ValidationError)
	if ok && errors.Is(ver.Inner, ErrorExpiredToken) {
	    return nil, ErrorExpiredToken
	}
	return nil, ErrorInvalidToken
    }

    payload, ok := jwtToken.Claims.(*Payload)
    if !ok {
	return nil, ErrorInvalidToken
    }
    return payload, nil
}
