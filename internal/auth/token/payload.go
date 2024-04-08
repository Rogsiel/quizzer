package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
    ErrorInvalidToken = errors.New("Token is invalid")
    ErrorExpiredToken = errors.New("this token has expired")
)
type Payload struct {
    ID		uuid.UUID   `json:"id"`
    UserName    string	    `json:"name"`
    IssuedAt	time.Time   `json:"issued_at"`
    ExpiredAt	time.Time   `json:"expired_at"`
}

func NewPayload(username string, duration time.Duration) (*Payload, error) {
    tokenID, err := uuid.NewRandom()
    if err != nil {
	return nil, err
    }
    
    payload := &Payload{
	ID: tokenID,
	UserName: username,
	IssuedAt: time.Now(),
	ExpiredAt: time.Now().Add(duration),
    }

    return payload, nil
}

func (payload *Payload) Valid() error {
    if time.Now().After(payload.ExpiredAt) {
	return ErrorExpiredToken
    }
    return nil
}
