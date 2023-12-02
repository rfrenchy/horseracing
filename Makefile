.PHONY: help
.PHONY: gen-scrape-params

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
