#!/bin/bash

for val in {1..9}
do
    echo VM$val Setup
    ssh tian23@fa22-cs425-220$val.cs.illinois.edu "go env -w GO111MODULE='off'; go env -w GOPATH='/home/tian23/mp1-hangy6-tian23'; exit"
done
echo VM10 Setup
ssh tian23@fa22-cs425-2210.cs.illinois.edu "go env -w GO111MODULE='off'; go env -w GOPATH='/home/tian23/mp1-hangy6-tian23'; exit"

echo "Set Up All VM Envs!"
