package main

import (
	"context"
	"log"
	"net/http"

	entdb "github.com/lehoangvuvt/go-ent-boilerplate/internal/infrastructure/ent"
	userrepository "github.com/lehoangvuvt/go-ent-boilerplate/internal/infrastructure/repository/user"
	httprouter "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/http/router"
	httpuser "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/http/user"
	userusecase "github.com/lehoangvuvt/go-ent-boilerplate/internal/usecase/user"
	_ "github.com/lib/pq"
)

func main() {
	ctx := context.Background()
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

	r := httprouter.NewRouter(httprouter.NewRouterArgs{
		UserHandler: userHandler,
	})

	log.Println("listening on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
