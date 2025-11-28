package jwtx

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrInvalidToken = errors.New("invalid token")
)

type Client struct {
	secretKey     string
	tokenDuration time.Duration
}

func NewJWTClient(secretKey string, tokenDuration time.Duration) *Client {
	return &Client{
		secretKey,
		tokenDuration,
	}
}

func (c *Client) Generate(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(c.secretKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func (c *Client) Verify(tokenStr string, claims jwt.Claims) error {
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return []byte(c.secretKey), nil
	})
	if err != nil || !token.Valid {
		return ErrInvalidToken
	}
	return nil
}
