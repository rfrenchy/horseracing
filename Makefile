.PHONY: help

help: 
	@echo "TODO help"

courses: 
	cat ./tools/rpscrape/courses/_courses

temp-insert-aintree:
	realpath tools/rpscrape/data/courses//Aintree/flat/* | xargs -L1 go run cmd/insert_results.go -f 2>&1 | tee log.txt

temp-insert-ascot:
	realpath tools/rpscrape/data/courses//Ascot/flat/* | xargs -L1 go run cmd/insert_results.go -f 2>&1 | tee log.txt

newmarket-insert-temp:
	realpath tools/rpscrape/data/courses/Newmarket/flat/* | \
	xargs -L1 go run cmd/insert_results.go -f 2>&1 | \
	tee log.txt

racecard-today: 
	cd ./tools/rpscrape/scripts; ./racecards.py today

newmarket:
	cd ./tools/rpscrape/scripts; ./rpscrape.py -c 38 -y 2020-2023 -t flat

scrape:
	go run scripts/racing_post.go
