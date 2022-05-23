package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"tjweldon/oim8/src/repl"
)

func main() {
	DoRepl()
}

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
