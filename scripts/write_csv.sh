#! /bin/bash

set -e

if [[ $# -ne 2 ]] ; then
        echo 'course id and year required'
        exit 0
fi

go run ./cmd/repository/... csv -f /home/ryan/dev/punts/tools/rpscrape/data/all/flat/$2.csv -cid $1 -y $2
