APP_NAME = github.com/lehoangvuvt/go-ent-boilerplate
CMD_PATH = ./cmd

fmt: 
	go fmt ./...

lint: 
	go vet ./...

run:
	go run $(CMD_PATH)/main.go

build:
	go build -o bin/$(APP_NAME) $(CMD_PATH)

ent-create:
	go run -mod=mod entgo.io/ent/cmd/ent new ${name}

ent-gen:
	go generate ./ent

install-tools:
	go install github.com/swaggo/swag/cmd/swag@latest
	go get entgo.io/ent/cmd/ent