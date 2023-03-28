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
func startCommand(command Command, index int) {
	args := strings.Split(command.Cmd, " ")
	defer wg.Done()

	cmd := exec.Command(args[0], args[1:]...)
	COMMANDS = append(COMMANDS, cmd)

	if command.Cwd != "" {
		commandPrint(index, fmt.Sprintf("changing working directory to: '%s'", command.Cwd))
		cmd.Dir = command.Cwd
	}

	commandPrint(index, fmt.Sprintf("executing: '%s'", command.Cmd))

	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()

	go func() {
		scanner := bufio.NewScanner(stdout)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			commandPrint(index, scanner.Text())
		}
	}()
	go func() {
		scanner := bufio.NewScanner(stderr)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			commandPrint(index, scanner.Text())
		}
	}()

	err := cmd.Start()

	if err != nil {
		commandPrint(index, fmt.Sprintf("error: %s", err))
		logErr("error while executing command, aborting")
	}

	cmd.Wait()

	commandPrint(index, "finished...")
}

func startCommands() {
	logInfo(fmt.Sprintf("found '%d' commands", len(CONFIG.Commands)))
	if len(CONFIG.Commands) == 0 {
		logErr("no commands found, aborting")
	}
	for i, c := range CONFIG.Commands {
		wg.Add(1)
		go startCommand(c, i)
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
