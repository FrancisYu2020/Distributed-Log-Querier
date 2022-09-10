package utils

import (
	"os/exec"
	"strings"
)

// get the command and execute it by command line, return the grep result
func Grep(command string) (string, bool) {
	params := strings.Split(command, " ")
	var cmd *exec.Cmd
	// execute the grep command
	if len(params) == 5 { // grep -Ec [regex] *.log [log file path]
		cmd = exec.Command(params[0], params[1], params[2], params[4]+params[3])
	} else if len(params) == 6 { // grep -Ec [regex] *.log [output path] [log file path]
		cmd = exec.Command(params[0], params[1], params[2], params[5]+params[3])
	}

	data, err := cmd.CombinedOutput() // get the result
	if err != nil {                   // handle error
		return "failed to call command: " + err.Error(), false
	}
	return string(data), true
}
