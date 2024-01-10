package commands

import (
	"fmt"
	"github.com/mymmrac/telego"
	"mara/dbQueries"
	"mara/handlers"
)

func init() {
	handlers.Register("list", func(u *telego.Update, b *telego.Bot) {
		idGames := dbQueries.List(int(u.Message.Chat.ID))
		fmt.Println(idGames)
		mtext := "Here we go, that ur game list:\n"
		for _, v := range idGames {
			game := dbQueries.GameInfo(v)
			title := game["Title"]
			mtext = mtext + " --- " + fmt.Sprintf("%v", title) + "\n"
		}
		message := telego.SendMessageParams{
			ChatID: telego.ChatID{ID: u.Message.Chat.ID, Username: u.Message.From.Username},
			Text:   mtext,
		}
		_, err := b.SendMessage(&message)
		if err != nil {
			fmt.Println(err)
		}
	})
}
