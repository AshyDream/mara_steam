package commands

import (
	"github.com/mymmrac/telego"
	"mara/dbQueries"
	"mara/handlers"
)

func init() {
	handlers.Register("curr", func(u *telego.Update, b *telego.Bot) {
		chatId := u.Message.Chat.ID
		if !dbQueries.IsUser(int(chatId)) {
			dbQueries.AddNewUser(u, "UA")
			return
		}

	})
}
