package commands

import (
	"fmt"
	"github.com/mymmrac/telego"
	"mara/handlers"
	"reflect"
	"strings"
)

func RegisterAllCommands() {
	cmdType := reflect.TypeOf(handlers.Commands)
	for i := 0; i < cmdType.NumMethod(); i++ {
		method := cmdType.Method(i)
		cmdName := strings.ToLower(method.Name)
		handlers.Register(cmdName, func(u *telego.Update, b *telego.Bot) {
			cmd := handlers.Commands[cmdName]
			args := []reflect.Value{reflect.ValueOf(cmd), reflect.ValueOf(u), reflect.ValueOf(b)}
			method.Func.Call(args)
		})
	}
	var i int
	for k, _ := range handlers.Commands {
		i++
		fmt.Printf("[\033[1m%v - Command: \033[3m\033[33m%s\033[0m]\n", i, k)
	}
}
