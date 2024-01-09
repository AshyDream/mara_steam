package commands

import (
	"github.com/mymmrac/telego"
	"mara/handlers"
	"mara/utils"
)

func init() {
	handlers.Register("startTimer", func(u *telego.Update, b *telego.Bot) {
		if u.Message.From.ID == int64(utils.Admin_id) {
			handlers.DayTimer(b)
		}
	})
}
