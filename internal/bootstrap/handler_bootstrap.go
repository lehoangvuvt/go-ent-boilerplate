package bootstrap

import (
	"github.com/go-chi/chi/v5"
	bootstrapstack "github.com/lehoangvuvt/go-ent-boilerplate/internal/bootstrap/stack"
	repositoryports "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/core/ports/repository"
	httprouter "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/http/router"
)

type HandlerBootstrapArgs struct {
	Repositories Repositories
}

type Repositories struct {
	UserRepository repositoryports.UserRepository
}

func BootstrapHandler(args HandlerBootstrapArgs) *chi.Mux {
	userHandler := bootstrapstack.BuildUserStack(bootstrapstack.BuildUserStackArgs{
		UserRepository: args.Repositories.UserRepository,
	})

	authHandler := bootstrapstack.BuildAuthStack(bootstrapstack.BuildAuthStackArgs{
		UserRepository: args.Repositories.UserRepository,
	})

	router := httprouter.NewRouter(httprouter.NewRouterArgs{
		UserHandler: userHandler,
		AuthHandler: authHandler,
	})

	return router
}
