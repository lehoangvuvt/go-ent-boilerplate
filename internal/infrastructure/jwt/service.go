package jwtinfra

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	securityports "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/core/ports/security"
	"github.com/lehoangvuvt/go-ent-boilerplate/pkg/jwtx"
)

type Service struct {
	client *jwtx.Client
}

func NewService(secretKey string, tokenDuration time.Duration) *Service {
	return &Service{
		client: jwtx.NewJWTClient(secretKey, tokenDuration),
	}
}

func (s *Service) Generate(claims jwt.Claims) (string, error) {
	return s.client.Generate(claims)
}

func (s *Service) Verify(token string, claims jwt.Claims) error {
	return s.client.Verify(token, claims)
}

var _ securityports.JWTService = (*Service)(nil)
