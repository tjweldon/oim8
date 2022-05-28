package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
	"tjweldon/oim8/src/peer"
	"tjweldon/oim8/src/repl"
)

func main() {
	done := make(chan struct{})
	defer func(d chan<- struct{}) {
		d <- struct{}{}
		close(d)
	}(done)
	msgs := make(chan []byte)
	defer close(msgs)

	// start the server
	go RunServer(done)
	time.Sleep(time.Second)

	// Start a client
	client := peer.NewClient("127.0.0.1", "42069", msgs)
	go client.Run()

	msgs <- []byte("Hello?")
	msgs <- []byte("u there?")
	msgs <- []byte("ok speak later")

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
func RunServer(done <-chan struct{}) {
	// supply the server with the stdout.
	writer := bufio.NewWriter(os.Stdout)

	// initialise the server
	server, err := peer.NewServer(writer)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(server.Serve(done))
}
