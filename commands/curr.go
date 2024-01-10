package commands

import (
	"fmt"
	"github.com/mymmrac/telego"
	"mara/handlers"
)

var (
	buttons = [][]telego.InlineKeyboardButton{
		{{Text: "\U0001F1FA\U0001F1E6", CallbackData: "UA"}},
		{{Text: "\U0001F1FA\U0001F1F8", CallbackData: "US"}},
		{{Text: "\U0001F1E9\U0001F1EA", CallbackData: "GE"}},
		//{{Text: "\U0001F1F5\U0001F1F1", CallbackData: "PL"}},
	}
)

func init() {
	//newB := []telego.InlineKeyboardButton{
	//	{Text: "dsa", CallbackData: "dsa"},
	//}
	//buttons := append(buttons, newB)
	handlers.Register("curr", func(u *telego.Update, b *telego.Bot) {
		keyboardMarkup := telego.InlineKeyboardMarkup{
			InlineKeyboard: buttons,
		}

		message := telego.SendMessageParams{
			Text:        "Here u go",
			ChatID:      telego.ChatID{ID: u.Message.Chat.ID},
			ReplyMarkup: &keyboardMarkup,
		}
		_, err := b.SendMessage(&message)
		if err != nil {
			fmt.Println(err)
		}

	})
}
