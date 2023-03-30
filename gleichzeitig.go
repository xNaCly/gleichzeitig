package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
)

var VERSION = ""
var COMMITSHA = ""
var COMMITDATE = ""

var CONFIG Config = Config{}

func main() {
	log.SetFlags(log.Ltime)
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "run", "r":
			run(os.Args[2:])
		case "init", "i":
			initGleichzeitig()
		case "help", "h":
			printHelp()
		case "version", "v":
			printVersion()
		}
	} else {
		CONFIG = getConfig()
		pwd, _ := os.Getwd()
		logInfo("starting gleichzeitig with config located at '" + pwd + "/" + CONFIG_PATH + "'")
		if CONFIG.LogFile != "" {
			logInfo("found config value for 'config.log_file'")
			logInfo("Logging to file: '" + CONFIG.LogFile + "'")
			f, err := os.OpenFile(CONFIG.LogFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
			if err != nil {
				logErr("failed to open logging file: '" + err.Error() + "'")
			}
			defer f.Close()
			mv := io.MultiWriter(os.Stdout, f)
			log.SetOutput(mv)
		}
		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt)
		go func() {
			<-c
			fmt.Println()
			for i, c := range COMMANDS {
				err := c.Process.Kill()
				if err != nil {
					commandPrint(i, "command already stopped.")
				} else {
					commandPrint(i, "terminated.")
				}
			}
			os.Exit(1)
		}()
		startCommands()
	}
}
