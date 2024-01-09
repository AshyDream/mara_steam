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
	Admin_Id string
}

func cfg() Configuration {
	file, _ := os.Open("config.json")
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)
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
		{Command: "curr", Description: "Choose ur current currency"},
		{Command: "delete", Description: "Unsubscribe from game"},
		{Command: "list", Description: "List subsribed games"}}
	DbSource = cfgs.DBSource
	Admin_id = cfgs.Admin_Id
)
