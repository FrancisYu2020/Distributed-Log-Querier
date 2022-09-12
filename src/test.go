package main

import (
	"fmt"
	"os/exec"
)

var BIN_PATH_CLIENT string = "/home/hangy6/mp1-hangy6-tian23/bin/client"
var BIN_PATH_SERVER string = "/home/hangy6/mp1-hangy6-tian23/bin/server"
var SRC_PATH string = "/home/hangy6/mp1-hangy6-tian23/src/"
var VM1 string = "hangy6@fa22-cs425-2202.cs.illinois.edu"

func main() {
	// server_compile_cmd := exec.Command("ssh", VM1, "ls -al")
	// server_compile_cmd := exec.Command("ssh", VM1, "sh mp1-hangy6-tian23/src/scripts/build.sh")
	server_cmd := exec.Command("ssh", VM1, BIN_PATH_SERVER + " &")
	err := server_cmd.Start()
	fmt.Println(err)
	fmt.Println(server_cmd)
	// stdoutStderr, _ := server_compile_cmd.CombinedOutput()
	// fmt.Printf("%s", stdoutStderr)
	// cmd := server_compile_cmd.Run()
	// fmt.Println(cmd)
	// server_compile_cmd = exec.Command("echo \"hello\"")
	// stdoutStderr, _ = server_compile_cmd.CombinedOutput()
	// fmt.Println(stdoutStderr)
}