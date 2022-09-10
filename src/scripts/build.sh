#!/bin/bash

cd ~/mp1-hangy6-tian23
# export GOPATH=`pwd`
cd ~/mp1-hangy6-tian23/src/client
go build -o ../../bin/client ./client.go
cd ~/mp1-hangy6-tian23/src/server
go build -o ../../bin/server ./server.go