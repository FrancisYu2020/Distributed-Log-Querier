package main

import (
	"distributed-log/utils"
	"fmt"
	"log"
	"net"
	"net/rpc"
)

type grepLogService struct{}

func (p *grepLogService) GrepLog(request string, reply *string) error {
	fmt.Printf("grep command：%v\n", request)

	log := utils.Grep(request)
	*reply = "This is the log\n" + log
	return nil
}

func makeServer(port string) {
	rpc.RegisterName("grepLogService", new(grepLogService))

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error:", err)
	}

	rpc.ServeConn(conn)

}

func main() {
	go makeServer(":1234")
	go makeServer(":1235")
	for {
	}
}
