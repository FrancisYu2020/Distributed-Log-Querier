package localTest

import (
	"testing"
	"fmt"
	"os/exec"
	"strings"
	"time"
	"log"
)

var BIN_PATH_CLIENT string = "/home/hangy6/mp1-hangy6-tian23/bin/client"
var BIN_PATH_SERVER string = "/home/hangy6/mp1-hangy6-tian23/bin/server"
var SRC_PATH string = "/home/hangy6/mp1-hangy6-tian23/src/"
// var QUERY_EXPRESSION string = "\"w*ld???\""
var QUERY_EXPRESSION string = "w*ld???"
// var QUERY_EXPRESSION string = "'Hello world!\nIs go the best???'"
// var QUERY_EXPRESSION string = "\"Hello\""
// var QUERY_EXPRESSION string = "???"

func TestOneServerEmpty(t *testing.T) {
	// Test query log from server on local machine
	// Basic sanity check on client log collection. Check query expression in empty log on the local server

	// build the server
	server_compile_cmd := exec.Command("go", "build", "-o", BIN_PATH_SERVER, SRC_PATH + "server_main.go")
	stdoutStderr, err := server_compile_cmd.CombinedOutput()
	if err != nil {
		t.Errorf("Failed to build server executable! %s", stdoutStderr)
		return
	}
	
	// build the client
	client_complie_cmd := exec.Command("go", "build", "-o", BIN_PATH_CLIENT, SRC_PATH + "client_main.go")
	stdoutStderr, err = client_complie_cmd.CombinedOutput()
	if err != nil {
		t.Errorf("Failed to build client executable! %s", stdoutStderr)
		return
	}

	// start the server
	server_cmd := exec.Command(BIN_PATH_SERVER, " &")
	err = server_cmd.Start()
	if err != nil {
		t.Errorf("Failed to start server! %s", stdoutStderr)
		return
	}

	// execute the grep command from the client
	start := time.Now()
	cmd := exec.Command(BIN_PATH_CLIENT, "grep", "-c", QUERY_EXPRESSION, "empty.log")
	elapsed := time.Since(start)
	log.Printf("%s execution time is: %s", cmd, elapsed)
	stdoutStderr, err = cmd.CombinedOutput()
	if err != nil {
		t.Error(stdoutStderr)
		return
	}

	// check the log output in the following blocks
	log_content := string(fmt.Sprintf("%s", stdoutStderr))
	log_lines := strings.Split(log_content, "\n") //NOTE that the last line is empty line

	// check if the number of succeed log queries match
	last_line := log_lines[len(log_lines) - 2]
	if last_line[len(last_line) - 1] != '1' {
		t.Errorf("Number of successful log queries does not match! Expected 1, got %c", last_line[len(last_line) - 1])
		return
	}

	// check if the total number of word count is correct
	second_last_line := log_lines[len(log_lines) - 3]
	if second_last_line[len(second_last_line) - 1] != '0' {
		t.Errorf("Number of total word count does not match! Expected 0, got %c", second_last_line[len(second_last_line) - 1])
		return
	}

	// check if only one server is giving expected results
	var pass bool
	for _, line := range log_lines {
		if len(line) > 0 && line[len(line) - 1] == '0' {
			pass = true
		}
	}
	if !pass {
		t.Errorf("Number of word count is not correct! Log content is as below: %s", log_content)
		return
	}
}

