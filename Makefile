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

readme:
	eza --long --tree --level 4 -I tools --no-permissions --no-filesize --no-user --no-time --git-ignore | svgbob --scale 2.5 --font-size 21 --font-family 'courier new' --background transparent > assets/structure.svg

test:
	go test -v cmd/*_test.go
