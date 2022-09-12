package main

// main wrapper for the server package, for the covenience of working in scripts
import (
	server "src/server"
)

func main() {
	server.ServerMain()
}
