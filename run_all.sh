#!/bin/bash
#

for i in $(seq 0 29)
do
./implicit-genome-simulator-go --loci 1000 --iterations 1000 --envs 5 --datafile "output${i}.csv" > "progress${i}.out" &
done
