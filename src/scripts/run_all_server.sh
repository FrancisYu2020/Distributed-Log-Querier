#!/bin/bash

for val in {1..9}
do
    echo VM$val Starting
    ssh tian23@fa22-cs425-220$val.cs.illinois.edu "cd /home/tian23/mp1-hangy6-tian23/src/server/; nohup go run server.go>/dev/null 2>&1&"
done
echo VM10 Starting
ssh tian23@fa22-cs425-2210.cs.illinois.edu "cd /home/tian23/mp1-hangy6-tian23/src/server/; nohup go run server.go>/dev/null 2>&1&"

echo "All servers has been started!"