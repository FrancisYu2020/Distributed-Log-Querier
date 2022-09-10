#!/bin/bash

for val in {1..9}
do
    echo VM$val downloading source code
    ssh hangy6@fa22-cs425-220$val.cs.illinois.edu "git clone https://gitlab.engr.illinois.edu/hangy6/mp1-hangy6-tian23.git; git checkout hangy6; git branch; exit"
    echo VM$val source code downloaded
done

echo "All VMs Have Been Updated!"