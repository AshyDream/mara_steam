package commands

import (
	"fmt"
	"github.com/mymmrac/telego"
	"mara/handlers"
)

func init() {
	handlers.Register("pong", func(u *telego.Update, b *telego.Bot) {
		message := telego.SendMessageParams{
			ChatID: telego.ChatID{ID: u.Message.Chat.ID, Username: u.Message.From.Username},
			Text:   "Ping!",
		}

		_, err := b.SendMessage(&message)
		if err != nil {
			fmt.Println("Pong problem ")
			return
		}
	})
}
