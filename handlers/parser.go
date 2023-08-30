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
	//// Знаходимо елементи за їх CSS класами або іншими селекторами
	gameTitle := doc.Find(".apphub_AppName").First().Text()
	fmt.Println(gameTitle)
	buyBlock := doc.Find(".game_purchase_action")
	discBlock := buyBlock.Find(".discount_block.game_purchase_discount").First()
	discBlockText := buyBlock.Text()
	if discBlockText != "" {
		pricePct := discBlock.Find(".discount_pct").Text()
		oldPrice := discBlock.Find(".discount_original_price").Text()
		newPrice := discBlock.Find(".discount_final_price").Text()
		mText := gameTitle + ": \nDiscount: " + pricePct + "\nOld price: " + oldPrice + "\nNew price " + newPrice + "\n" + url
		message := telego.SendMessageParams{
			ChatID: telego.ChatID{ID: u.Message.Chat.ID, Username: u.Message.From.Username},
			Text:   mText,
		}
		b.SendMessage(&message)
		return
	}
	fmt.Printf("No discount for %s right now", gameTitle)
}
