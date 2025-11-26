# Go Ent Boilerplate

Minimal Go service skeleton using chi for HTTP routing, ent for data access, and a clean architecture split across domain/usecase/interface/infra layers.

## Layout
- `cmd/`: application entrypoints (`main.go` bootstraps the HTTP server).
- `internal/app/`: wiring (container) and server startup.
- `internal/domain/`: core entities and validation (`user`).
- `internal/usecase/`: business logic and DTOs (`user` create).
- `internal/interface/`: adapters (HTTP router/handlers) and ports (`core/ports`).
- `internal/infrastructure/`: ent client setup and repository implementations.
- `pkg/httpx/`: small HTTP helpers for JSON encoding/decoding.
- `ent/`: generated ent code and schemas.

## Quick start
1) Set environment/config for the database:
   - Driver example: `POSTGRES_DRIVER=postgres`
   - DSN example: `POSTGRES_DSN=postgres://user:pass@localhost:5432/dbname?sslmode=disable`
   - Note: the current `internal/app/container.go` uses hard-coded config; adjust it or export vars before running.
2) Run the service:
   ```bash
   go run ./cmd
   ```
   The server listens on `:8080`.

## HTTP surface
- `GET /health` → basic health check.
- `POST /api/v1/users/register` → create user (JSON: `email`, `password`).

## Architecture highlights
- Dependency direction flows inward: interface → usecase → domain; infrastructure implements ports.
- Ent client is built in `internal/app`, repositories translate to domain entities.
- Request validation happens in DTOs and domain entities; handlers stay thin and delegate to use cases.

## Development notes
- Generated ent code lives in `ent/`; run `go generate ./ent` (with ent installed) after schema changes.
- Tests are not yet included; add per-layer tests as you extend functionality.
