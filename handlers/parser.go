package handlers

import (
	"fmt"
	"github.com/mymmrac/telego"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func Parser(url string, u *telego.Update, b *telego.Bot) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error get response from %s: %s", url, err)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		fmt.Printf("Error status code %v", response.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	gameTitle := doc.Find(".apphub_AppName").First().Text()
	fmt.Println(gameTitle)

	firstPurchaseBlock := doc.Find(".game_area_purchase_game").First()
	//titleCheck := checkPurchaseBlock(firstPurchaseBlock)
	//fmt.Println(titleCheck)

	buyBlock := firstPurchaseBlock.Find(".game_purchase_action").First()
	discBlock := buyBlock.Find(".discount_block.game_purchase_discount").First()
	fmt.Println(discBlock.Text())
	//discCheck := discBlock.Find(".game_purchase_action_bg").First()
	//if len(discCheck.Nodes) > 1 {
	//	return
	//}
	discBlockText := discBlock.Text()
	if discBlockText != "" {
		pricePct := discBlock.Find(".discount_pct").Text()
		oldPrice := discBlock.Find(".discount_original_price").Text()
		newPrice := discBlock.Find(".discount_final_price").Text()
		mText := gameTitle + ": \nDiscount: " + pricePct + "\nOld price: " + oldPrice + "\nNew price " + newPrice + "\n" + url
		message := telego.SendMessageParams{
			ChatID: telego.ChatID{ID: u.Message.Chat.ID, Username: u.Message.From.Username},
			Text:   mText,
		}
		_, err := b.SendMessage(&message)
		if err != nil {
			fmt.Printf("\033[31mparser.go l52 can't send a message")
		}
		return
	}
	fmt.Printf("\033[31mNo discount for %s right now\033[0m", gameTitle)
}

//func checkPurchaseBlock(block *goquery.Selection) bool {
//
//}
