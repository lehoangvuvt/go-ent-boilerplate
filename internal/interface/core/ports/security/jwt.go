package securityports

import "github.com/golang-jwt/jwt/v5"

// JWTService defines signing and verification for JWT tokens.
type JWTService interface {
	Generate(claims jwt.Claims) (string, error)
	Verify(token string, claims jwt.Claims) error
}
