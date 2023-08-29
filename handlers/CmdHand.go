package handlers

import (
	"github.com/mymmrac/telego"
	"log"
	"strings"
	"unicode/utf8"
)

var Commands = make(map[string]command)

type command func(u *telego.Update, b *telego.Bot)

func Register(name string, handlerFunc command) {
	Commands[name] = handlerFunc
}

func HandleCommand(u *telego.Update, b *telego.Bot) {
	text := trimFirstRune(u.Message.Text)
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

func trimFirstRune(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}
