package main

import (
	"encoding/json"
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
			Cmd: "echo 'Hello World from command 1!'",
		},
		{
			Cmd: "echo 'Hello World from command 2!'",
		},
	},
	LogFile:        "gleich.log",
	SurpressOutput: false,
}

func getConfig() Config {
	conf := Config{}
	content, err := os.ReadFile(CONFIG_PATH)
	if err != nil {
		logErr("config not found, create one using 'gleichzeitig init'")
	}
	err = json.Unmarshal(content, &conf)
	if err != nil {
		logErr("config is not valid json")
	}
	return conf
}
