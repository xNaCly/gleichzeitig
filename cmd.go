package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var COMMANDS = []*exec.Cmd{}

// startCommand executes the given command in a goroutine and prints the output of it using commandPrint
// TODO: this currently only works with commands that stop after they are done
func startCommand(command Command, index int) {
	commandPrint(index, fmt.Sprintf("executing: '%s'", command.Cmd))
	args := strings.Split(command.Cmd, " ")
	defer wg.Done()
	cmd := exec.Command(args[0], args[1:]...)
	COMMANDS = append(COMMANDS, cmd)

	if command.Cwd != "" {
		commandPrint(index, fmt.Sprintf("changing working directory to: '%s'", command.Cwd))
		cmd.Dir = command.Cwd
	}

	stdout, _ := cmd.StdoutPipe()
	err := cmd.Start()

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		commandPrint(index, scanner.Text())
	}

	if err != nil {
		commandPrint(index, fmt.Sprintf("error: %s", err))
		logErr("error while executing command, aborting")
	}

	cmd.Wait()
}

func startCommands() {
	logInfo(fmt.Sprintf("found '%d' commands", len(CONFIG.Commands)))
	if len(CONFIG.Commands) == 0 {
		logErr("no commands found, aborting")
	}
	for i, c := range CONFIG.Commands {
		wg.Add(1)
		startCommand(c, i)
	}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Press 'q' to quit: \n")
	for scanner.Scan() {
		input := scanner.Text()
		if input == "q" {
			for i, c := range COMMANDS {
				commandPrint(i, "killing command")
				c.Process.Kill()
			}
		}
	}
	wg.Wait()
	os.Exit(0)
}

func initGleichzeitig() {
	err := os.Mkdir(".gleichzeitig", os.ModePerm)
	if err != nil {
		logWarn("couldn't create '.gleichzeitig' directory")
	} else {
		logInfo("created '.gleichzeitig' dir")
	}

	file, err := os.Create(".gleichzeitig/config.json")
	if err != nil {
		logErr("couldn't create '.gleichzeitig/config.json' file")
	}
	logInfo("created '.gleichzeitig/config.json' file")

	out, _ := json.MarshalIndent(DEFAULT_CONFIG, "", "\t")
	_, err = file.Write(out)
	if err != nil {
		logErr("couldn't write default config to '.gleichzeitig/config.json' file")
	}
	logInfo("wrote default config to '.gleichzeitig/config.json' file")
}

func printHelp() {
	fmt.Println(`gleichzeitig - run multiple commands in parallel
by: xnacly

Usage:
    gleichzeitig [command]


Commands:
    init, i        create a default config file
    help, h        print this help message
    version, v     print the version

Run 'gleichzeitig' without any command to execute the commands defined in the config file.

website: https://github.com/xnacly/gleichzeitig`)
}

func printVersion() {
	fmt.Printf("gleichzeitig v%s\n", VERSION)
}
