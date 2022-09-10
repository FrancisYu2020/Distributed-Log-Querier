#!/bin/bash

for val in {1..9}
do
    echo VM$val source code deleting
    ssh hangy6@fa22-cs425-220$val.cs.illinois.edu "rm -rf ./mp1-hangy6-tian23; exit"
    echo VM$val source code deleted
done

echo "All VMs Have Been Updated!"