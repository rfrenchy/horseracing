#! /bin/bash

# grab given region course ids from racingpost
TMP_COURSE=$(mktemp)
cat tools/rpscrape/courses/_courses | jq -r ".$1 | keys[]" > $TMP_COURSE

# create year range
TMP_YEARS=$(mktemp)
seq 2008 2023 > $TMP_YEARS

# create cartesian product of both then shuffle to look less bait
mkfifo params
join -j 2 $TMP_COURSE $TMP_YEARS | shuf > params &

# clean temp files
clean() {
        rm $TMP_COURSE
        rm $TMP_YEARS
        unlink params
}


cat params

trap clean EXIT

# TODO make so each line an rpscrape command i.e. rpscrape -c $1 -y $2 -t flat
# TODO intersperse with vpn connect/disconnect every x(10?) lines
# use AWK?
