.PHONY: help
.PHONY: gen-scrape-params

PG_URL="postgresql://localhost/horse_racing?sslmode=disable"

help:
	@echo "TODO help"

build:
	go build ./...

test:
	go test ./... -v

scrape-results-gb:
	@./scripts/scrape_params.sh gb

tags:
	gotags

test:
	go test -v cmd/*_test.go

migrate:
	migrate -database "postgresql://localhost/horse_racing?sslmode=disable" -path cmd/postgres/migrate up

