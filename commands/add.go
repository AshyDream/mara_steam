package commands

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/mymmrac/telego"
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
				Text:   "Pls proceed argument",
			}
			b.SendMessage(&message)
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

		is, title := isURLValid(url, cookies)

		if !is {
			message := telego.SendMessageParams{
				ChatID: telego.ChatID{ID: u.Message.Chat.ID, Username: u.Message.From.Username},
				Text:   "Invalid Steam game URL!",
			}
			b.SendMessage(&message)
			return
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
			mtext = "The Game was successfully added to ur list!"
		} else {
			mtext = "Game already in ur list!"
		}

		message := telego.SendMessageParams{
			ChatID: telego.ChatID{ID: u.Message.Chat.ID, Username: u.Message.From.Username},
			Text:   mtext,
		}
		b.SendMessage(&message)

		//handlers.Parser(url, u, b)
	})
}

func isURLValid(url string, cookies []*http.Cookie) (bool, string) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Error creating HTTP request: %v", err)
		return false, ""
	}

	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}

	response, err := client.Do(req)
	if err != nil {
		log.Printf("Error HTTP fetch: %v", err)
		return false, ""
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		log.Printf("Wrong Status CODE: %d", response.StatusCode)
		return false, ""
	}

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Printf("Error parsing page: %v", err)
		return false, ""
	}

	title := document.Find(".apphub_AppName").First().Text()

	if document.Find(".apphub_AppName").Length() == 0 {
		log.Println("Can't find .apphub_AppName element, it might be Steam homepage")
		return false, ""
	}

	return true, title
}
