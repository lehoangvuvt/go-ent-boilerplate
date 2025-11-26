package httpuser

import (
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r chi.Router, h *UserHandler) {
	r.Route("/users", func(ur chi.Router) {
		ur.Post("/register", h.CreateUser)
	})
}
