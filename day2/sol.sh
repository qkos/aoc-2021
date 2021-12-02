#!/bin/bash

IN=part1.in
echo "$(cat $IN| grep 'forward' | cut -d' ' -f2 | paste -sd+ | bc -l )*($(cat $IN| grep 'down' | cut -d' ' -f2 | paste -sd+ | bc -l )-$(cat $IN| grep 'up' | cut -d' ' -f2 | paste -sd+ | bc -l ))" | bc -l