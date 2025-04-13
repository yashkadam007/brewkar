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
make migrate-up
```

5. Run the application
```bash
make run
```

## Project Structure

- `cmd/api`: Application entry point
- `internal/domain`: Domain models
- `internal/repository`: Data access layer
- `internal/service`: Business logic layer
- `internal/controller`: API controllers
- `internal/middleware`: HTTP middleware
- `internal/config`: Configuration
- `pkg`: Reusable packages
- `migrations`: Database migration files

## API Documentation

See `docs/api.md` for detailed API documentation.

## Development

- Run tests: `make test`
- Create a new migration: `make migrate-create name=migration_name`
- Apply migrations: `make migrate-up`
- Rollback migrations: `make migrate-down`
- Run linter: `make lint`
