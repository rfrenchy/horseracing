temp-insert-all:
	realpath temp/* | xargs -L1 go run cmd/insert_results.go -f 2>&1 | tee log.txt

racecard-today: 
	cd ./tools/rpscrape/scripts; ./racecards.py today
