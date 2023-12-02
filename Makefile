.PHONY: help
.PHONY: gen-scrape-params

help:
	@echo "TODO help"

courses:
	bat ./tools/rpscrape/courses/_courses.json

migrate:
	migrate -source ./db/migrations/* -database postgres://localhost:5432/horse_racing up 2

racecard-today:
	cd ./tools/rpscrape/scripts; ./racecards.py today

racing-post:
	go run cmd/racing_post.go

gen-scrape-params-gb:
	@./scripts/scrape_params.sh gb

test:
	go test -v cmd/*_test.go
