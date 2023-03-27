package main

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	ColorOutput    bool      `json:"color_output"`
	Commands       []Command `json:"commands"`
	LogFile        string    `json:"log_file"` /* TODO */
	SurpressOutput bool      `json:"surpress_output"`
}

type Command struct {
	Cwd string `json:"cwd"`
	Cmd string `json:"cmd"`
}

const CONFIG_PATH = ".gleichzeitig/config.json"

var DEFAULT_CONFIG = Config{
	ColorOutput: true,
	Commands: []Command{
		{
			Cwd: ".gleichzeitig",
			Cmd: "ls -la",
		},
		{
			Cwd: ".",
			Cmd: "ls -la",
		},
	},
	LogFile:        "",
	SurpressOutput: false,
}

func getConfig() Config {
	conf := Config{}
	content, err := os.ReadFile(CONFIG_PATH)
	if err != nil {
		log.Fatalln("config not found, create one using 'gleichzeitig init'")
	}
	err = json.Unmarshal(content, &conf)
	if err != nil {
		log.Fatalln("config couldn't be parsed, generate a new one using 'gleichzeitig init'")
	}
	return conf
}
