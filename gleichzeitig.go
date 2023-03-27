package main

import (
	"fmt"
	"os"
	"os/signal"
)

var VERSION = "0.0.1"
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
			for i, c := range COMMANDS {
				fmt.Println()
				commandPrint(i, "terminating...")
				err := c.Process.Kill()
				if err != nil {
					commandPrint(i, "failed to terminate...")
				}
			}
			os.Exit(1)
		}()
		startCommands()
	}
}
