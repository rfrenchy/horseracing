.PHONY: help

help: 
	@echo "TODO help"

courses: 
	cat ./tools/rpscrape/courses/_courses

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

racecard-today: 
	cd ./tools/rpscrape/scripts; ./racecards.py today

newmarket:
	cd ./tools/rpscrape/scripts; ./rpscrape.py -c 38 -y 2020-2023 -t flat

scrape:
	go run scripts/racing_post.go
