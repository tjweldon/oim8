package main

import (
	"bufio"
	"fmt"
	"os"
	"tjweldon/oim8/src/repl"
)

func main() {
	// supply stdin to REPL
	reader := bufio.NewReader(os.Stdin)

	// Run the repl and catch errors
	//log.Fatal(SimpleRepl(reader))
	err := repl.SimpleRepl(reader)

	switch err.(type) {
	case repl.QuitErr:
		// this is the quit signal
		fmt.Println(err.Error())
		// exit gracefully
		return
	default:
        // exit code 1 and error message
        fmt.Printf("oim8! FATAL: %s\n",err)
	}
}
