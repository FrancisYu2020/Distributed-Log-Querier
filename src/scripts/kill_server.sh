#!/bin/bash

for val in {1..2}
do
    echo Killing VM$val Server
    ssh tian23@fa22-cs425-220$val.cs.illinois.edu "cd /home/tian23/mp1-hangy6-tian23/bin/; pkill server"
    echo VM$val Done
done
echo Killing VM10 Server
ssh tian23@fa22-cs425-2210.cs.illinois.edu "cd /home/tian23/mp1-hangy6-tian23/bin/; pkill server"
echo VM10 Done

echo "All Servers Shut Down!"