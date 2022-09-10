#!/bin/bash

# bash set_env.sh
for val in {1..9}
do
    echo VM$val Building
    ssh hangy6@fa22-cs425-220$val.cs.illinois.edu "cd /home/hangy6/mp1-hangy6-tian23/src/scripts/; bash build.sh"
    echo VM$val Built
done
echo VM10 Building
ssh hangy6@fa22-cs425-2210.cs.illinois.edu "cd /home/hangy6/mp1-hangy6-tian23/src/scripts/; bash build.sh"
echo VM10 Built

echo "All VMs Built!!"