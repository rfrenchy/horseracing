#! /bin/bash

# get all gb course ids from racingpost
cat tools/rpscrape/courses/_courses | jq -r '.gb | keys[]' > scripts/gb_course_ids

# get year range we want
seq 2008 2023 > scripts/years

# get cartesian product of both, shuffle
join -j 2 scripts/gb_course_ids scripts/years | shuf > scripts/rpscrape_params
