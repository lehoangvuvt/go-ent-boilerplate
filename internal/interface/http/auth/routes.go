package httpauth

import (
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r chi.Router, h *AuthHandler) {
	r.Route("/auth", func(ur chi.Router) {
		ur.Post("/login", h.Login)
	})
}
