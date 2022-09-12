package utils

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// get the command and execute it by command line, return the grep result
func Grep(command string) (string, bool) {
	params := strings.Split(command, "\"")
	grep := strings.Split(params[0], " ")
	path := strings.Split(params[2][1:], " ")

	var cmd *exec.Cmd
	// execute the grep command
	if len(path) == 2 { // grep -Ec [regex] *.log [log file path]
		cmd = exec.Command(grep[0], grep[1], params[1], path[1]+path[0])
	} else if len(params) == 3 { // grep -Ec [regex] *.log [output path] [log file path]
		cmd = exec.Command(grep[0], grep[1], params[1], path[2]+path[0])
	}

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout                                          // use stdout to store the output
	cmd.Stderr = &stderr                                          // use stderr to store the error message
	err := cmd.Run()                                              // get the result
	if err != nil && strings.Compare(stdout.String(), "0") != 1 { // handle error if grep result is not 0
		return "command error: " + stderr.String(), false
	}
	return stdout.String(), true
}
