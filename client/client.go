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
	ips := [2]ipAddress{{"127.0.0.1:1234", "machine 1"}, {"127.0.0.1:1235", "machine 2"}}

	wg.Add(2)
	for _, ip := range ips {
		go func(ip ipAddress) {
			client, err := rpc.Dial("tcp", ip.address)
			if err != nil {
				log.Fatal(err)
			}

			var reply string
			err = client.Call("grepLogService.GrepLog", "grep -c log ../test_logs/log1", &reply)
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
		f, err1 = os.OpenFile(filename, os.O_APPEND, 0666) //open the file
	} else {
		f, err1 = os.Create(filename) // if file not exists, create the file
	}
	if err1 != nil {
		log.Fatal(err1)
	}

	defer f.Close()
	for i := 0; i < 2; i++ {
		_, err1 := io.WriteString(f, <-c) // write logs to the file
		if err1 != nil {
			panic(err1)
		}
	}
}
