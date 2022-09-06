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

func main() {
	var wg sync.WaitGroup
	defer wg.Wait()

	c := make(chan string)
	ips := [3]ipAddress{{"172.22.156.72:1234", "machine 1"}, {"172.22.158.72:1234", "machine 2"}, {"172.22.94.72:1234", "machine 3"}}

	for _, ip := range ips {
		wg.Add(1)
		go func(ip ipAddress) {
			client, err := rpc.Dial("tcp", ip.address)
			if err != nil {
				log.Fatal(err)
			}

			var reply string
			err = client.Call("grepLogService.GrepLog", "grep -E log ../test_logs/log1", &reply)
			if err != nil {
				log.Fatal(err)
			}
			c <- reply

			wg.Done()
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
