package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"utils"
)

type grepLogService struct{}

type replyStruct struct {
	log string
	ok  bool
}

func (p *grepLogService) GrepLog(request string, reply *string) error {
	fmt.Printf("grep command：%v\n", request) // print the request command

	log, ok := utils.Grep(request) // get the log query results
	// *reply = replyStruct{log, ok}  // send reply back to client
	*reply = log
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
