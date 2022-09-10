#!/bin/bash

for val in {1..3}
do
    echo VM$val Updating
    ssh hangy6@fa22-cs425-220$val.cs.illinois.edu "git clone git@gitlab.engr.illinois.edu:hangy6/mp1-hangy6-tian23.git; git checkout hangy6; git branch; git config --global user.name 'hangy6'; git config --global user.email hangy6@illinois.edu; exit"
    echo VM$val Updated
done

echo "All VMs Have Been Updated!"