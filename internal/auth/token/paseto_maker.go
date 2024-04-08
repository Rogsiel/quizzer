package token

import (
	"fmt"
	"time"

	"aidanwoods.dev/go-paseto"
	"github.com/google/uuid"
	"golang.org/x/crypto/chacha20poly1305"
)

type PasetoMaker struct {
    paseto	    paseto.Token
    symmetricKey    paseto.V4SymmetricKey
}

func NewPasetoMaker(byteKey string) (Maker, error) {
    if len(byteKey) != chacha20poly1305.KeySize {
	return nil, fmt.Errorf("Invalid key: Key size must exactly be %d characters", chacha20poly1305.KeySize)
    }
    key, err := paseto.V4SymmetricKeyFromBytes([]byte(byteKey))
    if err != nil {
	return nil, err
    }
    
    maker := &PasetoMaker{
	paseto: paseto.NewToken(),
	symmetricKey: key,
    }

    return maker, nil
}

func (maker *PasetoMaker) CreateToken(username string, duration time.Duration) (string, error){
    payload, err := NewPayload(username, duration)
    if err != nil {
	return "", err
    }
    maker.paseto.SetIssuedAt(payload.IssuedAt)
    maker.paseto.SetExpiration(payload.ExpiredAt)
    maker.paseto.SetString("id", payload.ID.String())
    maker.paseto.SetString("name", payload.UserName)
    return maker.paseto.V4Encrypt(maker.symmetricKey, nil), nil
}
func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error){
    parser := paseto.NewParser()
    parser.AddRule(paseto.NotExpired())
    
    payload := &Payload{}
    
    parsedToken, err := parser.ParseV4Local(maker.symmetricKey, token, nil)
    if err != nil {
	return nil, err
    }
    
    id, err := parsedToken.GetString("id")
    if err != nil {
    	return nil, ErrorInvalidToken
    }
    payload.ID, err = uuid.Parse(id)
    if err != nil {
    	return nil, ErrorInvalidToken
    }
    payload.UserName, err = parsedToken.GetString("name")
    if err != nil {
    	return nil, ErrorInvalidToken
    }
    payload.IssuedAt, err = parsedToken.GetIssuedAt()
    if err != nil {
    	return nil, ErrorInvalidToken
    }
    payload.ExpiredAt, err = parsedToken.GetExpiration()
    if err != nil {
    	return nil, ErrorInvalidToken
    }

    err = payload.Valid()
    if err != nil {
	return nil, err
    }
    return payload, nil
}
