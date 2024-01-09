package utils

import (
	"encoding/json"
	"fmt"
	"github.com/mymmrac/telego"
	"os"
)

type Configuration struct {
	Token    string
	DBSource string
}

func cfg() Configuration {
	file, _ := os.Open("config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	cfg := Configuration{}
	err := decoder.Decode(&cfg)
	if err != nil {
		fmt.Println("Error cfg: ", err)
	}
	return cfg
}

var (
	cfgs     = cfg()
	BotToken = cfgs.Token
	Cmds     = []telego.BotCommand{
		{Command: "start", Description: "Launch the bot"},
		{Command: "add", Description: "Subscribe to game"},
		{Command: "curr", Description: "Choose ur current currency"}}
	DbSource = cfgs.DBSource
)