func TestOneServerMini(t *testing.T) {
	// Test query log from server on local machine
	// Basic sanity check on client log collection. Check word "Hello" in log "small.log" on the local server

	// build the server
	server_compile_cmd := exec.Command("go", "build", "-o", BIN_PATH_SERVER, SRC_PATH + "server_main.go")
	stdoutStderr, err := server_compile_cmd.CombinedOutput()
	if err != nil {
		t.Errorf("Failed to build server executable! %s", stdoutStderr)
		return
	}
	
	// build the client
	client_complie_cmd := exec.Command("go", "build", "-o", BIN_PATH_CLIENT, SRC_PATH + "client_main.go")
	stdoutStderr, err = client_complie_cmd.CombinedOutput()
	if err != nil {
		t.Errorf("Failed to build client executable! %s", stdoutStderr)
		return
	}

	// start the server
	server_cmd := exec.Command(BIN_PATH_SERVER, " &")
	err = server_cmd.Start()
	if err != nil {
		t.Errorf("Failed to start server! %s", stdoutStderr)
		return
	}

	// execute the grep command from the client
	start := time.Now()
	cmd := exec.Command(BIN_PATH_CLIENT, "grep", "-c", QUERY_EXPRESSION, "mini.log")
	elapsed := time.Since(start)
	log.Printf("%s execution time is: %s", cmd, elapsed)
	stdoutStderr, err = cmd.CombinedOutput()
	log.Printf("%s", stdoutStderr)
	if err != nil {
		t.Error(stdoutStderr)
		return
	}

	// check the log output in the following blocks
	log_content := string(fmt.Sprintf("%s", stdoutStderr))
	log_lines := strings.Split(log_content, "\n") //NOTE that the last line is empty line

	// check if the number of succeed log queries match
	last_line := log_lines[len(log_lines) - 2]
	if last_line[len(last_line) - 1] != '1' {
		t.Errorf("Number of successful log queries does not match! Expected 1, got %c", last_line[len(last_line) - 1])
		return
	}

	// check if the total number of word count is correct
	second_last_line := log_lines[len(log_lines) - 3]
	if second_last_line[len(second_last_line) - 1:] != "3" {
		t.Errorf("Number of total word count does not match! Expected 3, got %s", second_last_line[len(second_last_line) - 1:])
		return
	}

	// check if only one server is giving expected results
	var pass bool = false
	for _, line := range log_lines[:10] {
		if len(line) > 0 && line[len(line) - 1:] == "3" {
			pass = true
		}
	}
	if !pass {
		t.Errorf("Number of word count is not correct! Log content is as below: %s", log_content)
		return
	}
}

func TestOneServerSmall(t *testing.T) {
	// Test query log from server on local machine
	// Basic sanity check on client log collection. Check query expression in log "small.log" on the local server

	// build the server
	server_compile_cmd := exec.Command("go", "build", "-o", BIN_PATH_SERVER, SRC_PATH + "server_main.go")
	stdoutStderr, err := server_compile_cmd.CombinedOutput()
	if err != nil {
		t.Errorf("Failed to build server executable! %s", stdoutStderr)
		return
	}
	
	// build the client
	client_complie_cmd := exec.Command("go", "build", "-o", BIN_PATH_CLIENT, SRC_PATH + "client_main.go")
	stdoutStderr, err = client_complie_cmd.CombinedOutput()
	if err != nil {
		t.Errorf("Failed to build client executable! %s", stdoutStderr)
		return
	}

	// start the server
	server_cmd := exec.Command(BIN_PATH_SERVER, " &")
	err = server_cmd.Start()
	if err != nil {
		t.Errorf("Failed to start server! %s", stdoutStderr)
		return
	}

	// execute the grep command from the client
	start := time.Now()
	cmd := exec.Command(BIN_PATH_CLIENT, "grep", "-c", QUERY_EXPRESSION, "small.log")
	// log.Println(cmd)
	elapsed := time.Since(start)
	log.Printf("%s execution time is: %s", cmd, elapsed)
	stdoutStderr, err = cmd.CombinedOutput()
	log.Printf("%s", stdoutStderr)
	if err != nil {
		t.Error(stdoutStderr)
		return
	}

	// check the log output in the following blocks
	log_content := string(fmt.Sprintf("%s", stdoutStderr))
	log_lines := strings.Split(log_content, "\n") //NOTE that the last line is empty line

	// check if the number of succeed log queries match
	last_line := log_lines[len(log_lines) - 2]
	if last_line[len(last_line) - 1] != '1' {
		t.Errorf("Number of successful log queries does not match! Expected 1, got %c", last_line[len(last_line) - 1])
		return
	}

	// check if the total number of word count is correct
	second_last_line := log_lines[len(log_lines) - 3]
	if second_last_line[len(second_last_line) - 2:] != "30" {
		t.Errorf("Number of total word count does not match! Expected 30, got %s", second_last_line[len(second_last_line) - 2:])
		return
	}

	// check if only one server is giving expected results
	var pass bool
	for _, line := range log_lines[:10] {
		if len(line) > 1 && line[len(line) - 2:] == "30" {
			pass = true
		}
	}
	if !pass {
		t.Errorf("Number of word count is not correct! Log content is as below: %s", log_content)
		return
	}
}

