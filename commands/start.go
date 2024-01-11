package commands

import (
	"github.com/mymmrac/telego"
	"mara/handlers"
)

func init() {
	handlers.Register("start", func(u *telego.Update, b *telego.Bot) {
		message := telego.SendMessageParams{
			ChatID: telego.ChatID{ID: u.Message.Chat.ID, Username: u.Message.From.Username},
			Text: "Hello!" +
				"\n  To subscribe to the game discount newsletter, first, select a country from the list of available countries by typing the \"/curr\" command or by selecting the appropriate command from the bot menu." +
				"\n After that, you can add subscriptions using the `/add steam-url' command." +
				"\n  Also, you can display the names of games to which you are already subscribed (\"/list\" command) and delete the subscription (\"/delete\" command)" +
				"\n If you find a bug or bot is not working pls contact bot maintainer: @ashe_dream",
		}
		_, err := b.SendMessage(&message)
		if err != nil {
			panic(err)
		}
		return
	})
}
