.PHONY: help
.PHONY: repository

PG_URL="postgresql://localhost/horse_racing?sslmode=disable"

help:
	@echo "TODO help"

build:
	go build ./...

scrape:
	@./scripts/scrape_params.sh gb

tags:
	gotags

test:
	go test -v ./cmd/...

repository:
	./scripts/repository

start-db:
	docker run -d --name postgres-db -p 5433:5432 -e POSTGRES_PASSWORD=password postgres:latest

migrate:
	migrate -database "postgresql://localhost/horse_racing?sslmode=disable" -path cmd/postgres/migrate up

