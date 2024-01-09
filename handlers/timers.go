package handlers

import (
	"fmt"
	"github.com/mymmrac/telego"
	"mara/dbQueries"
	"mara/utils"
	"time"
)

func DayTimer(b *telego.Bot) {
	dayTime := time.Now()
	for {
		ParserTimer(b)
		dayTime = dayTime.Add(time.Hour * 24)
		until := time.Until(dayTime)
		time.Sleep(until)
	}

}

func ParserTimer(b *telego.Bot) {
	lastId := 0
	for {
		games := dbQueries.GetGamesGroup(lastId, 2)
		if len(games) == 0 {
			lastId = 0
			return
		}
		for _, game := range games {
			game := game
			tmpId := game["Id"]
			id := utils.StringToInt(fmt.Sprintf("%v", tmpId))
			lastId = id
			url := utils.UrlBuilder(id)
			Parser(url, b)
		}
	}
}