func TestOneServerMedium(t *testing.T) {
	// Test query log from server on local machine
	// Basic sanity check on client log collection. Check query expression in log "small.log" on the local server

	// build the server
	server_compile_cmd := exec.Command("go", "build", "-o", BIN_PATH_SERVER, SRC_PATH + "server_main.go")
	stdoutStderr, err := server_compile_cmd.CombinedOutput()
	if err != nil {
		t.Errorf("Failed to build server executable! %s", stdoutStderr)
		return
	}
	
	// build the client
	client_complie_cmd := exec.Command("go", "build", "-o", BIN_PATH_CLIENT, SRC_PATH + "client_main.go")
	stdoutStderr, err = client_complie_cmd.CombinedOutput()
	if err != nil {
		t.Errorf("Failed to build client executable! %s", stdoutStderr)
		return
	}

	// start the server
	server_cmd := exec.Command(BIN_PATH_SERVER, " &")
	err = server_cmd.Start()
	if err != nil {
		t.Errorf("Failed to start server! %s", stdoutStderr)
		return
	}

	// execute the grep command from the client
	start := time.Now()
	cmd := exec.Command(BIN_PATH_CLIENT, "grep", "-c", QUERY_EXPRESSION, "medium.log")
	// log.Println(cmd)
	elapsed := time.Since(start)
	log.Printf("%s execution time is: %s", cmd, elapsed)
	stdoutStderr, err = cmd.CombinedOutput()
	log.Printf("%s", stdoutStderr)
	if err != nil {
		t.Error(stdoutStderr)
		return
	}

	// check the log output in the following blocks
	log_content := string(fmt.Sprintf("%s", stdoutStderr))
	log_lines := strings.Split(log_content, "\n") //NOTE that the last line is empty line

	// check if the number of succeed log queries match
	last_line := log_lines[len(log_lines) - 2]
	if last_line[len(last_line) - 1] != '1' {
		t.Errorf("Number of successful log queries does not match! Expected 1, got %c", last_line[len(last_line) - 1])
		return
	}

	// check if the total number of word count is correct
	second_last_line := log_lines[len(log_lines) - 3]
	if second_last_line[len(second_last_line) - 3:] != "300" {
		t.Errorf("Number of total word count does not match! Expected 300, got %s", second_last_line[len(second_last_line) - 3:])
		return
	}

	// check if only one server is giving expected results
	var pass bool
	for _, line := range log_lines[:10] {
		if len(line) > 2 && line[len(line) - 3:] == "300" {
			pass = true
		}
	}
	if !pass {
		t.Errorf("Number of word count is not correct! Log content is as below: %s", log_content)
		return
	}
}

