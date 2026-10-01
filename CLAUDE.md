# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project

`devboard` is a Go HTTP API backend (learning project, early scaffolding stage). Code comments and Makefile help text are written in Spanish; follow that convention.

## Commands

All common tasks go through the Makefile (`make help` lists them):

- `make run` — run the API (`go run ./cmd/api/main.go`)
- `make build` — build binary to `bin/devboard`
- `make test` — `go test -race -cover ./...`
- Single test: `go test -race -run TestName ./internal/server/`
- `make lint` — `golangci-lint run ./...` (v2 config in `.golangci.yml`: errcheck with type-assertion checks, staticcheck, govet, gosec (skipped in `_test.go`), revive, misspell; formatters gofmt + goimports)
- `make tidy` — `go mod tidy && go mod verify`
- `make generate` — `go generate ./...`
- `make docker-up` / `docker-down` / `docker-logs` — local Postgres 16 (`localhost:5432`, user `postgres` / pw `password`, db `devboard`) and Redis 7 (`localhost:6379`) via `docker-compose.yml`
- `make migrate-up` / `migrate-down` — golang-migrate CLI against `./migrations` using the `DB_URL` in the Makefile (down reverts only one step)

## Architecture

Standard Go layout:

- `cmd/api/main.go` — entrypoint; should only wire config and start the server.
- `internal/server` — `Server` wraps an `*http.Server` plus a stdlib `*http.ServeMux` (no router framework). `server.New(addr)` sets read/write/idle/read-header timeouts; `Start()` calls `ListenAndServe`. Routes are meant to be registered on the internal `mux`.
- `configs/config.yaml` and `.env.example` — config placeholders (currently empty).

## Current state / gotchas

- Module path: `github.com/sergioIbz11/devboard` (public repo on GitHub).
- `cmd/api/main.go` does not compile: it imports a filesystem path instead of `github.com/sergioIbz11/devboard/internal/server` and references an undefined `httpServer`. It should call `server.New(addr)` and `Start()`.
- `./migrations` does not exist yet.
