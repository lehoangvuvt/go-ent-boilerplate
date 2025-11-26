package main

import (
	"context"

	"github.com/lehoangvuvt/go-ent-boilerplate/internal/app"
	_ "github.com/lib/pq"
)

func main() {
	ctx := context.Background()
	c := app.Build(ctx)
	app.StartServer(c.Router)
}
