package bootstrapstack

import (
	repositoryports "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/core/ports/repository"
	httpuser "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/http/user"
	userusecase "github.com/lehoangvuvt/go-ent-boilerplate/internal/usecase/user"
)

type BuildUserStackArgs struct {
	UserRepository repositoryports.UserRepository
}

func BuildUserStack(args BuildUserStackArgs) *httpuser.UserHandler {
	createUserUC := userusecase.NewUserUsecase(args.UserRepository)
	handler := httpuser.NewUserHandler(httpuser.NewUserHandlerArgs{
		CreateUserUC: createUserUC,
	})
	return handler
}
