package httprouter

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	idempotencyports "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/core/ports/idempotency"
	httpauth "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/http/auth"
	httpmiddleware "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/http/middleware"
	httptransaction "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/http/transaction"
	httpuser "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/http/user"
)

type NewRouterArgs struct {
	UserHandler        *httpuser.UserHandler
	AuthHandler        *httpauth.AuthHandler
	TransactionHandler *httptransaction.TransactionHandler
	AuthMiddleware     *httpmiddleware.AuthMiddleware
	IdempotencyStore   idempotencyports.IdempotencyStore
}

func NewRouter(args NewRouterArgs) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("ok"))
	})

	r.Route("/api/v1", func(ur chi.Router) {
		httpuser.RegisterRoutes(ur, args.UserHandler)
		httpauth.RegisterRoutes(ur, args.AuthHandler)

		ur.Group(func(pr chi.Router) {

			pr.Use(args.AuthMiddleware.RequireJWT)

			pr.Use(httpmiddleware.NewIdempotencyMiddleware(
				args.IdempotencyStore,
			).Handler)

			httptransaction.RegisterRoutes(pr, args.TransactionHandler)
		})
	})

	return r
}
