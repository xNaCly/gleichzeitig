package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"sync"
)

const (
	ANSI_RESET   = "\033[0m"
	ANSI_RED     = "\033[31m"
	ANSI_GREEN   = "\033[32m"
	ANSI_YELLOW  = "\033[33m"
	ANSI_BLUE    = "\033[34m"
	ANSI_MAGENTA = "\033[35m"
	ANSI_CYAN    = "\033[36m"
)

var colors = []string{
	ANSI_CYAN,
	ANSI_MAGENTA,
	ANSI_YELLOW,
	ANSI_GREEN,
	ANSI_BLUE,
	ANSI_RED,
}

var wg sync.WaitGroup

// prints not empty text with color to allow the user to separate the output of different commands.
// index is used to select a color, therefore a maximum of 6 commands can be executed in parallel
func commandPrint(index int, text string) {
	if len(text) == 0 {
		return
	}
	i := func() string {
		if index < 9 {
			return fmt.Sprintf("0%d", index)
		}
		return fmt.Sprint(index)
	}()
	fmt.Printf("%s%s | %s%s\n", colors[index], i, text, ANSI_RESET)
}

// startCommand executes the given command in a goroutine and prints the output of it using commandPrint
func startCommand(command string, index int) {
	commandPrint(index, fmt.Sprintf("executing: '%s'", command))
	args := strings.Split(command, " ")
	go func() {
		defer wg.Done()
		out, err := exec.Command(args[0], args[1:]...).Output()
		if err != nil {
			log.Fatalln(err)
		}
		for _, line := range strings.Split(string(out), "\n") {
			commandPrint(index, line)
		}
	}()
}

func main() {
	// TODO: get instructions to run in parallel
	// TODO: <Ctrl-C> should terminate all running scripts
	commands := []string{"ls -la", "whoami", "uptime -p", "who", "groups", "neofetch"}
	if len(commands) > len(colors) {
		log.Fatalln("too many commands, a maximum of 6 commands can be executed in parallel")
	}
	for i, c := range commands {
		wg.Add(1)
		startCommand(c, i)
	}
	wg.Wait()
}
