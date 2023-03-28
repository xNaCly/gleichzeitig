package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

var COMMANDS = []*exec.Cmd{}

func run(subArgs []string) {
	CONFIG = DEFAULT_CONFIG
	CONFIG.Commands = []Command{}
	if len(subArgs) == 0 {
		logErr("no command given, aborting")
	}
	for _, c := range subArgs {
		CONFIG.Commands = append(CONFIG.Commands, Command{Cmd: c})
	}
	startCommands()
}

// startCommand executes the given command in a goroutine and prints the output of it using commandPrint
func startCommand(command Command, index int) {
	startTime := time.Now()
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
		logErr("error while executing command '" + command.Cmd + "', aborting")
	}

	cmd.Wait()

	commandPrint(index, "done, took "+time.Since(startTime).String())
}

func startCommands() {
	logInfo(fmt.Sprintf("found '%d' commands", len(CONFIG.Commands)))
	startTime := time.Now()
	if len(CONFIG.Commands) == 0 {
		logErr("no commands found, aborting")
	}
	for i, c := range CONFIG.Commands {
		wg.Add(1)
		go startCommand(c, i)
	}
	wg.Wait()
	logInfo("all commands finished, took " + time.Since(startTime).String())
	os.Exit(0)
}

func initGleichzeitig() {
	if _, err := os.Stat(CONFIG_PATH); err == nil {
		logWarn("config file already exists, do you want to overwrite it?")
		fmt.Print("y/n: ")
		var input string
		fmt.Scanln(&input)
		if input == "n" {
			logInfo("aborting")
			os.Exit(1)
		}
		logInfo("got 'y', overwriting config file")
	}
	err := os.Mkdir(".gleichzeitig", os.ModePerm)
	if err != nil {
		logWarn("couldn't create '.gleichzeitig' directory: '" + err.Error() + "'")
	} else {
		logInfo("created '.gleichzeitig' dir")
	}

	file, err := os.Create(".gleichzeitig/config.json")
	if err != nil {
		logErr("couldn't create '.gleichzeitig/config.json' file: '" + err.Error() + "'")
	}
	logInfo("created '.gleichzeitig/config.json' file")

	out, _ := json.MarshalIndent(DEFAULT_CONFIG, "", "\t")
	_, err = file.Write(out)
	if err != nil {
		logErr("couldn't write default config to '.gleichzeitig/config.json' file: '" + err.Error() + "'")
	}
	logInfo("wrote default config to '.gleichzeitig/config.json' file")
}

func printHelp() {
	fmt.Println(`gleichzeitig - run multiple commands in parallel
by: xnacly

Usage:
    gleichzeitig [command]

Run 'gleichzeitig' without any command to execute the commands defined in the config file.

Commands:
    init, i        create a default config file
    run, r         run arguments as commands  
    help, h        print this help message
    version, v     print the version

Examples:
    gleichzeitig init
    gleichzeitig run "echo 'hello world'" "echo 'hello world 2'"
    gleichzeitig help
    gleichzeitig version

website: https://github.com/xnacly/gleichzeitig`)
}

func printVersion() {
	fmt.Printf("gleichzeitig [version=%s] [commit=%s] [commitDate=%s]\n", VERSION, COMMITSHA, COMMITDATE)
}
