# Go Ent Clean/Hex Boilerplate

Minimal Go service using chi for HTTP, ent for data access, and a hexagonal/Clean split (domain ↔ use case ↔ ports, with infra adapters at the edge).

## Overview
- `cmd/`: entrypoints (`main.go` wires config, builds container, starts HTTP).
- `internal/app/`: composition root (builds infra, DI wiring, server lifecycle).
- `internal/domain/`: enterprise models + validation (`user`).
- `internal/usecase/`: application services (`auth` login, `user` create), DTOs and claim types.
- `internal/interface/core/ports/`: inbound/outbound interfaces (repository, security/JWT).
- `internal/interface/http/`: web adapters (router, handlers for auth/users).
- `internal/infrastructure/`: outbound adapters (ent client, user repository, JWT service wrapper).
- `pkg/`: shared helpers (e.g., `jwtx`, `httpx`, logging).
- `ent/`: ent schemas and generated code.

Dependency flow: handlers → use cases → ports → domain; infra implements ports and is injected from `internal/app`.

```mermaid
flowchart LR
    subgraph Entrypoint
        CMD[cmd/main.go]
    end
    subgraph Composition
        APP[internal/app\n(container & server)]
        BOOT[internal/bootstrap\n(stacks & router wiring)]
    end
    subgraph Interface
        HTTP[internal/interface/http\n(router & handlers)]
        PORTS[internal/interface/core/ports\n(repo, security/JWT)]
    end
    subgraph Usecase
        UC[internal/usecase\n(auth, user)]
    end
    subgraph Domain
        DOM[internal/domain\n(user entities)]
    end
    subgraph Infrastructure
        ENT[internal/infrastructure/ent\n(db client)]
        REPO[internal/infrastructure/repository\n(user repo)]
        JWT[internal/infrastructure/jwt\n(jwt service)]
    end
    subgraph Shared
        PKG[pkg/\nhttpx, jwtx, logger]
    end

    CMD --> APP --> BOOT --> HTTP --> UC --> DOM
    HTTP --> PORTS
    UC --> PORTS
    PORTS -.implements .-> ENT
    PORTS -.implements .-> REPO
    PORTS -.implements .-> JWT
    ENT -.uses .-> PKG
    REPO -.uses .-> ENT
    JWT -.wraps .-> PKG
    UC -.uses .-> PKG
```

## HTTP Surface
- `GET /health`
- `POST /api/v1/users/register` — create user (JSON: `email`, `password`)
- `POST /api/v1/auth/login` — validate credentials, returns `{ token, user }`

## Configuration
Environment variables (via Viper; `.` becomes `_`):
- `APP_PORT` (int)
- `DB_HOST`, `DB_PORT`, `DB_NAME`, `DB_USER`, `DB_PASSWORD`, `DB_AUTO_MIGRATE`
- `JWT_SECRET` (string), `JWT_DURATION_MINUTES` (int)

Sample `.env` is included; set a real `JWT_SECRET` in non-dev environments.

## Run
```bash
go run ./cmd
```
Server listens on `:${APP_PORT}` (default 0 if unset).

## Development Notes
- Ent: edit schemas under `ent/schema`, run `go generate ./ent` to regenerate code.
- JWT: `internal/infrastructure/jwt.Service` wraps `pkg/jwtx` and is injected via the JWT port.
- Tests are not yet added; consider per-layer tests (usecase with mocks, infra with DB, handler/route with httptest).
