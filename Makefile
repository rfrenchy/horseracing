.PHONY: help
.PHONY: gen-scrape-params

PG_URL="postgresql://localhost/horse_racing?sslmode=disable"

help:
	@echo "TODO help"

courses:
	bat ./tools/rpscrape/courses/_courses.json

migrate:
	migrate -source ./db/migrations/* -database postgres://localhost:5432/horse_racing up 2

scrape-results-gb:
	@./scripts/scrape_params.sh gb

tags:
	gotags

test:
	go test -v cmd/*_test.go

migrate-up:
	migrate -database ${PG_URL} -path cmd/postgres/migrate up

migrate-down:
	migrate -database ${PG_URL} -path cmd/postgres/migrate down
