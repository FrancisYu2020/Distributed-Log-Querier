package main

import (
	"io"
	"log"
	"net/rpc"
	"os"
	"sync"
)

type ipAddress struct {
	address string
	name    string
}

// check wheter 'filename' file exists
func checkFileIsExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

func handleError(err error, c chan string, wg *sync.WaitGroup, ip ipAddress) {
	defer wg.Done()
	c <- string(ip.name + ".log: " + err.Error() + "\n")
}

func main() {
	var wg sync.WaitGroup // use wait group to keep synchronization
	defer wg.Wait()

	c := make(chan string) // use chanel to send logs safely
	ips := [10]ipAddress{
		{"fa22-cs425-2201.cs.illinois.edu:1234", "vm.1"},
		{"fa22-cs425-2202.cs.illinois.edu:1234", "vm.2"},
		{"fa22-cs425-2203.cs.illinois.edu:1234", "vm.3"},
		{"fa22-cs425-2204.cs.illinois.edu:1234", "vm.4"},
		{"fa22-cs425-2205.cs.illinois.edu:1234", "vm.5"},
		{"fa22-cs425-2206.cs.illinois.edu:1234", "vm.6"},
		{"fa22-cs425-2207.cs.illinois.edu:1234", "vm.7"},
		{"fa22-cs425-2208.cs.illinois.edu:1234", "vm.8"},
		{"fa22-cs425-2209.cs.illinois.edu:1234", "vm.9"},
		{"fa22-cs425-2210.cs.illinois.edu:1234", "vm.10"},
	}

	for _, ip := range ips {
		wg.Add(1) // add one when set a new task
		// use goutine to execute concurrently
		go func(ip ipAddress) { // connect to one server and try to execute RPC on that server
			client, err := rpc.Dial("tcp", ip.address) // set connection
			if err != nil {
				handleError(err, c, &wg, ip)
				return
			}

			var reply string
			err = client.Call("grepLogService.GrepLog", "grep -Ec log ../test_logs/log1 "+ip.name+".log: ", &reply) // RPC
			if err != nil {
				handleError(err, c, &wg, ip)
				return
			}
			c <- reply // use channel send logs back

			wg.Done() // minus one when finish a task
		}(ip)
	}

	var filename = "./test.txt" // path of the log file
	var f *os.File
	var err1 error
	if checkFileIsExist(filename) { // if file exists
		err1 = os.Remove(filename) // remove this file
	}
	f, err1 = os.Create(filename) // if file not exists, create the file

	if err1 != nil {
		log.Fatal(err1)
	}

	defer f.Close()
	for i := 0; i < len(ips); i++ {
		_, err1 := io.WriteString(f, <-c) // write logs to the file
		if err1 != nil {
			panic(err1)
		}
	}
}
