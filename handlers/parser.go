package handlers

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/mymmrac/telego"
	"log"
	"net/http"
)

func Parser(url string, u *telego.Update, b *telego.Bot) {

	url += "?l=english"

	doc := fetch(url)

	gameTitle := doc.Find(".apphub_AppName").First().Text()
	fmt.Println(gameTitle)

	purchaseBox := doc.Find(".game_area_purchase").Find(".game_area_purchase_game")

	discCheck := purchaseBox.First().Find(".game_purchase_discount").Text()

	if discCheck == "" {
		return
	}

	var gameInfo = map[string]map[string]string{
		"UA": {},
		"US": {},
		"GE": {},
		"UK": {},
		"JP": {},
		"CA": {},
		"AU": {},
		"NZ": {},
		"NO": {},
		"CH": {},
		"TW": {},
		"AR": {},
		"BR": {},
		"ID": {},
		"KR": {},
		"MX": {},
		"PL": {},
		"EU": {},
	}

	for key := range gameInfo {
		nUrl := url + "&cc=" + key
		fmt.Println(key)
		doc := fetch(nUrl)
		purchaseBox := doc.Find(".game_area_purchase").Find(".game_area_purchase_game")
		info := gatherInfo(purchaseBox)
		gameInfo[key]["pct"] = info[0]
		gameInfo[key]["priceActl"] = info[1]
		gameInfo[key]["priceDscnt"] = info[2]
	}

	fmt.Println(gameInfo)

}

func gatherInfo(doc *goquery.Selection) []string {
	gameBox := doc.First()

	//gameName := gameBox.Find("h1").Text()
	//fmt.Println(gameName)

	gamePrices := gameBox.Find(".discount_prices").Children()
	gameDscntPct := gameBox.Find(".discount_pct").Text()
	gameActlPrice := gamePrices.First().Text()
	gameDscntPrice := gamePrices.Next().Text()

	//for i, box := range purchaseBox {
	//	fmt.Printf("\033[3%vm %v", i+1, box.Attr)
	//}

	info := []string{gameDscntPct, gameActlPrice, gameDscntPrice}
	return info
}

func fetch(url string) *goquery.Document {

	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error get response from %s: %s", url, err)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		fmt.Printf("Error status code: %v", response.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return doc
}
