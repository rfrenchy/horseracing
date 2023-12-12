#! /bin/bash

main() {
        # exit if VPN errors
        set -e

        if [[ $# -eq 0 ]] ; then
                echo 'need a course id in order to mine'
                exit 0
        fi

        # connect to vpn
        vpn

        # mine racingpost
        gen $1 | shuf | mine
}

# generate mining commands
gen() {
        seq 2008 2022 |
                while read Y; do echo "./rpscrape.py -c $1 -y $Y -t flat"; done
}

# execute mining commands
mine() {
        cd ~/dev/punts/tools/rpscrape/scripts

        while read -r com; do eval $com; done

        cd --
}

main $1

