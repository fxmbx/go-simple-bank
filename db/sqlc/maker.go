package db

import (
	"errors"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/google/uuid"
	"github.com/o1egl/paseto"
)

type Maker interface {
	CreateToken(username string, duration time.Duration) (string, error)

	VertifyToken(tokenString string) (*Payload, error)
}

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuaedAt time.Time `json:"issued_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

func CreatePayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	return &Payload{
		ID:        tokenID,
		Username:  username,
		IssuaedAt: time.Now(),
		ExpiresAt: time.Now().Add(duration),
	}, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiresAt) {
		return errors.New("Expired Token")
	}
	return nil
}

type pasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func NewPasetoMaker(symmetricKet []byte) (Maker, error) {
	if len(symmetricKet) > 0 && len(symmetricKet) != chacha20poly1305.KeySize {
		return nil, errors.New("Invalid symmetric key")
	}

	maker := &pasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: symmetricKet,
	}

	return maker, nil
}

func (maker *pasetoMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := CreatePayload(username, duration)
	if err != nil {
		return "", err
	}
	token, err := maker.paseto.Encrypt(maker.symmetricKey, payload, nil)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (maker *pasetoMaker) VertifyToken(tokenString string) (*Payload, error) {
	payload := &Payload{}
	err := maker.paseto.Decrypt(tokenString, maker.symmetricKey, payload, nil)
	if err != nil {
		return nil, err
	}

	if err := payload.Valid(); err != nil {
		return nil, err
	}
	return payload, nil
}
