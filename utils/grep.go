package utils

import (
	"log"
	"os/exec"
	"strings"
)

func Grep(command string) string {
	params := strings.Split(command, " ")
	cmd := exec.Command("grep", params[1], params[2], params[3])
	data, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("failed to call cmd.Run(): %v", err)
	}
	return string(data)
}
