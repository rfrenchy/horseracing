.PHONY: help
.PHONY: gen-scrape-params

# README DIR SVG CHECKSUM
# Should update whenever there's a change to the repository's directory structure
# e.g rename file/dir, new file/dir, delete file/dir
# * * *  Does NOT Update ON Change in File Contents * * *
RDM_DIR_CKSUM='3086355355 14988'

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
	eza --long --tree --level 4 -I tools --git-ignore --no-user --no-time | svgbob --scale 1.5 --font-size 21 --font-family 'courier new' --background transparent > assets/dir.svg; \
	git add assets/dir.svg; \
	git commit --amend -C HEAD;

test:
	go test -v cmd/*_test.go