func TestOneServerLarge(t *testing.T) {
	// Test query log from server on local machine
	// Basic sanity check on client log collection. Check query expression in log "small.log" on the local server

	// build the server
	server_compile_cmd := exec.Command("go", "build", "-o", BIN_PATH_SERVER, SRC_PATH + "server_main.go")
	stdoutStderr, err := server_compile_cmd.CombinedOutput()
	if err != nil {
		t.Errorf("Failed to build server executable! %s", stdoutStderr)
		return
	}
	
	// build the client
	client_complie_cmd := exec.Command("go", "build", "-o", BIN_PATH_CLIENT, SRC_PATH + "client_main.go")
	stdoutStderr, err = client_complie_cmd.CombinedOutput()
	if err != nil {
		t.Errorf("Failed to build client executable! %s", stdoutStderr)
		return
	}

	// start the server
	server_cmd := exec.Command(BIN_PATH_SERVER, " &")
	err = server_cmd.Start()
	if err != nil {
		t.Errorf("Failed to start server! %s", stdoutStderr)
		return
	}

	// execute the grep command from the client
	start := time.Now()
	cmd := exec.Command(BIN_PATH_CLIENT, "grep", "-c", QUERY_EXPRESSION, "large.log")
	// log.Println(cmd)
	elapsed := time.Since(start)
	log.Printf("%s execution time is: %s", cmd, elapsed)
	stdoutStderr, err = cmd.CombinedOutput()
	log.Printf("%s", stdoutStderr)
	if err != nil {
		t.Error(stdoutStderr)
		return
	}

	// check the log output in the following blocks
	log_content := string(fmt.Sprintf("%s", stdoutStderr))
	log_lines := strings.Split(log_content, "\n") //NOTE that the last line is empty line

	// check if the number of succeed log queries match
	last_line := log_lines[len(log_lines) - 2]
	if last_line[len(last_line) - 1] != '1' {
		t.Errorf("Number of successful log queries does not match! Expected 1, got %c", last_line[len(last_line) - 1])
		return
	}

	// check if the total number of word count is correct
	second_last_line := log_lines[len(log_lines) - 3]
	if second_last_line[len(second_last_line) - 5:] != "30000" {
		t.Errorf("Number of total word count does not match! Expected 30000, got %s", second_last_line[len(second_last_line) - 5:])
		return
	}

	// check if only one server is giving expected results
	var pass bool
	for _, line := range log_lines[:10] {
		if len(line) > 2 && line[len(line) - 5:] == "30000" {
			pass = true
		}
	}
	if !pass {
		t.Errorf("Number of word count is not correct! Log content is as below: %s", log_content)
		return
	}
}

func TestOneServerHuge(t *testing.T) {
	// Test query log from server on local machine
	// Basic sanity check on client log collection. Check query expression in log "small.log" on the local server

	// build the server
	server_compile_cmd := exec.Command("go", "build", "-o", BIN_PATH_SERVER, SRC_PATH + "server_main.go")
	stdoutStderr, err := server_compile_cmd.CombinedOutput()
	if err != nil {
		t.Errorf("Failed to build server executable! %s", stdoutStderr)
		return
	}
	
	// build the client
	client_complie_cmd := exec.Command("go", "build", "-o", BIN_PATH_CLIENT, SRC_PATH + "client_main.go")
	stdoutStderr, err = client_complie_cmd.CombinedOutput()
	if err != nil {
		t.Errorf("Failed to build client executable! %s", stdoutStderr)
		return
	}

	// start the server
	server_cmd := exec.Command(BIN_PATH_SERVER, " &")
	err = server_cmd.Start()
	if err != nil {
		t.Errorf("Failed to start server! %s", stdoutStderr)
		return
	}

	// execute the grep command from the client
	start := time.Now()
	cmd := exec.Command(BIN_PATH_CLIENT, "grep", "-c", QUERY_EXPRESSION, "huge.log")
	// log.Println(cmd)
	elapsed := time.Since(start)
	log.Printf("%s execution time is: %s", cmd, elapsed)
	stdoutStderr, err = cmd.CombinedOutput()
	log.Printf("%s", stdoutStderr)
	if err != nil {
		t.Error(stdoutStderr)
		return
	}

	// check the log output in the following blocks
	log_content := string(fmt.Sprintf("%s", stdoutStderr))
	log_lines := strings.Split(log_content, "\n") //NOTE that the last line is empty line

	// check if the number of succeed log queries match
	last_line := log_lines[len(log_lines) - 2]
	if last_line[len(last_line) - 1] != '1' {
		t.Errorf("Number of successful log queries does not match! Expected 1, got %c", last_line[len(last_line) - 1])
		return
	}

	// check if the total number of word count is correct
	second_last_line := log_lines[len(log_lines) - 3]
	if second_last_line[len(second_last_line) - 6:] != "300000" {
		t.Errorf("Number of total word count does not match! Expected 30000, got %s", second_last_line[len(second_last_line) - 6:])
		return
	}

	// check if only one server is giving expected results
	var pass bool
	for _, line := range log_lines[:10] {
		if len(line) > 2 && line[len(line) - 6:] == "300000" {
			pass = true
		}
	}
	if !pass {
		t.Errorf("Number of word count is not correct! Log content is as below: %s", log_content)
		return
	}
}