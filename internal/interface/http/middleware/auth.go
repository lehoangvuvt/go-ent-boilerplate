package httpmiddleware

import (
	"context"
	"net/http"
	"strings"

	securityports "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/core/ports/security"
	authusecasedto "github.com/lehoangvuvt/go-ent-boilerplate/internal/usecase/auth/dto"
)

type ctxKey string

const authClaimsCtxKey ctxKey = "authClaims"

// AuthMiddleware verifies JWT bearer tokens and attaches auth claims to the request context.
type AuthMiddleware struct {
	jwtService securityports.JWTService
}

func NewAuthMiddleware(jwtService securityports.JWTService) *AuthMiddleware {
	return &AuthMiddleware{jwtService: jwtService}
}

func (m *AuthMiddleware) RequireJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(strings.ToLower(authHeader), "bearer ") {
			http.Error(w, "missing bearer token", http.StatusUnauthorized)
			return
		}

		raw := strings.TrimSpace(authHeader[len("Bearer "):])
		if raw == "" {
			http.Error(w, "missing bearer token", http.StatusUnauthorized)
			return
		}

		claims := &authusecasedto.AuthClaims{}
		if err := m.jwtService.Verify(raw, claims); err != nil {
			http.Error(w, "invalid or expired token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), authClaimsCtxKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// FromContext extracts AuthClaims set by RequireJWT middleware.
func FromContext(ctx context.Context) (authusecasedto.AuthClaims, bool) {
	val := ctx.Value(authClaimsCtxKey)
	if claims, ok := val.(*authusecasedto.AuthClaims); ok && claims != nil {
		return *claims, true
	}
	return authusecasedto.AuthClaims{}, false
}
