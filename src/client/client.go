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

func checkFileIsExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

func handleError(err error, c chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	c <- string(err.Error())
}

func main() {
	var wg sync.WaitGroup // use wait group to keep synchronization
	defer wg.Wait()

	c := make(chan string) // use chanel to send logs safely
	ips := [3]ipAddress{{"172.22.156.72:1234", "machine 1"}, {"172.22.158.72:1234", "machine 2"}, {"172.22.94.72:1234", "machine 3"}}

	for _, ip := range ips {
		wg.Add(1)               // add one when set a new task
		go func(ip ipAddress) { // connect to one server and try to execute RPC on that server
			client, err := rpc.Dial("tcp", ip.address) // set connection
			if err != nil {
				handleError(err, c, &wg)
				return
			}

			var reply string
			err = client.Call("grepLogService.GrepLog", "grep -E log ../test_logs/log1", &reply) // RPC
			if err != nil {
				handleError(err, c, &wg)
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
	for i := 0; i < 3; i++ {
		_, err1 := io.WriteString(f, <-c) // write logs to the file
		if err1 != nil {
			panic(err1)
		}
	}
}
