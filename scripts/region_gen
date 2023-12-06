#! /bin/bash


# partial insert script value string

cat tools/rpscrape/courses/_regions \
        | jq '.' \
        | sed 's/:/,/gm' \
        | sed '1,2d;$d' \
        | awk '{ print "(" NR "," $0  }'


