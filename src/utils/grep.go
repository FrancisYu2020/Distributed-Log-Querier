package utils

import (
	"os/exec"
	"strings"
)

// get the command and execute it by command line, return the grep result
func Grep(command string) string {
	params := strings.Split(command, " ")
	cmd := exec.Command("grep", params[1], params[2], params[3]) // execute the grep command
	data, err := cmd.CombinedOutput()                            // get the result
	if err != nil {                                              // handle error
		return "failed to call command: " + err.Error()
	}
	return params[4] + string(data)
}
