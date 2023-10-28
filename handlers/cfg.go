package handlers

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	Token    string
	DBSource string
}

func Cfg() Configuration {
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
