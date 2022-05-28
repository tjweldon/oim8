package repl

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

// debug mode flag
var debug = os.Getenv("DEBUG") == "true"

// QuitErr is the signal to exit
type QuitErr struct{}

// Quit returns a QuitErr with the exit msg
func Quit() *QuitErr {
	return &QuitErr{}
}

func (q *QuitErr) Error() string {
	return "Quitting!"
}

// UserErr is when the user gets a key wrong or smth
type UserErr struct {
	err error
}

// Error proxies to the internal err
func (ue *UserErr) Error() string {
	return ue.err.Error()
}

// User wraps an err in the appropriate type
func User(err error) *UserErr {
	return &UserErr{err}
}

type Handler func(text string) error

// Mode is the enum of client modes, i.e. Message or Command
type Mode int

const (
	Message Mode = iota
	Command
)

// String returns the human-readable mode value
func (m Mode) String() string {
	return []string{
		"Message",
		"Command",
	}[m]
}

// Handle is the mapping of mode value to Handle function call
func (m Mode) Handle(text string) error {
	return []Handler{
		messageMode,
		commandMode,
	}[m](text)
}

// GetMode is the mapping of prefix
func GetMode(b byte) (mode Mode) {
	mode, _ = map[byte]Mode{
		byte('>'): Command,
	}[b]

	return mode
}

// printHeader is the first message a user sees.
func printHeader() {
	fmt.Println("Simple Repl")
	fmt.Println("-----------")
}

// promptUser is the 'echo $PROMPT' equivalent
func promptUser() {
	yourName := "<user>@<host>"
	fmt.Printf("%s : ", yourName)
}

// SimpleRepl encapsulates the Read-Eval-Print loop
func SimpleRepl(reader *bufio.Reader) error {
	// First prompt
	printHeader()
	var (
		mode  Mode   // the mode to parse the text
		text  string // the text to parse
		input []byte // as bytes
		err   error  // the error if there is one
	)
	for {
		// PRINT: the promptUser message
		promptUser()

		// READ: user input until newline
		input, err = reader.ReadBytes('\n')
		if err != nil {
			return err
		}

		// EVAL: the mode
		mode = takeMode(&input)
		text = string(input)
		if debug {
			_, _ = os.Stderr.WriteString(
				fmt.Sprintf("oim8! DEBUG: input: % x\n", input),
			)
		}

		// EVAL: the input
		// & PRINT: to stdout/stderr
		err = mode.Handle(text)

		// If there's an error, handle it
		if err != nil {
			// handle any UserErr errors without Fatal
			switch err.(type) {
			case *UserErr:
				// ...just user message and drop the error
				fmt.Printf("oim8!: %s\n", err)
				err = nil
			default:
				return err
			}
		}
	}
}

// takeMode takes the first byte to see if the user is entering commandMode,
// trims mode prefix if so.
func takeMode(input *[]byte) (mode Mode) {
	// default and no command
	if len(*input) == 0 {
		return
	}

	// trim trailing newlines
	trimmed := bytes.TrimRight(*input, "\n\r")

	// Modes are controlled with the first byte of input
	modeByte := (*input)[0]

	// modeMapping encodes the bindings for choosing a mode.
	modeMapping := map[byte]Mode{
		byte('>'): Command,
	}
	var ok bool
	mode, ok = modeMapping[modeByte]
	if ok {
		// trim the incoming bytes
		trimmed = trimmed[1:]
	}

	*input = trimmed

	return mode
}

// ApplicationCommand is a subcommand for the mode context
type ApplicationCommand func(args ...string) error
