package commands

import (
	"github.com/mymmrac/telego"
	"mara/handlers"
	"mara/utils"
)

func init() {
	handlers.Register("startTimer", func(u *telego.Update, b *telego.Bot) {
		tmpAdmin_id := utils.Admin_id
		admin_id := utils.StringToInt(tmpAdmin_id)
		if u.Message.From.ID == int64(admin_id) {
			handlers.DayTimer(b)
		}
	})
}
