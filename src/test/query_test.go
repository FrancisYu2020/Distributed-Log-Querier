package test

import (
	"testing"
	// "bytes"
	"fmt"
	"os/exec"
	"time"
	// "reflect"
	// "strings"
)

var BIN_PATH_CLIENT string = "/home/hangy6/mp1-hangy6-tian23/bin/client"
var BIN_PATH_SERVER string = "/home/hangy6/mp1-hangy6-tian23/bin/server"
var SRC_PATH string = "/home/hangy6/mp1-hangy6-tian23/src/"

func TestBasic1(t *testing.T) {
	// Test query log from server on local machine
	// Basic sanity check on client log collection. Check word "Hello" in log "log1" on each server

	fmt.Println("start build server")
	// build the server
	server_compile_cmd := exec.Command("go", "build", "-o", BIN_PATH_SERVER, SRC_PATH + "server_main.go")
	stdoutStderr, err := server_compile_cmd.CombinedOutput()
	if err != nil {
		t.Errorf("Failed to build server executable! %s", stdoutStderr)
	}
	fmt.Println("finished build server")
	
	// build the client
	client_complie_cmd := exec.Command("go", "build", "-o", BIN_PATH_CLIENT, SRC_PATH + "client_main.go")
	stdoutStderr, err = client_complie_cmd.CombinedOutput()
	if err != nil {
		t.Errorf("Failed to build client executable! %s", stdoutStderr)
	}

	// start the server
	server_cmd := exec.Command("nohup", BIN_PATH_SERVER + ">/dev/null", "2>&1&")
	// server_cmd := exec.Command(BIN_PATH_SERVER, "&")
	// stdoutStderr, err = server_cmd.CombinedOutput()
	fmt.Println(err)
	err = server_cmd.Run()
	fmt.Println(err)
	if err != nil {
		t.Errorf("Failed to start server! %s", stdoutStderr)
	}

	time.Sleep(0 * time.Second)
	// execute the grep command from the client
	cmd := exec.Command(BIN_PATH_CLIENT, "grep", "-c", "Hello", "log1")
	stdoutStderr, err = cmd.CombinedOutput()
	if err != nil {
		t.Error(stdoutStderr)
	}

	log_content := fmt.Sprintf("%s", stdoutStderr)
	if log_content[len(log_content)-2] == '1' {
		t.Errorf("Number of successful log queries does not match! Expected 1, got %c", log_content[len(log_content)-2])
	}
}