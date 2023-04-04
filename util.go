package main

import (
	"fmt"
	"log"
	"sync"
)

const (
	ANSI_RESET         = "\x1b[0m"
	ANSI_RED           = "\x1b[31m"
	ANSI_GREEN         = "\x1b[32m"
	ANSI_YELLOW        = "\x1b[33m"
	ANSI_BLUE          = "\x1b[34m"
	ANSI_MAGENTA       = "\x1b[35m"
	ANSI_CYAN          = "\x1b[36m"
	ANSI_RED_BRIGHT    = "\x1b[91m"
	ANSI_YELLOW_BRIGHT = "\x1b[93m"
	ANSI_CYAN_BRIGHT   = "\x1b[96m"
)

var colors = []string{
	ANSI_BLUE,
	ANSI_MAGENTA,
	ANSI_GREEN,
	ANSI_CYAN,
	ANSI_YELLOW,
	ANSI_RED,
}

var wg sync.WaitGroup

// prints not empty text with color to allow the user to separate the output of different commands.
// index is used to select a color, therefore a maximum of 6 commands can be executed in parallel
func commandPrint(index int, text string) {
	i := func() string {
		if index < 9 {
			return fmt.Sprintf("0%d", index)
		}
		return fmt.Sprint(index)
	}()

	if CONFIG.SurpressOutput {
		return
	}

	if CONFIG.ColorOutput {
		if CONFIG.OnlyColorPrefix {
			log.Printf("%s%s |%s %s\n", colors[index%len(colors)], i, ANSI_RESET, text)
			return
		}
		log.Printf("%s%s | %s%s\n", colors[index%len(colors)], i, text, ANSI_RESET)
	} else {
		log.Printf("%s | %s\n", i, text)
	}
}

func logInfo(text string) {
	log.Printf("%sinfo%s: %s\n", ANSI_CYAN_BRIGHT, ANSI_RESET, text)
}

func logWarn(text string) {
	log.Printf("%swarn%s: %s\n", ANSI_YELLOW_BRIGHT, ANSI_RESET, text)
}

func logErr(text string) {
	log.Fatalf("%serr%s: %s\n", ANSI_RED_BRIGHT, ANSI_RESET, text)
}
