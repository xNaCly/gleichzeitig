package main

import (
	"fmt"
	"os"
	"os/signal"
)

var VERSION = ""
var COMMITSHA = ""
var COMMITDATE = ""

var CONFIG Config = Config{}

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "init", "i":
			initGleichzeitig()
		case "help", "h":
			printHelp()
		case "version", "v":
			printVersion()
		}
	} else {
		CONFIG = getConfig()
		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt)
		go func() {
			<-c
			fmt.Println()
			for i, c := range COMMANDS {
				err := c.Process.Kill()
				if err != nil {
					commandPrint(i, "failed to terminate...")
				} else {
					commandPrint(i, "terminated...")
				}
			}
			os.Exit(1)
		}()
		startCommands()
	}
}
