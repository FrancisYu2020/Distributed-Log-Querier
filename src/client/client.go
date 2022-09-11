package client

import (
	// "bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/rpc"
	"os"
	"strconv"
	"strings"
	"sync"
	utils "src/utils"
)

type replyStruct struct {
	Log string `json:log`
	Ok  bool   `json:ok`
}

// Outer method wrappers
func CheckFileIsExist(filename string) bool { return checkFileIsExist(filename) }
func WriteFile(filename string, c chan string, taskNum int, totalMatch *int, totalSuccessNum *int) { writeFile(filename, c, taskNum,totalMatch, totalSuccessNum) }
func PrintQueryResult(taskNum int, c chan string, totalMatch *int, totalSuccessNum *int) { printQueryResult(taskNum, c, totalMatch, totalSuccessNum) }

// check wheter 'filename' file exists
func checkFileIsExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

func writeFile(filename string, c chan string, taskNum int, totalMatch *int, totalSuccessNum *int) {
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
	for i := 0; i < taskNum; i++ {
		_, err1 := io.WriteString(f, <-c) // write logs to the file
		if err1 != nil {
			panic(err1)
		}
	}
	_, err1 = io.WriteString(f, "total match number: "+strconv.Itoa(*totalMatch)+"\n")
	if err1 != nil {
		panic(err1)
	}
	_, err1 = io.WriteString(f, "number of successful log queries: "+strconv.Itoa(*totalSuccessNum)+"\n")
	if err1 != nil {
		panic(err1)
	}
	fmt.Println("All tasks done!")
	fmt.Println("Please see " + filename + " for results.")
}

func printQueryResult(taskNum int, c chan string, totalMatch *int, totalSuccessNum *int) {
	for i := 0; i < taskNum; i++ {
		fmt.Print(<-c) // output the results
	}
	fmt.Println("total match number: " + strconv.Itoa(*totalMatch))
	fmt.Println("number of successful log queries: " + strconv.Itoa(*totalSuccessNum))
}

func handleError(err error, c chan string, wg *sync.WaitGroup, server utils.Server) {
	c <- string(server.Name + ".log: " + err.Error() + "\n")
}

func ClientMain() {
	var wg sync.WaitGroup // use wait group to keep synchronization
	defer wg.Wait()

	c := make(chan string) // use chanel to send logs safely
	servers := utils.LoadConfig()

	// fmt.Println("Please enter the query...")

    argsWithoutProg := os.Args[1:]
	query := strings.Join(argsWithoutProg, " ")

	totalSuccessNum := 0
	totalMatch := 0
	// fmt.Println("Querying log...")
	for _, server := range servers {
		wg.Add(1) // add one when set a new task
		// use goutine to execute concurrently
		go func(server utils.Server, totalSuccessNum, totalMatch *int, query string) { // connect to one server and try to execute RPC on that server
			defer wg.Done()                                               // minus one when finish a task
			client, err := rpc.Dial("tcp", server.IpAddr+":"+server.Port) // set connection
			if err != nil {
				handleError(err, c, &wg, server)
				return
			}

			var reply string
			command := query + " " + server.FilePath
			err = client.Call("grepLogService.GrepLog", command, &reply) // RPC
			if err != nil {
				handleError(err, c, &wg, server)
				return
			}


			var message replyStruct
			json.Unmarshal([]byte(reply), &message)

			c <- server.Name + ": " + message.Log // use channel send logs back
			if message.Ok {
				*totalSuccessNum += 1
				match, err := strconv.Atoi(message.Log[:len(message.Log)-1])
				if err != nil {
				} else {
					*totalMatch += match
				}
			}
		}(server, &totalSuccessNum, &totalMatch, query)
	}

	queryParmas := strings.Split(query, " ")
	if len(queryParmas) == 4 { // grep -Ec [regex] *.log
		printQueryResult(len(servers), c, &totalMatch, &totalSuccessNum)
	} else if len(queryParmas) == 5 { // grep -Ec [regex] *.log [output path]
		var filename = queryParmas[4] // path of the log file
		writeFile(filename, c, len(servers), &totalMatch, &totalSuccessNum)
	}
}
