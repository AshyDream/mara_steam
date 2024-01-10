package handlers

import (
	"fmt"
	"github.com/mymmrac/telego"
	"mara/dbQueries"
	"mara/utils"
)

func DiscountMessageRoad(gameId int, gameInfo map[string]map[string]string, b *telego.Bot) {
	lastId := 0
	for {
		users := dbQueries.GetUserGroup(lastId, 5, gameId)
		if len(users) == 0 {
			lastId = 0
			return
		}
		for _, user := range users {
			tmpUserId := user["User_id"]
			userId := utils.StringToInt(fmt.Sprintf("%v", tmpUserId))
			curr := dbQueries.GetUserCurr(userId)
			tmpId := user["Id"]
			id := utils.StringToInt(fmt.Sprintf("%v", tmpId))
			lastId = id
			SendDiscountMessage(gameId, curr, userId, gameInfo, b)
		}
		break
	}
}

func SendDiscountMessage(gameId int, curr string, userId int, gameInfo map[string]map[string]string, b *telego.Bot) {
	url := utils.UrlBuilder(gameId)
	fmt.Println(gameInfo[curr])
	pricePct := gameInfo[curr]["pct"]
	oldPrice := gameInfo[curr]["priceActl"]
	newPrice := gameInfo[curr]["priceDscnt"]
	gameTitle := gameInfo[curr]["gameTitle"]
	mText := gameTitle + ": \nDiscount: " + pricePct + "\nOld price: " + oldPrice + "\nNew price " + newPrice + "\n" + url
	message := telego.SendMessageParams{
		ChatID: telego.ChatID{ID: int64(userId)},
		Text:   mText,
	}
	_, err := b.SendMessage(&message)
	if err != nil {
		fmt.Println(err)
	}
}
