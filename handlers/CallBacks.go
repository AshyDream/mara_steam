package handlers

import (
	"fmt"
	"github.com/mymmrac/telego"
	"mara/dbQueries"
	"mara/utils"
)

var (
	currCBs = []string{"UA", "US", "GE"}
)

func CallbackRoad(u *telego.Update, b *telego.Bot) {
	cb := u.CallbackQuery
	if isCurrChange(cb) {
		currCallBackRoad(cb, u, b)
	} else {
		gameId := utils.StringToInt(cb.Data)
		if dbQueries.DeleteUserGame(*u, gameId) {
			msg := "You successful unsubscribed!"
			message := telego.SendMessageParams{
				Text:   msg,
				ChatID: telego.ChatID{ID: u.CallbackQuery.Message.Chat.ID},
			}
			_, err := b.SendMessage(&message)
			if err != nil {
				fmt.Println(err)
				return
			}
			msgPrms := telego.DeleteMessageParams{
				ChatID: telego.ChatID{
					ID: u.CallbackQuery.Message.Chat.ID,
				},
				MessageID: u.CallbackQuery.Message.MessageID,
			}
			err = b.DeleteMessage(&msgPrms)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

func isCurrChange(cb *telego.CallbackQuery) bool {
	for _, data := range currCBs {
		if data == cb.Data {
			return true
		}
	}
	return false
}

func currCallBackRoad(cb *telego.CallbackQuery, u *telego.Update, b *telego.Bot) {
	chatId := u.CallbackQuery.Message.Chat.ID
	if !dbQueries.IsUser(int(chatId)) {
		dbQueries.AddNewUser(*u, cb.Data)
	} else {
		dbQueries.UpdateCurr(*u, cb.Data)
	}
	msg := "Your currency successful changed to: " + string(cb.Data)
	message := telego.SendMessageParams{
		Text:   msg,
		ChatID: telego.ChatID{ID: u.CallbackQuery.Message.Chat.ID},
	}
	_, err := b.SendMessage(&message)
	if err != nil {
		fmt.Println(err)
		return
	}
	msgPrms := telego.DeleteMessageParams{
		ChatID: telego.ChatID{
			ID: u.CallbackQuery.Message.Chat.ID,
		},
		MessageID: u.CallbackQuery.Message.MessageID,
	}
	err = b.DeleteMessage(&msgPrms)
	if err != nil {
		fmt.Println(err)
		return
	}
}
