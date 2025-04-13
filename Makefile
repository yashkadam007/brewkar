.PHONY: run build test migrate-up migrate-down

run:
	go run cmd/api/main.go

build:
	go build -o bin/api cmd/api/main.go

test:
	@echo "Running tests..."
	@go test ./tests/... -v

migrate-create:
	migrate create -ext sql -dir migrations -seq $(name)

migrate-up:
	migrate -path migrations -database "postgres://postgres:postgres@localhost:5432/brewkar?sslmode=disable" up

migrate-down:
	migrate -path migrations -database "postgres://postgres:postgres@localhost:5432/brewkar?sslmode=disable" down

lint:
	golangci-lint run

# Add migration command
migrate:
	@echo "Running database migrations..."
	@go run migrations/001_create_tables.go

# Generate Wire code
wire:
	@echo "Generating dependency injection code..."
	@go run github.com/google/wire/cmd/wire ./internal/di
