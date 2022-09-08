package main

import (
	"io"
	"log"
	"net/rpc"
	"os"
	"strconv"
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

	totalSuccessNum := 0
	totalMatch := 0
	for _, server := range servers {
		wg.Add(1)       // add one when set a new task
		defer wg.Done() // minus one when finish a task

		// use goutine to execute concurrently
		go func(server utils.Server) { // connect to one server and try to execute RPC on that server
			client, err := rpc.Dial("tcp", server.IpAddr+":"+server.Port) // set connection
			if err != nil {
				handleError(err, c, &wg, server)
				return
			}

			var reply utils.ReplyStruct
			err = client.Call("grepLogService.GrepLog", "grep -Ec log "+server.FilePath+" "+server.Name+".log: ", &reply) // RPC
			if err != nil {
				handleError(err, c, &wg, server)
				return
			}
			c <- server.Name + ": " + reply.log // use channel send logs back
			if reply.ok {
				totalSuccessNum += 1
				match, err := strconv.Atoi(reply.log)
				if err != nil {
				} else {
					totalMatch += match
				}
			}
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
	_, err1 = io.WriteString(f, "match number: "+string(totalMatch)+"\n")
	if err1 != nil {
		panic(err1)
	}
	_, err1 = io.WriteString(f, "numer of successful log queries: "+string(totalSuccessNum)+"\n")
	if err1 != nil {
		panic(err1)
	}
}
