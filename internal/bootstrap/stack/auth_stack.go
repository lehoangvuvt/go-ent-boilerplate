package bootstrapstack

import (
	"time"

	repositoryports "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/core/ports/repository"
	securityports "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/core/ports/security"
	httpauth "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/http/auth"
	authusecase "github.com/lehoangvuvt/go-ent-boilerplate/internal/usecase/auth"
)

type BuildAuthStackArgs struct {
	UserRepository repositoryports.UserRepository
	JWTService     securityports.JWTService
	TokenDuration  time.Duration
}

func BuildAuthStack(args BuildAuthStackArgs) *httpauth.AuthHandler {
	loginUC := authusecase.NewLoginUsercase(args.UserRepository, args.JWTService, args.TokenDuration)
	authHandler := httpauth.NewAuthHandler(httpauth.NewAuthHandlerArgs{
		LoginUC: loginUC,
	})
	return authHandler
}
