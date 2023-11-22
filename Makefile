.PHONY: help

GBK := $(shell cat './tools/rpscrape/courses/_courses' | jq '.gb' | jq -c 'to_entries[] | [.key] | .[]')

GBV := $(shell cat './tools/rpscrape/courses/_courses' | jq '.gb' | jq -c 'to_entries[] | [.value] | .[]')

help: 
	@echo "TODO help"

courses: 
	cat ./tools/rpscrape/courses/_courses

temp-insert-all:
	realpath temp/* | xargs -L1 go run cmd/insert_results.go -f 2>&1 | tee log.txt

newmarket-insert-temp:
	realpath tools/rpscrape/data/courses/Newmarket/flat/* | \
	xargs -L1 go run cmd/insert_results.go -f 2>&1 | \
	tee log.txt

racecard-today: 
	cd ./tools/rpscrape/scripts; ./racecards.py today

newmarket:
	cd ./tools/rpscrape/scripts; ./rpscrape.py -c 38 -y 2020-2023 -t flat

vpn: 
	sudo openvpn --data-ciphers 'AES-256-CBC' --auth-nocache --config /etc/openvpn/ovpn_tcp/uk2287.nordvpn.com.tcp.ovpn

scrape-gb-flat:
	./scripts/scrape.sh
