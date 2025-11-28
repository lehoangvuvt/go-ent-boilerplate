package authusecasedto

import "github.com/golang-jwt/jwt/v5"

type AuthClaims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

type LoginResponse struct {
	Token string             `json:"token"`
	User  *UserResponseBrief `json:"user"`
}

type UserResponseBrief struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}
