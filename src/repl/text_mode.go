package repl

import (
	"fmt"
	"strings"
)

// messageMode is sending messages
func messageMode(text string) error {
	commands := map[string]ApplicationCommand{
		"@w": whisper,
	}

	args := strings.Split(text, " ")
	if len(args) == 0 {
		return User(fmt.Errorf("Type Something!"))
	}

	command, ok := commands[args[0]]
	if !ok {
		command = send 
	}

	return command(args...)
}

// send is TODO
func send(args ...string) error {
    return nil
}

// whisper should send a message to a single peer (TODO)
func whisper(args ...string) error {
	if len(args) < 2 {
		User(fmt.Errorf("Not enough args supplied"))
	}
	fmt.Printf(" [whispered]@%s", args[1])
	return nil
}
