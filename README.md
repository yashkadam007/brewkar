# Brewkar - Coffee Tracker API

A Go-based backend API for the Brewkar coffee tracking application.

## Prerequisites

- Go 1.18+
- PostgreSQL 13+
- Redis 6+

## Setup

1. Clone the repository
```bash
git clone https://github.com/yashkadam007/brewkar.git
cd brewkar
```

2. Install dependencies
```bash
go mod download
```

3. Create the database
```bash
createdb brewkar
```

4. Run database migrations
```bash
make migrate
```

5. Generate dependency injection code
```bash
make wire
```

6. Run the application
```bash
make run
```

## Project Structure

- `cmd/api`: Application entry point
- `internal/domain`: Domain models
- `internal/repository`: Data access layer
- `internal/service`: Business logic layer
- `internal/controller`: API controllers
- `internal/router`: Route definitions
- `internal/middleware`: HTTP middleware
- `internal/config`: Configuration
- `internal/di`: Dependency injection
- `pkg`: Reusable packages
- `migrations`: Database migration files

## API Documentation

See `docs/api.md` for detailed API documentation.

## Development

- Run tests: `make test`
- Apply migrations: `make migrate`
- Generate dependency injection code: `make wire`
- Run linter: `make lint`

## Dependency Injection

This project uses Google Wire for dependency injection. When making changes to the dependency graph:

1. Update the provider sets or provider functions in `internal/di/wire.go`
2. Run `make wire` to regenerate the implementation
3. The application will be automatically wired together at startup
