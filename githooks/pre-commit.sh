#! /bin/bash

# Update SVG README Asset WHEN DIRECTORY STRUCTURE CHANGES
diff <(cat ./assets/dir.svg | cksum) <(cat Makefile | grep -e RDM_DIR_CKSUM= | awk -F '"' '{ print $2 }') &> /dev/null || make readme
