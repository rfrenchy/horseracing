.PHONY: help
.PHONY: gen-scrape-params

help:
	@echo "TODO help"

courses:
	bat ./tools/rpscrape/courses/_courses.json

insert-aintree:
	realpath tools/rpscrape/data/courses/Aintree/flat/* | \
	xargs -L1 go run cmd/insert_results.go -c -f 2>&1 | \
	tee log.txt

insert-ascot:
	realpath tools/rpscrape/data/courses/Ascot/flat/* | \
	xargs -L1 go run cmd/insert_results.go -c -f 2>&1 | \
	tee log.txt

insert-bangor:
	realpath tools/rpscrape/data/courses/Bangor-on-Dee/flat/* | \
	xargs -L1 go run cmd/insert_results.go -c -f 2>&1 | \
	tee log.txt

insert-doncaster:
	realpath tools/rpscrape/data/courses/Doncaster/flat/* | \
	xargs -L1 go run cmd/insert_results.go -c -f 2>&1 | \
	tee log.txt

insert-epsom:
	realpath tools/rpscrape/data/courses/Epsom/flat/* | \
	xargs -L1 go run cmd/insert_results.go -c -f 2>&1 | \
	tee log.txt

insert-goodwood:
	realpath tools/rpscrape/data/courses/Goodwood/flat/* | \
	xargs -L1 go run cmd/insert_results.go -c -f 2>&1 | \
	tee log.txt

insert-lingfield:
	realpath tools/rpscrape/data/courses/Lingfield/flat/* | \
	xargs -L1 go run cmd/insert_results.go -c -f 2>&1 | \
	tee log.txt

insert-newmarket:
	realpath tools/rpscrape/data/courses/Newmarket/flat/* | \
	xargs -L1 go run cmd/insert_results.go -c -f 2>&1 | \
	tee log.txt

insert-sandown:
	realpath tools/rpscrape/data/courses/Sandown/flat/* | \
	xargs -L1 go run cmd/insert_results.go -c -f 2>&1 | \
	tee log.txt

insert-york:
	realpath tools/rpscrape/data/courses/York/flat/* | \
	xargs -L1 go run cmd/insert_results.go -c -f 2>&1 | \
	tee log.txt

migrate:
	migrate -source ./db/migrations/* -database postgres://localhost:5432/horse_racing up 2

racecard-today:
	cd ./tools/rpscrape/scripts; ./racecards.py today

racing-post:
	go run cmd/racing_post.go

gen-scrape-params: clean-scrape-gen
	$(shell ./scripts/gen_rpscrape_params.sh)

clean-scrape-gen:
	@rm -f scripts/rpscrape_params
