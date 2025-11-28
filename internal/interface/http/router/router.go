package httprouter

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpauth "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/http/auth"
	httpuser "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/http/user"
)

type NewRouterArgs struct {
	UserHandler *httpuser.UserHandler
	AuthHandler *httpauth.AuthHandler
}

func NewRouter(args NewRouterArgs) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	r.Route("/api/v1", func(ur chi.Router) {
		httpuser.RegisterRoutes(ur, args.UserHandler)
		httpauth.RegisterRoutes(ur, args.AuthHandler)
	})

	return r
}
