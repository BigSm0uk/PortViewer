# Go Simple Service Template

A small, opinionated starting point for Go HTTP services. It keeps application
wiring explicit, uses the standard library for HTTP, and includes a compact set
of development tasks without introducing a framework.

## What's included

- Application lifecycle wiring under `internal/app`
- YAML configuration loaded from an explicit `--config` path
- Structured JSON or console logging with [Zap](https://github.com/uber-go/zap)
- A standard-library `net/http` server
- A `GET /ping` health endpoint
- [Task](https://taskfile.dev/) commands for tidy, test, vet, build, run, and release

## Requirements

- Go 1.27 or newer
- Task 3

## Use this template

1. Select **Use this template** on GitHub and create a new repository.
2. Clone your new repository and enter its directory.
3. Change the module path:

   ```bash
   go mod edit -module github.com/<owner>/<repository>
   ```

4. Replace `github.com/bigsm0uk/go-simple-service-template` in Go imports with
   your new module path.
5. Rename `cmd/go-simple-service-template` and update `APP_DIR` and `MAIN_DIR`
   in `taskfiles/app.yml` if you want the executable to use your service name.
6. Refresh the module metadata:

   ```bash
   go mod tidy
   ```

## Quick start

Build and run the service with the included configuration:

```bash
task app:run
```

In another terminal, check the health endpoint:

```bash
curl http://127.0.0.1:8080/ping
```

The response is `pong`.

You can also run the service without Task:

```bash
go run ./cmd/go-simple-service-template --config ./config.yaml
```

## Tasks

Run `task --list` to see the available commands.

| Command | Description |
| --- | --- |
| `task app:tidy` | Synchronize `go.mod` and `go.sum` |
| `task app:test` | Run all Go tests |
| `task app:vet` | Run Go static analysis |
| `task app:build` | Build `bin/go-simple-service-template` |
| `task app:run` | Tidy, build, and run with `config.yaml` |
| `task app:release` | Run tidy, tests, vet, and build |

The short alias also works, for example `task a:test`.

## Configuration

The service requires a configuration path through the `--config` flag. The
included `config.yaml` provides a local development setup:

```yaml
server:
  host: 127.0.0.1
  port: 8080
  open_browser: true
logging:
  level: debug
  format: json
```

Logging accepts `debug`, `info`, `warn`, `error`, and the other Zap levels.
Use `json` for structured logs or `text`/`console` for console output.

## Project structure

```text
.
├── cmd/
│   └── go-simple-service-template/
│       └── main.go
├── internal/
│   └── app/
│       ├── config/
│       │   ├── config.go
│       │   └── config_test.go
│       ├── logger/
│       │   └── zap.go
│       ├── application.go
│       └── routes.go
├── taskfiles/
│   └── app.yml
├── config.yaml
├── go.mod
└── Taskfile.yml
```

## Customize

- [ ] Change the Go module path and update internal imports.
- [ ] Rename the command directory and binary paths for your service.
- [ ] Replace the example configuration with your service settings.
- [ ] Add routes and dependencies through `internal/app`.
- [ ] Replace `/ping` with the health contract your deployment expects.
- [ ] Add project-specific tests and CI checks.
- [ ] Choose and add a license before public distribution.
