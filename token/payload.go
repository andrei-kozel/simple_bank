package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrExpiredToken    = errors.New("token has expired")
	ErrInvalidToken    = errors.New("invalid token")
	ErrInvalidKey      = errors.New("invalid key")
	ErrInvalidIssuer   = errors.New("invalid issuer")
	ErrInvalidAudience = errors.New("invalid audience")
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  int64     `json:"issued_at"`
	ExpiredAt int64     `json:"expired_at"`
}

func NewPayload(username string, duration time.Duration) (*Payload, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        id,
		Username:  username,
		IssuedAt:  time.Now().Unix(),
		ExpiredAt: time.Now().Add(duration).Unix(),
	}

	return payload, nil
}

func (p *Payload) Valid() error {
	if time.Now().After(time.Unix(p.ExpiredAt, 0)) {
		return ErrExpiredToken
	}
	return nil
}
