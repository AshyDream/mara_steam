package commands

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/mymmrac/telego"
	"io"
	"log"
	"mara/dbQueries"
	"mara/handlers"
	"mara/utils"
	"net/http"
	"strings"
)

func init() {
	handlers.Register("add", func(u *telego.Update, b *telego.Bot) {
		text := u.Message.Text
		text = utils.TrimFirstRune(text)
		parts := strings.Fields(text)
		//cc := "us"

		if len(parts) < 2 {
			message := telego.SendMessageParams{
				ChatID: telego.ChatID{ID: u.Message.Chat.ID, Username: u.Message.From.Username},
				Text:   "Pls write /add steam-url",
			}
			_, err := b.SendMessage(&message)
			if err != nil {
				fmt.Println(err)
			}
			return
		}
		url := parts[1]
		cookies := []*http.Cookie{
			&http.Cookie{Name: "wants_mature_content", Value: "1"},
			&http.Cookie{Name: "lastagecheckage", Value: "1-0-1983"},
			&http.Cookie{Name: "birthtime", Value: "407541601"},
		}
		urlExmpl := "https://store.steampowered.com/app/"

		if len(parts) != 2 || !strings.Contains(url, urlExmpl) {
			message := telego.SendMessageParams{
				ChatID: telego.ChatID{ID: u.Message.Chat.ID, Username: u.Message.From.Username},
				Text:   "Wrong arguments!",
			}
			b.SendMessage(&message)
			return
		}

		is, title, isFree, isOnSale, isNotYet := isURLValid(url, cookies)

		if !is {
			message := telego.SendMessageParams{
				ChatID: telego.ChatID{ID: u.Message.Chat.ID, Username: u.Message.From.Username},
				Text:   "Invalid Steam game URL!",
			}
			_, err := b.SendMessage(&message)
			if err != nil {
				fmt.Println(err)
			}
			return
		}
		if isNotYet {
			message := telego.SendMessageParams{
				ChatID: telego.ChatID{ID: u.Message.Chat.ID, Username: u.Message.From.Username},
				Text:   "Game is not yet available, try again when it will release!",
			}
			_, err := b.SendMessage(&message)
			if err != nil {
				fmt.Println(err)
			}
			return
		}

		if isFree {
			message := telego.SendMessageParams{
				ChatID: telego.ChatID{ID: u.Message.Chat.ID, Username: u.Message.From.Username},
				Text:   "Game must be free!",
			}
			_, err := b.SendMessage(&message)
			if err != nil {
				fmt.Println(err)
			}
			return
		}

		if isOnSale {
			message := telego.SendMessageParams{
				ChatID: telego.ChatID{ID: u.Message.Chat.ID, Username: u.Message.From.Username},
				Text:   "Game already on sale.",
			}
			_, err := b.SendMessage(&message)
			if err != nil {
				fmt.Println(err)
			}
		}
		id := utils.IdTrimer(url)
		var (
			mb    bool
			mtext string
		)

		if !dbQueries.IsGame(id) {
			mb = dbQueries.AddNewGame(*u, id, title)
		} else {
			mb = dbQueries.UserToGame(*u, id)
		}

		if mb {
			mtext = "The Game was successfully added to ur list! You will be alerted about next sale!"
		} else {
			mtext = "Game already in ur list!"
		}

		message := telego.SendMessageParams{
			ChatID: telego.ChatID{ID: u.Message.Chat.ID, Username: u.Message.From.Username},
			Text:   mtext,
		}
		_, err := b.SendMessage(&message)
		if err != nil {
			fmt.Println(err)
		}

	})
}

func isURLValid(url string, cookies []*http.Cookie) (bool, string, bool, bool, bool) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Error creating HTTP request: %v", err)
		return false, "", false, false, false
	}

	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}

	response, err := client.Do(req)
	if err != nil {
		log.Printf("Error HTTP fetch: %v", err)
		return false, "", false, false, false
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(response.Body)

	if response.StatusCode != 200 {
		log.Printf("Wrong Status CODE: %d", response.StatusCode)
		return false, "", false, false, false
	}

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Printf("Error parsing page: %v", err)
		return false, "", false, false, false
	}

	title := document.Find(".apphub_AppName").First().Text()

	if document.Find(".apphub_AppName").Length() == 0 {
		log.Println("Can't find .apphub_AppName element, it might be Steam homepage")
		return false, "", false, false, false
	}

	purchaseBox := document.Find(".game_area_purchase").Find(".game_area_purchase_game")
	wishlistBox := document.Find(".wishlist_add_reminder").Text()
	if purchaseBox.Text() == "" && wishlistBox != "" {
		return true, title, false, false, true
	}

	var gameBox goquery.Selection

	if document.Find(".demo_above_purchase").Text() != "" {
		gameBox = *purchaseBox.First().Next().Find(".game_purchase_price")
	} else {
		gameBox = *purchaseBox.First().Find(".game_purchase_price")
	}

	isFree := gameBox.Text()

	if isFree == "Free to Play" {
		log.Println("Game must be free")
		return true, title, true, false, false
	}

	if isOnSale(*document, url) {
		return true, title, false, true, false
	}

	return true, title, false, false, false
}

func isOnSale(doc goquery.Document, url string) bool {
	var discCheck string

	purchaseBox := doc.Find(".game_area_purchase").Find(".game_area_purchase_game")

	if doc.Find(".demo_above_purchase").Text() != "" {
		discCheck = purchaseBox.First().Next().Find(".game_purchase_discount").Text()
	} else {
		discCheck = purchaseBox.First().Find(".game_purchase_discount").Text()
	}

	id := utils.IdTrimer(url)

	if dbQueries.OnSale(id) {
		if discCheck == "" {
			dbQueries.OnSaleChanger(1, id)
			return false
		}
		return true
	} else {
		if discCheck == "" {
			return false
		} else {
			dbQueries.OnSaleChanger(0, id)
			return true
		}
	}
}
