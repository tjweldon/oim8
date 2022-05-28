package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"tjweldon/oim8/src/peer"
	"tjweldon/oim8/src/repl"
)

func main() {
	RunServer()
	// DoRepl()
}

// DoRepl initialises the REPL and runs it as the main function
func DoRepl() {
	// supply stdin to REPL
	reader := bufio.NewReader(os.Stdin)
	// Run the repl and catch errors
	err := repl.SimpleRepl(reader)

	switch err.(type) {
	case *repl.QuitErr:
		// this is the quit signal
		fmt.Println(err.Error())
		// exit gracefully
		return
	default:
		// exit code 1 and error message
		log.Fatalf("oim8! FATAL: %s\n", err)
	}
}

// RunServer initialises the server and runs it as the main function.
func RunServer() {
	// supply the server with the stdout.
	writer := bufio.NewWriter(os.Stdout)

	// initialise the server
	server, err := peer.NewServer(writer)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(server.Serve())
}
