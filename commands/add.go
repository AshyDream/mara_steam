package commands

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/mymmrac/telego"
	"log"
	"mara/handlers"
	"net/http"
	"strings"
)

func init() {
	handlers.Register("add", func(u *telego.Update, b *telego.Bot) {
		text := u.Message.Text
		text = handlers.TrimFirstRune(text)
		parts := strings.Fields(text)
		url := parts[1]
		urlExmpl := "https://store.steampowered.com/app/"

		if len(parts) != 2 || !strings.Contains(url, urlExmpl) {
			message := telego.SendMessageParams{
				ChatID: telego.ChatID{ID: u.Message.Chat.ID, Username: u.Message.From.Username},
				Text:   "Wrong arguments!",
			}
			b.SendMessage(&message)
			return
		}

		if !isURLValid(url) {
			message := telego.SendMessageParams{
				ChatID: telego.ChatID{ID: u.Message.Chat.ID, Username: u.Message.From.Username},
				Text:   "Invalid Steam game URL!",
			}
			b.SendMessage(&message)
			return
		}

		handlers.Parser(url, u, b)

	})
}

func isURLValid(url string) bool {
	response, err := http.Get(url)
	if err != nil {
		log.Printf("Помилка HTTP запиту: %v", err)
		return false
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		log.Printf("Недійсний статус код: %d", response.StatusCode)
		return false
	}

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Printf("Помилка парсингу сторінки: %v", err)
		return false
	}

	if document.Find(".apphub_AppName").Length() == 0 {
		log.Println("Елемент .apphub_AppName не знайдено, це може бути головна сторінка Steam")
		return false
	}

	return true
}
