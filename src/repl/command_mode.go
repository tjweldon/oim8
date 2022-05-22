package repl

import (
    "fmt"
    "strings"
)

// commandMode is where to do things like quit.
func commandMode(text string) error {
    commands := map[string]ApplicationCommand{
        "q": quit,
    }
    
    args := strings.Split(text, " ")
    if len(args) == 0 {
        fmt.Printf("%+v", args)
        return User(fmt.Errorf("No command supplied"))
    }
 
    // attempt to find the command
    command, ok := commands[args[0]]
    if !ok {
        // otherwise error (mode specific implementation)
        return User(fmt.Errorf("Command does not exist"))
    }

    // run the command
    return command(args...)
}

// quit is the ApplicationCommand for quitting
func quit(args ...string) error {
    return Quit()
}
