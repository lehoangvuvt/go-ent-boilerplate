package bootstrap

import (
	"time"

	"github.com/go-chi/chi/v5"
	bootstrapstack "github.com/lehoangvuvt/go-ent-boilerplate/internal/bootstrap/stack"
	cacheports "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/core/ports/cache"
	idempotencyports "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/core/ports/idempotency"
	mailports "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/core/ports/mail"
	repositoryports "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/core/ports/repository"
	securityports "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/core/ports/security"
	httpmiddleware "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/http/middleware"
	httprouter "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/http/router"
)

type HandlerBootstrapArgs struct {
	Repositories Repositories
	Services     Services
	Stores       Stores
}

type Repositories struct {
	UserRepository        repositoryports.UserRepository
	TransactionRepository repositoryports.TransactionRepository
}

type Services struct {
	JWTService   securityports.JWTService
	JWTDuration  time.Duration
	CacheService cacheports.Cache
	MailService  mailports.MailService
}

type Stores struct {
	IdempotencyStore idempotencyports.IdempotencyStore
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

	transactionHandler := bootstrapstack.BuildTransactionStack(bootstrapstack.BuildTransactionStackArgs{
		TransactionRepository: args.Repositories.TransactionRepository,
	})

	authMW := httpmiddleware.NewAuthMiddleware(args.Services.JWTService)

	router := httprouter.NewRouter(httprouter.NewRouterArgs{
		UserHandler:        userHandler,
		AuthHandler:        authHandler,
		TransactionHandler: transactionHandler,
		AuthMiddleware:     authMW,
		IdempotencyStore:   args.Stores.IdempotencyStore,
	})

	return router
}
