# MP1-hangy6-tian23 Distributed Log Querier

## Description
A Go distributed log querier with concurrency and fault tolerance that allows you to query distributed log files on multiple machines from any one of those machines. This project is implemented by hangy6(Hang Yu) and tian23(Tian Luan).


## Installation

You can clone this project to the machines you need to grep log from using following command:

```
ssh: git clone git@gitlab.engr.illinois.edu:hangy6/mp1-hangy6-tian23.git
```
```
https: git clone https://gitlab.engr.illinois.edu/hangy6/mp1-hangy6-tian23.git
``` 

## Build

To build the client and server programs easily, please follow the commands below:

```
cd src/scripts/
bash build.sh
```

or you can build the client and server program by yourself following the commands below:

```
cd src/
go build -o ../bin/client ./client_main.go
cd src/
go build -o ../bin/server ./server_main.go
```

## Usage

To grep log from multiple machines, please build the server and client program first following previous instruction, and then run server program on these machines. After that please run client program on the machine you want to use.

To run server program on a machine, please use following command:

```
bin/server
```

You can use config.json to configure the machines you want to query log from.

To run client program on a machine, please use following command:

```
bin/client
```

Then you need input the query command like following:

```
grep [options] [pattern] [log_name] [output_file_path](optional)
```

Here for [options] we have:

-c  counts the number of lines that contain matching pattern in a file and prints it or (output to specified file)

-Ec counts the number of lines that contain matching pattern using regex in a file and prints it or (output to specified file)

[output_file_path] is optional, if you input this, the result of grep command will be output to [output_file_path]


If you want to run server on multiple machines at the same time, you can use the run_all_server.sh at src/scripts/ but please make slight change for your ipaddress.


## Tests
To run unit test, please use following command to run the corresponding script:

```
go test src/test/[test_script_name]
```

If you want to see the execute time of tests, please use option -v.

## Support
If you have any questions, please contact tian23@illinois.edu or hangy6@illinois.edu

## Authors 
Tian Luan & Hang Yu

