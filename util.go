package main

import (
	"fmt"
	"log"
	"sync"
)

const (
	ANSI_RESET   = "\033[0m"
	ANSI_RED     = "\033[91m"
	ANSI_GREEN   = "\033[92m"
	ANSI_YELLOW  = "\033[93m"
	ANSI_BLUE    = "\033[94m"
	ANSI_MAGENTA = "\033[95m"
	ANSI_CYAN    = "\033[96m"
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
	if !CONFIG.SurpressOutput {
		if CONFIG.ColorOutput {
			fmt.Printf("%s%s | %s%s\n", colors[index%len(colors)], i, text, ANSI_RESET)
		} else {
			fmt.Printf("%s | %s\n", i, text)
		}
	}
}

func logInfo(text string) {
	log.Printf("info: %s\n", text)
}

func logWarn(text string) {
	log.Printf("warn: %s\n", text)
}

func logErr(text string) {
	log.Fatalf("err: %s\n", text)
}
