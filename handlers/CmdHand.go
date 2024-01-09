package handlers

import (
	"github.com/mymmrac/telego"
	"log"
	"mara/utils"
	"strings"
)

var Commands = make(map[string]command)

type command func(u *telego.Update, b *telego.Bot)

func Register(name string, handlerFunc command) {
	Commands[name] = handlerFunc
}

func HandleCommand(u *telego.Update, b *telego.Bot) {
	text := utils.TrimFirstRune(u.Message.Text)
	parts := strings.Fields(text)
	if len(parts) == 0 {
		return
	}
	commandName := parts[0]

	cmd, ok := Commands[commandName]
	if !ok {
		message := telego.SendMessageParams{
			ChatID: telego.ChatID{ID: u.Message.Chat.ID, Username: u.Message.From.Username},
			Text:   "Unknown command!",
		}
		log.Printf("Unknown command: %s", commandName)
		b.SendMessage(&message)
		return
	}
	cmd(u, b)
}
