package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
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

func commandPrint(index int, text string) {
	i := func() string {
		if index < 9 {
			return fmt.Sprintf("0%d", index)
		}
		return fmt.Sprint(index)
	}()
	fmt.Printf("%s%s | %s%s\n", colors[index], i, text, ANSI_RESET)
}

func startCommand(command string, index int) {
	args := strings.Split(command, " ")
	out, err := exec.Command(args[0], args[1:]...).Output()
	if err != nil {
		log.Fatalln(err)
	}
	for _, line := range strings.Split(string(out), "\n") {
		commandPrint(index, line)
	}
}

func main() {
	// TODO: get instructions to run in parallel
	// TODO: execute programs in parallel
	// TODO: wait for execution
	// TODO: <Ctrl-C> should terminate all running scripts
	commands := []string{"ls -la", "whoami", "uptime -p", "who", "groups", "echo $SHELL"}
	for i, c := range commands {
		commandPrint(i, fmt.Sprintf("executing: '%s'", c))
		startCommand(c, i)
	}
}
