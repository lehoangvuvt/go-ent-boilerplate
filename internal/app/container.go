package app

import (
	"context"
	"log"

	"github.com/go-chi/chi/v5"
	entdb "github.com/lehoangvuvt/go-ent-boilerplate/internal/infrastructure/ent"
	userrepository "github.com/lehoangvuvt/go-ent-boilerplate/internal/infrastructure/repository/user"
	httprouter "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/http/router"
	httpuser "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/http/user"
	userusecase "github.com/lehoangvuvt/go-ent-boilerplate/internal/usecase/user"
)

type Container struct {
	Router *chi.Mux
}

func Build(ctx context.Context) *Container {
	entDB, err := entdb.New(ctx, entdb.Config{
		Driver:      "postgres",
		DSN:         "",
		AutoMigrate: true,
	})
	if err != nil {
		log.Fatalf("failed initialzing ent db. Error: %w", err)
	}

	userInfra := userrepository.NewUserRepository(entDB.Client())
	createUserUC := userusecase.NewUserUsecase(userInfra)
	userHandler := httpuser.NewUserHandler(httpuser.NewUserHandlerArgs{
		CreateUserUC: createUserUC,
	})

	router := httprouter.NewRouter(httprouter.NewRouterArgs{
		UserHandler: userHandler,
	})

	return &Container{
		Router: router,
	}
}
