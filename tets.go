package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	cmd := exec.Command("grep", "-Ec", "das", "src/test_logs/log1")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run() // get the result
	if err != nil {  // handle error
		fmt.Print("command error: " + stderr.String())
	}
	fmt.Print(strings.Compare(stdout.String(), "0") == 1)
}
