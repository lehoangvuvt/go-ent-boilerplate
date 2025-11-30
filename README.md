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
git clone https://github.com/yourusername/go-ent-boilerplate
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
```

---

## Running

```bash
go run ./cmd
```

API starts at:

```
http://localhost:8080
```

---

## Ent ORM

### Generate code:

```bash
go generate ./ent
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
go generate ./internal/app
```

---

## API Examples

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

---
