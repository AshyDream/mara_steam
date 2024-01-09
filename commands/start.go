package commands

import (
	"github.com/mymmrac/telego"
	"mara/handlers"
)

func init() {
	handlers.Register("start", func(u *telego.Update, b *telego.Bot) {
		message := telego.SendMessageParams{
			ChatID: telego.ChatID{ID: u.Message.Chat.ID, Username: u.Message.From.Username},
			Text:   "Привіт! \n Щоб підписатися на розсилку про знижки на гру, спочатку вибери країну із списку доступних, ввівши команду curr, або вибравши відповідну команду з меню бота. Після чого можеш додавати підписки за допомогою команди `add steam-url`. \n Також, можеш вивести назви ігор на які ти уже підписанний (команда list), та видалити підписку (команда delete)",
		}
		_, err := b.SendMessage(&message)
		if err != nil {
			panic(err)
		}
		return
	})
}
