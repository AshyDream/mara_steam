package commands

import (
	"fmt"
	"github.com/mymmrac/telego"
	"mara/dbQueries"
	"mara/handlers"
	"strconv"
)

func init() {
	handlers.Register("delete", func(u *telego.Update, b *telego.Bot) {
		idGames := dbQueries.List(int(u.Message.Chat.ID))
		fmt.Println(idGames)
		var gameButtons [][]telego.InlineKeyboardButton
		for _, v := range idGames {
			game := dbQueries.GameInfo(v)
			fmt.Println(game)
			title := game["Title"]
			newB := []telego.InlineKeyboardButton{
				{Text: fmt.Sprintf("%v", title), CallbackData: strconv.Itoa(v)},
			}
			gameButtons = append(gameButtons, newB)
		}
		keyboardMarkup := telego.InlineKeyboardMarkup{InlineKeyboard: gameButtons}
		message := telego.SendMessageParams{
			Text:        "Pick the game u want to unsubscribe",
			ChatID:      telego.ChatID{ID: u.Message.Chat.ID},
			ReplyMarkup: &keyboardMarkup,
		}
		_, err := b.SendMessage(&message)
		if err != nil {
			fmt.Println(err)
		}
	})
}
