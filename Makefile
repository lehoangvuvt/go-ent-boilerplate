APP_NAME = github.com/lehoangvuvt/go-ent-boilerplate
CMD_PATH = ./cmd
SERVER_BIN=bin/server
WORKER_BIN=bin/worker

.PHONY: help build run server worker docker up down logs

help:
	@echo "Targets:"
	@echo "  make build        - Build server & worker"
	@echo "  make server       - Run server locally"
	@echo "  make worker       - Run worker locally"
	@echo "  make docker       - Build Docker image"
	@echo "  make up           - Start docker-compose"
	@echo "  make down         - Stop docker-compose"
	@echo "  make logs         - View docker-compose logs"

# ==================
# LOCAL BUILD
# ==================

build:
	@echo "Building server..."
	@go build -o $(SERVER_BIN) ./cmd/server
	@echo "Building worker..."
	@go build -o $(WORKER_BIN) ./cmd/worker

server:
	@go run ./cmd/server

worker:
	@go run ./cmd/worker

# ==================
# DOCKER
# ==================

docker:
	docker build -t $(APP_NAME) .

up:
	docker-compose up -d

down:
	docker-compose down

logs:
	docker-compose logs -f --tail=200


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

wire-gen:
	wire ./internal/app

install-tools:
	go install github.com/swaggo/swag/cmd/swag@latest
	go install entgo.io/ent/cmd/ent
	go install github.com/google/wire/cmd/wire@latest