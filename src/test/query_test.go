package test

import (
	"testing"
	// "bytes"
	"fmt"
	"os/exec"
	// "time"
	// "reflect"
	"strings"
)

var BIN_PATH_CLIENT string = "/home/hangy6/mp1-hangy6-tian23/bin/client"
var BIN_PATH_SERVER string = "/home/hangy6/mp1-hangy6-tian23/bin/server"
var SRC_PATH string = "/home/hangy6/mp1-hangy6-tian23/src/"

func TestBasic1(t *testing.T) {
	// Test query log from server on local machine
	// Basic sanity check on client log collection. Check word "Hello" in log "log1" on each server

	// build the server
	server_compile_cmd := exec.Command("go", "build", "-o", BIN_PATH_SERVER, SRC_PATH + "server_main.go")
	stdoutStderr, err := server_compile_cmd.CombinedOutput()
	if err != nil {
		t.Errorf("Failed to build server executable! %s", stdoutStderr)
	}
	
	// build the client
	client_complie_cmd := exec.Command("go", "build", "-o", BIN_PATH_CLIENT, SRC_PATH + "client_main.go")
	stdoutStderr, err = client_complie_cmd.CombinedOutput()
	if err != nil {
		t.Errorf("Failed to build client executable! %s", stdoutStderr)
	}

	// start the server
	server_cmd := exec.Command(BIN_PATH_SERVER, " &")
	err = server_cmd.Start()
	if err != nil {
		t.Errorf("Failed to start server! %s", stdoutStderr)
	}

	// execute the grep command from the client
	cmd := exec.Command(BIN_PATH_CLIENT, "grep", "-c", "Hello", "log1")
	stdoutStderr, err = cmd.CombinedOutput()
	if err != nil {
		t.Error(stdoutStderr)
	}

	// check the log output in the following blocks
	log_content := string(fmt.Sprintf("%s", stdoutStderr))
	log_lines := strings.Split(log_content, "\n") //NOTE that the last line is empty line

	// check if the number of succeed log queries match
	last_line := log_lines[len(log_lines) - 2]
	if last_line[len(last_line) - 1] != '1' {
		t.Errorf("Number of successful log queries does not match! Expected 1, got %s", log_lines[len(log_lines) - 2])
	}

	// check if the total number of word count is correct
	second_last_line := log_lines[len(log_lines) - 3]
	if second_last_line[len(second_last_line) - 1] != '1' {
		t.Errorf("Number of total word count does not match! Expected 1, got %s", log_lines[len(log_lines) - 1])
	}

	// check if only one server is giving expected results
	var pass bool
	for _, line := range log_lines {
		if len(line) > 0 && line[len(line) - 1] == '1' {
			pass = true
		}
	}
	if !pass {
		t.Errorf("Number of word count is not correct! Log content is as below: %s", log_content)
	}
}