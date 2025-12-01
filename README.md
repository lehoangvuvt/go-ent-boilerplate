# go-ent-boilerplate

[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8.svg)](https://golang.org/doc/go1.22)
[![Build Status](https://img.shields.io/badge/build-passing-brightgreen)]()
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-green.svg)]()

A clean, modular, and production-ready Go backend starter kit built with **Clean Architecture**, **Ent ORM**, **Google Wire DI**, **JWT Authentication**, and **Redis**.

This project is designed for developers and teams who want a **high-quality, scalable, testable, and open-source friendly backend foundation**.

---

## Features

- **Clean Architecture** separation
- **Ent ORM** (schema-as-code, type-safe queries)
- **Google Wire** dependency injection
- **Chi router** HTTP stack
- **Redis** cache & idempotency
- **Structured module layout**
- **Production-ready Makefile**
- **Environment-based configuration**

---

## Architecture Overview

```
cmd/                 # Application entrypoint (main.go)
internal/
  app/               # DI, container, wire injector, server bootstrap
  bootstrap/         # Router + middleware bootstrap
  config/            # Viper configuration loader
  domain/            # Domain entities and logic
  infrastructure/     # Database, Redis, JWT, external services
  interface/
    core/ports       # Interfaces for repositories & services
    http/            # Handlers, middleware, router
  usecase/           # Application business logic
ent/                 # Ent schema + generated code
pkg/                 # Shared utilities
```

---

## Architecture Diagram

```
                      ┌────────────────────────────────────────────┐
                      │                  cmd/                      │
                      │        (main.go / server startup)          │
                      └────────────────────────────────────────────┘
                                        │
                                        ▼
                      ┌────────────────────────────────────────────┐
                      │              internal/app/                 │
                      │     Wire DI → Build Container → Router    │
                      │  wire.go / wire_sets.go / wire_gen.go     │
                      └────────────────────────────────────────────┘
                                        │
                                        ▼
         ┌──────────────────────────────────────────────────────────────────────────┐
         │                             Bootstrap Layer                              │
         │                      internal/bootstrap / router                         │
         └──────────────────────────────────────────────────────────────────────────┘
                                        │
                                        ▼
        ┌────────────────────────────────────────────────────────────────────────────┐
        │                              HTTP Interface                                │
        │                internal/interface/http/{handlers,middleware}               │
        └────────────────────────────────────────────────────────────────────────────┘
                                        │
                                        ▼
        ┌────────────────────────────────────────────────────────────────────────────┐
        │                                Usecases                                    │
        └────────────────────────────────────────────────────────────────────────────┘
                                        │
                                        ▼
        ┌────────────────────────────────────────────────────────────────────────────┐
        │                               Repositories                                 │
        └────────────────────────────────────────────────────────────────────────────┘
                                        │
                                        ▼
        ┌────────────────────────────────────────────────────────────────────────────┐
        │                              Infrastructure                                │
        └────────────────────────────────────────────────────────────────────────────┘
                                        │
                                        ▼
                         ┌───────────────────────────────────────┐
                         │         External Systems               │
                         │      Postgres | Redis | Services       │
                         └───────────────────────────────────────┘
```

---

## Installation

```bash
git clone https://github.com/lehoangvuvt/go-ent-boilerplate
cd go-ent-boilerplate
go mod tidy
```

---

## Configuration

Create a `.env` file:

```
APP_PORT=8080

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=example
DB_AUTO_MIGRATE=true

REDIS_ADDRESS=localhost:6379
REDIS_PASSWORD=

JWT_SECRET=supersecret
JWT_DURATION=3600

SMTP_HOST=
SMTP_PORT=
SMTP_USER=
SMTP_PASS=
```

---

## Running

```bash
go run ./cmd
```

API starts at:

```
http://localhost:{{ your-configured-port }}
```

---

## Deployment (Docker Compose)

```bash
# fill secrets in .env.prod first: JWT_SECRET, SMTP_USER, SMTP_PASS, RESEND_API_KEY
docker compose --env-file .env.prod up -d --build
docker compose logs -f server worker
```

Notes:
- Containers: `server` (API) and `worker` (background jobs) share the same image.
- External deps run locally via Compose: Postgres on 5432, Redis on 6379, RabbitMQ on 5672/15672, RedisInsight on 5540.
- `APP_PORT` in `.env.prod` maps the API to `http://localhost:${APP_PORT}`.
- Set real production secrets before bringing the stack up; `.env.prod` supports overriding via env vars when the Compose file is parsed.

---

## Ent ORM

### Create new schema:

```bash
make ent-create name=Example123
```

### Generate code:

```bash
make ent-gen
```

### Auto migration:

Enabled via:

```
DB_AUTO_MIGRATE=true
```

---

## Dependency Injection (Google Wire)

Provider sets are in:

```
internal/app/wire_sets.go
```

Generate Wire code:

```bash
make wire-gen
```

---

## API Examples

### Register

**POST** `/api/v1/users/register`

```json
{
  "email": "user@example.com",
  "password": "123456"
}
```

### Login

**POST** `/api/v1/auth/login`

```json
{
  "email": "user@example.com",
  "password": "123456"
}
```

### Create Transaction

**POST** `/api/v1/transactions`

Headers:

```
Authorization: Bearer <token>
Idempotency-Key: <uuid>
```

Payload:

```json
{
  "ID": "14a61d4f-147e-4930-bdb2-f6672a2022cf",
  "Amount": 200000,
  "Currency": "VND",
  "UserID": "7938c4fa-701f-48b6-8d47-f048e97d0901",
  "Method": "banking",
  "Visa": null,
  "Banking": {
    "bank_code": "VCB",
    "bank_name": "Vietcombank",
    "account_number": "1234567890",
    "reference_id": "BANK987654321",
    "transaction_time": "2025-01-20T10:05:00Z"
  },
  "EWallet": null,
  "QRPay": null,
  "Status": "pending",
  "CreatedAt": "2025-11-30T12:16:26.0223125+07:00",
  "UpdatedAt": "2025-11-30T12:16:26.0223125+07:00"
}
```

---
