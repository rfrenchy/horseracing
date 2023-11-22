#!/bin/bash

# Run from Makefile, paths relative to project root
# Only for flat for now
# $1 = [course_id, coursename]

course_type=flat
#range=(2008 2009 2010 2011 2012 2013 2014 2015 2016 2017 2018 2019 2020 2021 2022 2023)
range=(2008)


courses=(cat './tools/rpscrape/courses/_courses' | jq '.gb' | jq -c 'to_entries[] | [.key, value] | .[]')

echo $courses

# echo "course id: $course_id"
# echo "course name: $course_name"

#echo -e $courses | head -n 2


# input an array of json kvp




# cat all the course ids? or a selection for now?
#for c in "${courses[@]}"; do:w:
#      echo $("$c" | jq -r '.key')
#done

# for year in "${range[@]}"; do
  #      echo "$course_type"
     #   (cd tools/rpscrape/scripts; ./rpscrape.py -y $year -c 17 -t $course_type)
        # ~/dev/horse_racing/tools/rpscrape/scripts/rpscrape.py; # add flags
#done





