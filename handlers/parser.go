package handlers

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/mymmrac/telego"
	"io"
	"log"
	"mara/dbQueries"
	"mara/utils"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func Parser(url string, b *telego.Bot) {

	url += "?l=english"

	cookies := []*http.Cookie{
		{Name: "wants_mature_content", Value: "1"},
		{Name: "lastagecheckage", Value: "1-0-1983"},
		{Name: "birthtime", Value: "407541601"},
	}

	doc := fetch(url, cookies)

	gameTitle := doc.Find(".apphub_AppName").First().Text()
	fmt.Println(gameTitle)

	purchaseBox := doc.Find(".game_area_purchase").Find(".game_area_purchase_game")

	var discCheck string
	if doc.Find(".demo_above_purchase").Text() != "" {
		discCheck = purchaseBox.First().Next().Find(".game_purchase_discount").Text()
	} else {
		discCheck = purchaseBox.First().Find(".game_purchase_discount").Text()
	}

	id := utils.IdTrimer(url)

	if dbQueries.OnSale(id) {
		if discCheck == "" {
			dbQueries.OnSaleChanger(1, id)
			return
		}
		return
	} else {
		if discCheck == "" {
			return
		} else {
			dbQueries.OnSaleChanger(0, id)
		}
	}

	var gameInfo = map[string]map[string]string{
		"UA": {},
		"US": {},
		"GE": {},
		//"PL": {},
	}

	for key := range gameInfo {
		nUrl := url + "&cc=" + key
		key := key
		wg.Add(1)
		go func() {
			doc = fetch(nUrl, cookies)
			info := gatherInfo(doc, gameTitle)
			gameInfo[key]["pct"] = info[0]
			gameInfo[key]["priceActl"] = info[1]
			gameInfo[key]["priceDscnt"] = info[2]
			gameInfo[key]["gameTitle"] = info[3]
			wg.Done()
		}()
	}
	wg.Wait()
	DiscountMessageRoad(id, gameInfo, b)
}

func gatherInfo(doc *goquery.Document, gameTitle string) []string {
	purchaseBox := doc.Find(".game_area_purchase").Find(".game_area_purchase_game")

	var gameBox goquery.Selection
	if doc.Find(".demo_above_purchase").Text() != "" {
		gameBox = *purchaseBox.First().Next().Find(".game_purchase_discount")
	} else {
		gameBox = *purchaseBox.First().Find(".game_purchase_discount")
	}

	gamePrices := gameBox.Find(".discount_prices").Children()
	gameDscntPct := gameBox.Find(".discount_pct").Text()
	gameActlPrice := gamePrices.First().Text()
	gameDscntPrice := gamePrices.Next().Text()

	info := []string{gameDscntPct, gameActlPrice, gameDscntPrice, gameTitle}
	return info
}

func fetch(url string, cookies []*http.Cookie) *goquery.Document {

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Error creating HTTP request: %v", err)
	}

	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}

	response, err := client.Do(req)
	if err != nil {
		log.Printf("Error HTTP fetch: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(response.Body)

	if response.StatusCode != 200 {
		log.Printf("Wrong Status CODE: %d", response.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Printf("Error parsing page: %v", err)
	}

	return doc
}
