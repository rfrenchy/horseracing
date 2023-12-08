.PHONY: help
.PHONY: repository

PG_URL="postgresql://localhost/horse_racing?sslmode=disable"

help:
	@echo "TODO help"

build:
	go build ./...

scrape-results-gb:
	@./scripts/scrape_params.sh gb

tags:
	gotags

test:
	go test -v ./cmd/...

repository:
	./scripts/repository

migrate:
	migrate -database "postgresql://localhost/horse_racing?sslmode=disable" -path cmd/postgres/migrate up

