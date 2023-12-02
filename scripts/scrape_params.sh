#! /bin/bash

# grab given region course ids from racingpost
TMP_COURSE=$(mktemp)
cat tools/rpscrape/courses/_courses | jq -r ".$1 | keys[]" > $TMP_COURSE

# create year range
TMP_YEARS=$(mktemp)
seq 2008 2023 > $TMP_YEARS

# create cartesian product of both then shuffle to look less bait
join -j 2 $TMP_COURSE $TMP_YEARS | shuf

# clean temp files
trap "rm $TMP_COURSE; rm $TMP_YEARS" EXIT

# TODO make so each line an rpscrape command i.e. rpscrape -c $1 -y $2 -t flat
