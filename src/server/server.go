package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/rpc"
	utils "src/utils"
)

type replyStruct struct {
	Log string `json:log`
	Ok  bool   `json:ok`
}

type grepLogService struct{}

// Outer method wrappers
func OpenLogServer(port string) { openLogServer(port) }

func (p *grepLogService) GrepLog(request string, reply *string) error {
	fmt.Printf("grep command：%v\n", request) // print the request command

	log, ok := utils.Grep(request)    // get the log query results
	data := replyStruct{log, ok}      // use struct store log and successful message
	jsonData, _ := json.Marshal(data) // convert to json
	str := string(jsonData)
	*reply = str // send back to client

	return nil
}

func openLogServer(port string) {
	rpc.RegisterName("grepLogService", new(grepLogService)) // register RPC service
	listener, err := net.Listen("tcp", port)                // listen at particular port
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for { // keep listen
		conn, err := listener.Accept() // build connection with client
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		go rpc.ServeConn(conn) // provide RPC service
	}
}

func ServerMain() {
	openLogServer(":1234") // open server and keep listening at port 1234
}
