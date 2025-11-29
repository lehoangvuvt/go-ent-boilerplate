# Go Ent Clean/Hex Boilerplate

Minimal Go HTTP service built with chi and ent, organized around clean/hexagonal boundaries (domain + use cases at the center, interfaces/ports at the edges, adapters from infrastructure).

## Architecture at a Glance

- Entrypoint: `cmd/main.go` loads config, builds the app container, and starts the HTTP server with graceful shutdown.
- Composition root: `internal/app` + `internal/bootstrap` wire infrastructure (ent/Postgres, Redis cache, JWT, idempotency store) and handler stacks.
- Domain core: `internal/domain` holds entities, value objects, and behaviors; no IO.
- Use cases: `internal/usecase` orchestrate workflows against domain and ports.
- Ports: `internal/interface/core/ports` define the IO contracts that use cases consume.
- Adapters: `internal/infrastructure` implements the ports (ent repositories, JWT, Redis cache/idempotency, RabbitMQ queue).
- Transport: `internal/interface/http` handles routing, auth middleware, idempotency middleware, request/response shaping.

## Directory Map

- `cmd/` � service entrypoint.
- `internal/app/` � container + server lifecycle.
- `internal/bootstrap/` � DSN builders, handler stacks, router bootstrap.
- `internal/config/` � config loader + validation.
- `internal/domain/` � business entities and behaviors (user, transaction + payment methods).
- `internal/usecase/` � auth, user creation, transaction creation/list/find/confirm/cancel/fail.
- `internal/interface/core/ports/` � repository, cache, queue, JWT, idempotency contracts.
- `internal/interface/http/` � chi router, handlers, middleware (auth, idempotency).
- `internal/infrastructure/` � ent client, repos, Redis cache, Redis idempotency store, JWT service, RabbitMQ adapter.
- `ent/` � schemas and generated ent code.
- `pkg/` � helpers (`logger`, `jwtx`, `httpx`).

## Features

- JWT auth with login flow (`/api/v1/auth/login`), bearer validation middleware, and claim injection.
- User registration (`/api/v1/users/register`) with password hashing and entity validation.
- Transaction lifecycle: create, list, fetch by ID, confirm, cancel, fail; enforces state transitions in domain.
- Idempotency middleware (Redis-backed) for authenticated routes: honors `Idempotency-Key` header, caches successful responses (TTL 10m), rejects duplicate in-flight requests.
- Redis cache adapter implementing a simple key/value TTL cache port.
- Ent/Postgres persistence with optional auto-migrate on startup.
- RabbitMQ producer/consumer adapter + stack builder available for queue-enabled features.

## Data Model (core)

- User: `id`, `email`, `hashed_password`, timestamps.
- Transaction: `id`, `user_id`, `amount`, `currency`, `status` (`pending`, `completed`, `failed`, `rejected`), `method` (`visa`, `banking`, `ewallet`, `qr`), optional method-specific detail blobs (visa/banking/ewallet/qr), timestamps.

## Request Flow

1. `cmd/main.go` ? `config.Load()` reads env (dotenv + env vars, `.` ? `_`), validates required fields.
2. `internal/app.Build()` boots ent/Postgres, Redis cache, JWT service, Redis idempotency store, repositories, and handler stacks.
3. `internal/interface/http/router` registers routes; auth middleware protects `/api/v1/transactions`, then idempotency middleware runs, then handlers call use cases.
4. Use cases enforce validation/state and call ports; adapters hit DB/Redis/JWT.

## API Surface (HTTP)

- `GET /health`
- `POST /api/v1/users/register` `{ "email": "...", "password": "..." }`
- `POST /api/v1/auth/login` `{ "email": "...", "password": "..." }` ? `{ token, user }`
- Authenticated (Bearer token):
  - `POST /api/v1/transactions` with `Idempotency-Key: <uuid>` and body:
    ```json
    {
      "amount": 1000,
      "currency": "USD",
      "method": "visa",
      "visa": {
        "card_last_4": "4242",
        "card_network": "visa",
        "authorization_code": "abc",
        "reference_id": "ref-1",
        "is_3d_secure": true
      }
    }
    ```
  - `GET /api/v1/transactions?user_id=<uuid>&status=pending&method=visa&limit=20&offset=0`
  - `GET /api/v1/transactions/{id}`
  - `POST /api/v1/transactions/{id}/confirm`
  - `POST /api/v1/transactions/{id}/cancel`
  - `POST /api/v1/transactions/{id}/fail`

## Configuration

Env keys (dotenv loaded first):

- `APP_PORT`
- `DB_HOST`, `DB_PORT`, `DB_NAME`, `DB_USER`, `DB_PASSWORD`, `DB_AUTO_MIGRATE`
- `JWT_SECRET`, `JWT_DURATION` (seconds)
- `REDIS_ADDRESS`, `REDIS_PASSWORD`
- RabbitMQ DSN (if you use the queue stack): e.g. `amqp://guest:guest@localhost:5672/`
  Sample values live in `.env`.

## Run & Develop

- Prereqs: Go 1.24+, Postgres, Redis; RabbitMQ optional for queue experiments.
- Install deps: `go mod download`
- Start (auto-migrate if enabled): `go run ./cmd`
- Make targets: `make fmt`, `make lint`, `make run`, `make build`, `make ent-create name=<Schema>`, `make ent-gen` (regenerate ent after schema edits).
- Tests: `go test ./...` (add focused tests around use cases/ports as you extend).

## Idempotency Notes

- Add `Idempotency-Key` to POST/PUT calls under `/api/v1/transactions` to dedupe.
- On first request: store `pending`, then persist `success`/`failed` with response; later identical keys return cached success or `409 Conflict` if still pending.
- TTL configured in `internal/app/container.go` (default 10m via Redis).

## Extending

- New dependencies: add a port under `internal/interface/core/ports`, implement in `internal/infrastructure`, wire in `internal/app.Build()`.
- New features: add use cases under `internal/usecase/<feature>`, domain rules in `internal/domain`, HTTP handlers/routes under `internal/interface/http/<feature>`, register via a stack in `internal/bootstrap/stack`.
- Schema changes: edit `ent/schema`, run `make ent-gen`, then restart.

## Troubleshooting

- DB issues: confirm DSN in env and Postgres running; disable auto-migrate if you manage schema manually.
- Redis required for idempotency/cache; startup pings Redis in `internal/app/container.go`.
- JWT errors: ensure `JWT_SECRET` and `JWT_DURATION` are set; auth middleware expects `Authorization: Bearer <token>`.
