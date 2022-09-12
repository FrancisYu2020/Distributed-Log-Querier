package distributedTest

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
var VM1 string = "hangy6@fa22-cs425-2201.cs.illinois.edu"
var VM2 string = "hangy6@fa22-cs425-2202.cs.illinois.edu"
var VM3 string = "hangy6@fa22-cs425-2203.cs.illinois.edu"
var VM4 string = "hangy6@fa22-cs425-2204.cs.illinois.edu"
var QUERY_EXPRESSION string = "w*ld???"
// var QUERY_EXPRESSION string = "'Hello world!\nIs go the best???'"
// var QUERY_EXPRESSION string = "Hello"
// var QUERY_EXPRESSION string = "???"

func TestDistributedServerEmpty(t *testing.T) {
	// Test query log from server on local machine
	// Basic sanity check on client log collection. Check query expression in empty log on the local server

	// build the servers and clients
	server_compile_cmd := exec.Command("ssh", VM1, "sh mp1-hangy6-tian23/src/scripts/build.sh")
	server_compile_cmd.Start()
	server_compile_cmd = exec.Command("ssh", VM2, "sh mp1-hangy6-tian23/src/scripts/build.sh")
	server_compile_cmd.Start()
	server_compile_cmd = exec.Command("ssh", VM3, "sh mp1-hangy6-tian23/src/scripts/build.sh")
	server_compile_cmd.Start()
	server_compile_cmd = exec.Command("ssh", VM4, "sh mp1-hangy6-tian23/src/scripts/build.sh")
	server_compile_cmd.Start()
	stdoutStderr, err := server_compile_cmd.CombinedOutput()

	time.Sleep(1 * time.Second)
	// start the server
	server_cmd := exec.Command("ssh", VM1, BIN_PATH_SERVER + " &")
	err = server_cmd.Start()
	server_cmd = exec.Command("ssh", VM2, BIN_PATH_SERVER + " &")
	err = server_cmd.Start()
	server_cmd = exec.Command("ssh", VM3, BIN_PATH_SERVER + " &")
	err = server_cmd.Start()
	server_cmd = exec.Command("ssh", VM4, BIN_PATH_SERVER + " &")
	err = server_cmd.Start()
	if err != nil {
		t.Errorf("Failed to start server! %s", stdoutStderr)
		return
	}

	time.Sleep(1 * time.Second)

	// build the client
	client_complie_cmd := exec.Command("go", "build", "-o", BIN_PATH_CLIENT, SRC_PATH + "client_main.go")
	stdoutStderr, err = client_complie_cmd.CombinedOutput()
	if err != nil {
		t.Errorf("Failed to build client executable! %s", stdoutStderr)
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
	if last_line[len(last_line) - 1] != '4' {
		t.Errorf("Number of successful log queries does not match! Expected 4, got %c", last_line[len(last_line) - 1])
		return
	}

	// check if the total number of word count is correct
	second_last_line := log_lines[len(log_lines) - 3]
	if second_last_line[len(second_last_line) - 1] != '0' {
		t.Errorf("Number of total word count does not match! Expected 0, got %c", second_last_line[len(second_last_line) - 1])
		return
	}

	// check if only one server is giving expected results
	var pass int
	for _, line := range log_lines {
		if len(line) <= 6 || line[len(line) - 1:] != "0" { continue }
		if line[3:5] == "01" || line[3:5] == "02" || line[3:5] == "03" || line[3:5] == "04" {
			pass ++
		} 
	}
	if pass != 4 {
		t.Errorf("Collected log not correct! Log content is as below: %s", log_content)
		return
	}
}

func TestDistributedServerMini(t *testing.T) {
	// Test query log from server on local machine
	// Basic sanity check on client log collection. Check query expression in empty log on the local server

	// build the servers and clients
	server_compile_cmd := exec.Command("ssh", VM1, "sh mp1-hangy6-tian23/src/scripts/build.sh")
	server_compile_cmd.Start()
	server_compile_cmd = exec.Command("ssh", VM2, "sh mp1-hangy6-tian23/src/scripts/build.sh")
	server_compile_cmd.Start()
	server_compile_cmd = exec.Command("ssh", VM3, "sh mp1-hangy6-tian23/src/scripts/build.sh")
	server_compile_cmd.Start()
	server_compile_cmd = exec.Command("ssh", VM4, "sh mp1-hangy6-tian23/src/scripts/build.sh")
	server_compile_cmd.Start()
	stdoutStderr, err := server_compile_cmd.CombinedOutput()

	time.Sleep(1 * time.Second)
	// start the server
	server_cmd := exec.Command("ssh", VM1, BIN_PATH_SERVER + " &")
	err = server_cmd.Start()
	server_cmd = exec.Command("ssh", VM2, BIN_PATH_SERVER + " &")
	err = server_cmd.Start()
	server_cmd = exec.Command("ssh", VM3, BIN_PATH_SERVER + " &")
	err = server_cmd.Start()
	server_cmd = exec.Command("ssh", VM4, BIN_PATH_SERVER + " &")
	err = server_cmd.Start()
	if err != nil {
		t.Errorf("Failed to start server! %s", stdoutStderr)
		return
	}

	time.Sleep(1 * time.Second)

	// build the client
	client_complie_cmd := exec.Command("go", "build", "-o", BIN_PATH_CLIENT, SRC_PATH + "client_main.go")
	stdoutStderr, err = client_complie_cmd.CombinedOutput()
	if err != nil {
		t.Errorf("Failed to build client executable! %s", stdoutStderr)
		return
	}

	// execute the grep command from the client
	start := time.Now()
	cmd := exec.Command(BIN_PATH_CLIENT, "grep", "-c", QUERY_EXPRESSION, "mini.log")
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
	if last_line[len(last_line) - 1] != '4' {
		t.Errorf("Number of successful log queries does not match! Expected 4, got %c", last_line[len(last_line) - 1])
		return
	}

	// check if the total number of word count is correct
	second_last_line := log_lines[len(log_lines) - 3]
	if second_last_line[len(second_last_line) - 2:] != "12" {
		t.Errorf("Number of total word count does not match! Expected 12, got %s", second_last_line[len(second_last_line) - 2:])
		return
	}

	// check if only one server is giving expected results
	var pass int
	for _, line := range log_lines {
		if len(line) <= 6 || line[len(line) - 1:] != "3" { continue }
		if line[3:5] == "01" || line[3:5] == "02" || line[3:5] == "03" || line[3:5] == "04" {
			pass ++
		} 
	}
	if pass != 4 {
		t.Errorf("Collected log not correct! Log content is as below: %s", log_content)
		return
	}
}

func TestDistributedServerSmall(t *testing.T) {
	// Test query log from server on local machine
	// Basic sanity check on client log collection. Check query expression in empty log on the local server

	// build the servers and clients
	server_compile_cmd := exec.Command("ssh", VM1, "sh mp1-hangy6-tian23/src/scripts/build.sh")
	server_compile_cmd.Start()
	server_compile_cmd = exec.Command("ssh", VM2, "sh mp1-hangy6-tian23/src/scripts/build.sh")
	server_compile_cmd.Start()
	server_compile_cmd = exec.Command("ssh", VM3, "sh mp1-hangy6-tian23/src/scripts/build.sh")
	server_compile_cmd.Start()
	server_compile_cmd = exec.Command("ssh", VM4, "sh mp1-hangy6-tian23/src/scripts/build.sh")
	server_compile_cmd.Start()
	stdoutStderr, err := server_compile_cmd.CombinedOutput()

	time.Sleep(1 * time.Second)
	// start the server
	server_cmd := exec.Command("ssh", VM1, BIN_PATH_SERVER + " &")
	err = server_cmd.Start()
	server_cmd = exec.Command("ssh", VM2, BIN_PATH_SERVER + " &")
	err = server_cmd.Start()
	server_cmd = exec.Command("ssh", VM3, BIN_PATH_SERVER + " &")
	err = server_cmd.Start()
	server_cmd = exec.Command("ssh", VM4, BIN_PATH_SERVER + " &")
	err = server_cmd.Start()
	if err != nil {
		t.Errorf("Failed to start server! %s", stdoutStderr)
		return
	}

	time.Sleep(1 * time.Second)

	// build the client
	client_complie_cmd := exec.Command("go", "build", "-o", BIN_PATH_CLIENT, SRC_PATH + "client_main.go")
	stdoutStderr, err = client_complie_cmd.CombinedOutput()
	if err != nil {
		t.Errorf("Failed to build client executable! %s", stdoutStderr)
		return
	}

	// execute the grep command from the client
	start := time.Now()
	cmd := exec.Command(BIN_PATH_CLIENT, "grep", "-c", QUERY_EXPRESSION, "small.log")
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
	if last_line[len(last_line) - 1] != '4' {
		t.Errorf("Number of successful log queries does not match! Expected 4, got %c", last_line[len(last_line) - 1])
		return
	}

	// check if the total number of word count is correct
	second_last_line := log_lines[len(log_lines) - 3]
	if second_last_line[len(second_last_line) - 3:] != "120" {
		t.Errorf("Number of total word count does not match! Expected 120, got %s", second_last_line[len(second_last_line) - 3:])
		return
	}

	// check if only one server is giving expected results
	var pass int
	for _, line := range log_lines {
		if len(line) <= 6 || line[len(line) - 2:] != "30" { continue }
		if line[3:5] == "01" || line[3:5] == "02" || line[3:5] == "03" || line[3:5] == "04" {
			pass ++
		} 
	}
	if pass != 4 {
		t.Errorf("Collected log not correct! Log content is as below: %s", log_content)
		return
	}
}

func TestDistributedServerMedium(t *testing.T) {
	// Test query log from server on local machine
	// Basic sanity check on client log collection. Check query expression in empty log on the local server

	// build the servers and clients
	server_compile_cmd := exec.Command("ssh", VM1, "sh mp1-hangy6-tian23/src/scripts/build.sh")
	server_compile_cmd.Start()
	server_compile_cmd = exec.Command("ssh", VM2, "sh mp1-hangy6-tian23/src/scripts/build.sh")
	server_compile_cmd.Start()
	server_compile_cmd = exec.Command("ssh", VM3, "sh mp1-hangy6-tian23/src/scripts/build.sh")
	server_compile_cmd.Start()
	server_compile_cmd = exec.Command("ssh", VM4, "sh mp1-hangy6-tian23/src/scripts/build.sh")
	server_compile_cmd.Start()
	stdoutStderr, err := server_compile_cmd.CombinedOutput()

	time.Sleep(1 * time.Second)
	// start the server
	server_cmd := exec.Command("ssh", VM1, BIN_PATH_SERVER + " &")
	err = server_cmd.Start()
	server_cmd = exec.Command("ssh", VM2, BIN_PATH_SERVER + " &")
	err = server_cmd.Start()
	server_cmd = exec.Command("ssh", VM3, BIN_PATH_SERVER + " &")
	err = server_cmd.Start()
	server_cmd = exec.Command("ssh", VM4, BIN_PATH_SERVER + " &")
	err = server_cmd.Start()
	if err != nil {
		t.Errorf("Failed to start server! %s", stdoutStderr)
		return
	}

	time.Sleep(1 * time.Second)

	// build the client
	client_complie_cmd := exec.Command("go", "build", "-o", BIN_PATH_CLIENT, SRC_PATH + "client_main.go")
	stdoutStderr, err = client_complie_cmd.CombinedOutput()
	if err != nil {
		t.Errorf("Failed to build client executable! %s", stdoutStderr)
		return
	}

	// execute the grep command from the client
	start := time.Now()
	cmd := exec.Command(BIN_PATH_CLIENT, "grep", "-c", QUERY_EXPRESSION, "medium.log")
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
	if last_line[len(last_line) - 1] != '4' {
		t.Errorf("Number of successful log queries does not match! Expected 4, got %c", last_line[len(last_line) - 1])
		return
	}

	// check if the total number of word count is correct
	second_last_line := log_lines[len(log_lines) - 3]
	if second_last_line[len(second_last_line) - 4:] != "1200" {
		t.Errorf("Number of total word count does not match! Expected 1200, got %s", second_last_line[len(second_last_line) - 4:])
		return
	}

	// check if only one server is giving expected results
	var pass int
	for _, line := range log_lines {
		if len(line) <= 6 || line[len(line) - 3:] != "300" { continue }
		if line[3:5] == "01" || line[3:5] == "02" || line[3:5] == "03" || line[3:5] == "04" {
			pass ++
		} 
	}
	if pass != 4 {
		t.Errorf("Collected log not correct! Log content is as below: %s", log_content)
		return
	}
}

func TestDistributedServerLarge(t *testing.T) {
	// Test query log from server on local machine
	// Basic sanity check on client log collection. Check query expression in empty log on the local server

	// build the servers and clients
	server_compile_cmd := exec.Command("ssh", VM1, "sh mp1-hangy6-tian23/src/scripts/build.sh")
	server_compile_cmd.Start()
	server_compile_cmd = exec.Command("ssh", VM2, "sh mp1-hangy6-tian23/src/scripts/build.sh")
	server_compile_cmd.Start()
	server_compile_cmd = exec.Command("ssh", VM3, "sh mp1-hangy6-tian23/src/scripts/build.sh")
	server_compile_cmd.Start()
	server_compile_cmd = exec.Command("ssh", VM4, "sh mp1-hangy6-tian23/src/scripts/build.sh")
	server_compile_cmd.Start()
	stdoutStderr, err := server_compile_cmd.CombinedOutput()

	time.Sleep(1 * time.Second)
	// start the server
	server_cmd := exec.Command("ssh", VM1, BIN_PATH_SERVER + " &")
	err = server_cmd.Start()
	server_cmd = exec.Command("ssh", VM2, BIN_PATH_SERVER + " &")
	err = server_cmd.Start()
	server_cmd = exec.Command("ssh", VM3, BIN_PATH_SERVER + " &")
	err = server_cmd.Start()
	server_cmd = exec.Command("ssh", VM4, BIN_PATH_SERVER + " &")
	err = server_cmd.Start()
	if err != nil {
		t.Errorf("Failed to start server! %s", stdoutStderr)
		return
	}

	time.Sleep(1 * time.Second)

	// build the client
	client_complie_cmd := exec.Command("go", "build", "-o", BIN_PATH_CLIENT, SRC_PATH + "client_main.go")
	stdoutStderr, err = client_complie_cmd.CombinedOutput()
	if err != nil {
		t.Errorf("Failed to build client executable! %s", stdoutStderr)
		return
	}

	// execute the grep command from the client
	start := time.Now()
	cmd := exec.Command(BIN_PATH_CLIENT, "grep", "-c", QUERY_EXPRESSION, "large.log")
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
	if last_line[len(last_line) - 1] != '4' {
		t.Errorf("Number of successful log queries does not match! Expected 4, got %c", last_line[len(last_line) - 1])
		return
	}

	// check if the total number of word count is correct
	second_last_line := log_lines[len(log_lines) - 3]
	if second_last_line[len(second_last_line) - 6:] != "120000" {
		t.Errorf("Number of total word count does not match! Expected 120000, got %s", second_last_line[len(second_last_line) - 6:])
		return
	}

	// check if only one server is giving expected results
	var pass int
	for _, line := range log_lines {
		if len(line) <= 6 || line[len(line) - 5:] != "30000" { continue }
		if line[3:5] == "01" || line[3:5] == "02" || line[3:5] == "03" || line[3:5] == "04" {
			pass ++
		} 
	}
	if pass != 4 {
		t.Errorf("Collected log not correct! Log content is as below: %s", log_content)
		return
	}
}

func TestDistributedServerHuge(t *testing.T) {
	// Test query log from server on local machine
	// Basic sanity check on client log collection. Check query expression in empty log on the local server

	// build the servers and clients
	server_compile_cmd := exec.Command("ssh", VM1, "sh mp1-hangy6-tian23/src/scripts/build.sh")
	server_compile_cmd.Start()
	server_compile_cmd = exec.Command("ssh", VM2, "sh mp1-hangy6-tian23/src/scripts/build.sh")
	server_compile_cmd.Start()
	server_compile_cmd = exec.Command("ssh", VM3, "sh mp1-hangy6-tian23/src/scripts/build.sh")
	server_compile_cmd.Start()
	server_compile_cmd = exec.Command("ssh", VM4, "sh mp1-hangy6-tian23/src/scripts/build.sh")
	server_compile_cmd.Start()
	stdoutStderr, err := server_compile_cmd.CombinedOutput()

	time.Sleep(1 * time.Second)
	// start the server
	server_cmd := exec.Command("ssh", VM1, BIN_PATH_SERVER + " &")
	err = server_cmd.Start()
	server_cmd = exec.Command("ssh", VM2, BIN_PATH_SERVER + " &")
	err = server_cmd.Start()
	server_cmd = exec.Command("ssh", VM3, BIN_PATH_SERVER + " &")
	err = server_cmd.Start()
	server_cmd = exec.Command("ssh", VM4, BIN_PATH_SERVER + " &")
	err = server_cmd.Start()
	if err != nil {
		t.Errorf("Failed to start server! %s", stdoutStderr)
		return
	}

	time.Sleep(1 * time.Second)

	// build the client
	client_complie_cmd := exec.Command("go", "build", "-o", BIN_PATH_CLIENT, SRC_PATH + "client_main.go")
	stdoutStderr, err = client_complie_cmd.CombinedOutput()
	if err != nil {
		t.Errorf("Failed to build client executable! %s", stdoutStderr)
		return
	}

	// execute the grep command from the client
	start := time.Now()
	cmd := exec.Command(BIN_PATH_CLIENT, "grep", "-c", QUERY_EXPRESSION, "huge.log")
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
	if last_line[len(last_line) - 1] != '4' {
		t.Errorf("Number of successful log queries does not match! Expected 4, got %c", last_line[len(last_line) - 1])
		return
	}

	// check if the total number of word count is correct
	second_last_line := log_lines[len(log_lines) - 3]
	if second_last_line[len(second_last_line) - 7:] != "1200000" {
		t.Errorf("Number of total word count does not match! Expected 1200000, got %s", second_last_line[len(second_last_line) - 7:])
		return
	}

	// check if only one server is giving expected results
	var pass int
	for _, line := range log_lines {
		if len(line) <= 6 || line[len(line) - 6:] != "300000" { continue }
		if line[3:5] == "01" || line[3:5] == "02" || line[3:5] == "03" || line[3:5] == "04" {
			pass ++
		} 
	}
	if pass != 4 {
		t.Errorf("Collected log not correct! Log content is as below: %s", log_content)
		return
	}
}