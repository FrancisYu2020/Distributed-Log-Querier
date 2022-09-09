package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"utils"
)

type replyStruct struct {
	Log string `json:log`
	Ok  bool   `json:ok`
}

type grepLogService struct{}

func (p *grepLogService) GrepLog(request string, reply *string) error {
	fmt.Printf("grep command：%v\n", request) // print the request command

	log, ok := utils.Grep(request) // get the log query results
	data := replyStruct{log, ok}
	jsonData, _ := json.Marshal(data)
	str := string(jsonData)
	fmt.Println(str)
	*reply = str

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

func main() {
	openLogServer(":1234") // open server and keep listening at port 1234
}
