# Go Ent Clean/Hex Boilerplate

Minimal Go HTTP service built with chi and ent, structured around Clean Architecture and hexagonal boundaries (domain + use cases at the center, interfaces/ports at the edges, adapters plugged in from infrastructure).

## Why this structure
- Clear direction of dependencies: handlers -> use cases -> ports -> domain, with infrastructure implementing ports and injected inward.
- Composition root lives at the boundary (`internal/app` + `internal/bootstrap`), keeping wiring separate from core logic.
- Domain owns business rules (entities, invariants), use cases orchestrate workflows, interfaces describe contracts, and adapters translate to real IO (DB, JWT, cache).

## Project layout (goal of each folder)
- `cmd/`: Entrypoint (`main.go`); load config, build container, start server with graceful shutdown.
- `internal/app/`: Composition root; bootstraps ent, Redis cache, JWT service, repositories; exposes router and DB lifecycle.
- `internal/bootstrap/`: Wiring helpers (build DSN, create stacks). `stack/` assembles feature stacks (auth, user, transaction) and router routes.
- `internal/config/`: Config loading/validation via Viper and env; central config structs.
- `internal/domain/`: Core business model. `user/` validation; `transaction/` aggregate state machine and payment method value objects + factories.
- `internal/usecase/`: Application services. Auth login; user creation; transaction creation/list/find and state transitions (confirm/cancel/fail); DTOs per feature.
- `internal/interface/core/ports/`: Boundary interfaces (repositories, JWT service, cache) that use cases depend on.
- `internal/interface/http/`: HTTP adapters. Router setup, auth/user/transaction handlers, request parsing, response shaping.
- `internal/infrastructure/`: Outbound adapters. `ent/` client setup; repositories for user/transaction; `jwt/` wrapping `jwtx`; `cache/redis` implementing cache port.
- `ent/`: Ent schemas (`user`, `post`, `transaction`) and generated ORM code (run `go generate ./ent` after schema changes).
- `pkg/`: Shared utilities (`httpx` JSON helpers, `jwtx` thin JWT client, `logger` setup).
- `Makefile`: fmt/lint/run/build targets and ent helper commands.
- `.env`, `.gitignore`, `go.mod`, `go.sum`: environment sample, ignores, module metadata.

## How the layers collaborate
1) `cmd/main.go` loads configuration and calls `internal/app.Build`, which bootstraps infrastructure (ent DB, Redis, JWT) and constructs repositories.
2) `internal/bootstrap` builds handler stacks: use cases receive port interfaces, handlers receive use cases, and the chi router registers routes.
3) HTTP handlers only parse/validate transport payloads and delegate to use cases.
4) Use cases enforce business rules and call ports; repositories and services live behind those ports.
5) Infrastructure adapters translate ports to real implementations (ent persistence, JWT signing/verification, Redis cache).

Clean rules: inner layers never import outer ones; adapters depend on interfaces, not the other way around. Ports keep IO replaceable for tests and alternative adapters.

## Domain highlights
- `internal/domain/user`: entity validation (email, hashed password required).
- `internal/domain/transaction`: transaction aggregate with states (`pending`, `completed`, `failed`, `rejected`) and behaviors (`Confirm`, `Cancel`, `Fail`). Payment methods captured as value objects (visa, banking, ewallet, qr) with factories for creation.

## Use cases
- Auth: `LoginUsecase` verifies credentials, produces JWT with claims and expiry.
- User: `CreateUserUsecase` hashes password, validates entity, persists through `UserRepository`.
- Transaction: create/list/find plus state transitions (`Confirm`, `Cancel`, `Fail`) executed via repository updates.

## Ports and adapters
- Repository ports: `UserRepository`, `TransactionRepository` with filtering and lookups. Implemented by ent-backed adapters under `internal/infrastructure/repository/...`.
- Security port: `JWTService`, implemented by `internal/infrastructure/jwt.Service` wrapping `pkg/jwtx`.
- Cache port: Redis adapter under `internal/infrastructure/cache/redis` with a simple key/value TTL API.

## HTTP surface
- `GET /health`
- `POST /api/v1/users/register` — create user (body: `email`, `password`)
- `POST /api/v1/auth/login` — returns `{ token, user }`
- `POST /api/v1/transactions` — create transaction
- `GET /api/v1/transactions` — list with optional filters (`user_id`, `status`, `method`, `limit`, `offset`)
- `GET /api/v1/transactions/{id}` — fetch one
- `POST /api/v1/transactions/{id}/confirm` — confirm
- `POST /api/v1/transactions/{id}/cancel` — cancel

## Configuration
Environment variables (Viper maps `.` to `_`; `.env` is loaded first):
- `APP_PORT`
- `DB_HOST`, `DB_PORT`, `DB_NAME`, `DB_USER`, `DB_PASSWORD`, `DB_AUTO_MIGRATE`
- `JWT_SECRET`, `JWT_DURATION` (seconds)
- `REDIS_ADDRESS`, `REDIS_PASSWORD`

## Run
```bash
go run ./cmd
```
Server listens on `:${APP_PORT}`. Ent auto-migrate runs when enabled.

## Extending the system
1) Add or refine a port under `internal/interface/core/ports/...` if a new dependency is needed.
2) Implement the use case in `internal/usecase/<feature>` operating on domain models.
3) Provide an adapter in `internal/infrastructure/...` that satisfies the port, and wire it in `internal/app.Build`.
4) Expose via HTTP by adding a handler and route in `internal/interface/http/<feature>`, and register it in the router bootstrap.
