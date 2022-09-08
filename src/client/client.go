package main

import (
	"io"
	"log"
	"net/rpc"
	"os"
	"sync"
	"utils"
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

func handleError(err error, c chan string, wg *sync.WaitGroup, server utils.Server) {
	defer wg.Done()
	c <- string(server.Name + ".log: " + err.Error() + "\n")
}

func main() {
	var wg sync.WaitGroup // use wait group to keep synchronization
	defer wg.Wait()

	c := make(chan string) // use chanel to send logs safely
	servers := utils.LoadConfig()

	for _, server := range servers {
		wg.Add(1) // add one when set a new task
		// use goutine to execute concurrently
		go func(server utils.Server) { // connect to one server and try to execute RPC on that server
			client, err := rpc.Dial("tcp", server.IpAddr+":"+server.Port) // set connection
			if err != nil {
				handleError(err, c, &wg, server)
				return
			}

			var reply string
			err = client.Call("grepLogService.GrepLog", "grep -Ec log ../test_logs/log1 "+server.Name+".log: ", &reply) // RPC
			if err != nil {
				handleError(err, c, &wg, server)
				return
			}
			c <- reply // use channel send logs back

			wg.Done() // minus one when finish a task
		}(server)
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
	for i := 0; i < len(servers); i++ {
		_, err1 := io.WriteString(f, <-c) // write logs to the file
		if err1 != nil {
			panic(err1)
		}
	}
}
