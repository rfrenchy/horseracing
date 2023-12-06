#!/bin/bash

OUT="$(mktemp)"

echo "username:"
read USR
echo $USR >> $OUT

echo "password:"
read -s PASS
echo $PASS >> $OUT

function con() {
        sudo openvpn --data-ciphers 'AES-256-CBC' --config /etc/openvpn/ovpn_tcp/$(ls /etc/openvpn/ovpn_tcp | grep -i -e uk | shuf -n 1) --auth-user-pass $1
}

for x in $@
do
        con $OUT
        echo $x
        (cd ./tools/rpscrape/scripts; ./rpscrape.py -c $x -y 2008-2023 -t flat)

        sudo killall openvpn
done

trap '{ sudo killall openvpn; rm $OUT; }' EXIT
