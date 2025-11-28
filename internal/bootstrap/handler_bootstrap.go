package bootstrap

import (
	"time"

	"github.com/go-chi/chi/v5"
	bootstrapstack "github.com/lehoangvuvt/go-ent-boilerplate/internal/bootstrap/stack"
	cacheports "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/core/ports/cache"
	repositoryports "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/core/ports/repository"
	securityports "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/core/ports/security"
	httprouter "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/http/router"
)

type HandlerBootstrapArgs struct {
	Repositories Repositories
	Services     Services
}

type Repositories struct {
	UserRepository repositoryports.UserRepository
}

type Services struct {
	JWTService   securityports.JWTService
	JWTDuration  time.Duration
	CacheService cacheports.Cache
}

func BootstrapHandler(args HandlerBootstrapArgs) *chi.Mux {
	userHandler := bootstrapstack.BuildUserStack(bootstrapstack.BuildUserStackArgs{
		UserRepository: args.Repositories.UserRepository,
	})

	authHandler := bootstrapstack.BuildAuthStack(bootstrapstack.BuildAuthStackArgs{
		UserRepository: args.Repositories.UserRepository,
		JWTService:     args.Services.JWTService,
		TokenDuration:  args.Services.JWTDuration,
	})

	router := httprouter.NewRouter(httprouter.NewRouterArgs{
		UserHandler: userHandler,
		AuthHandler: authHandler,
	})

	return router
}
