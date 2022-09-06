#!/bin/bash

for val in {1..9}
do
    echo VM$val Updated
    ssh tian23@fa22-cs425-220$val.cs.illinois.edu "cd ./mp1-hangy6-tian23; git pull; exit"
done
echo VM10 Updated
ssh tian23@fa22-cs425-2210.cs.illinois.edu "cd ./mp1-hangy6-tian23; git pull; exit"

echo "Updated Done!"