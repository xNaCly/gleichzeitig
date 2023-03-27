package main

import (
	"os"
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
		startCommands()
	}
}
