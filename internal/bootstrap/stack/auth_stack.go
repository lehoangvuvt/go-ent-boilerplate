package bootstrapstack

import (
	repositoryports "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/core/ports/repository"
	httpauth "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/http/auth"
	authusecase "github.com/lehoangvuvt/go-ent-boilerplate/internal/usecase/auth"
)

type BuildAuthStackArgs struct {
	UserRepository repositoryports.UserRepository
}

func BuildAuthStack(args BuildAuthStackArgs) *httpauth.AuthHandler {
	loginUC := authusecase.NewLoginUsercase(args.UserRepository)
	authHandler := httpauth.NewAuthHandler(httpauth.NewAuthHandlerArgs{
		LoginUC: loginUC,
	})
	return authHandler
}
