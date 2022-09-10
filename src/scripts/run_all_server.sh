#!/bin/bash

for val in {1..9}
do
    echo VM$val Starting
    ssh hangy6@fa22-cs425-220$val.cs.illinois.edu "cd /home/hangy6/mp1-hangy6-tian23/bin/; nohup ./server>/dev/null 2>&1&"
    echo VM$val Started
done
echo VM10 Starting
ssh hangy6@fa22-cs425-2210.cs.illinois.edu "cd /home/hangy6/mp1-hangy6-tian23/bin/; nohup ./server>/dev/null 2>&1&"
echo VM10 Started

echo "All Servers Have Been Started!"