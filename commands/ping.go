package commands

import (
	"fmt"
	"github.com/mymmrac/telego"
	"mara/handlers"
)

func init() {
	handlers.Register("ping", func(u *telego.Update, b *telego.Bot) {
		message := telego.SendMessageParams{
			ChatID: telego.ChatID{ID: u.Message.Chat.ID, Username: u.Message.From.Username},
			Text:   "Pong!",
		}

		_, err := b.SendMessage(&message)
		if err != nil {
			fmt.Println("Ping problem ")
			return
		}
	})
}
