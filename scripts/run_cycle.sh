#! /bin/bash

set -e

if [[ $# -ne 2 ]] ; then
        echo 'course id and year required'
        exit 0
fi

go run ./cmd/repository/... model -cid $1 -y $2
