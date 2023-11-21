insert-all-temp:
	realpath temp/* | xargs -L1 go run cmd/insert_results.go -f 2>&1 | tee log.txt
