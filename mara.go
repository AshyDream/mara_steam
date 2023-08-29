package main

import (
	"fmt"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	"mara/commands"
	"mara/handlers"
	"os"
	"time"
)

var (
	cfg      = handlers.Cfg()
	botToken = cfg.Token
)

func main() {
	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println("Error-starting: ", err)
		os.Exit(1)
	}

	//botUser, err := bot.GetMe()
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}

	commands.RegisterAllCommands()

	//fmt.Printf("Bot user: %+v\n", botUser)

	// Get updates channel
	options := telego.GetUpdatesParams{Timeout: 60, Limit: 1}
	updates, _ := bot.UpdatesViaLongPolling(&options)

	// Create bot handler
	bh, _ := th.NewBotHandler(bot, updates)

	bh.Handle(func(b *telego.Bot, u telego.Update) {
		handlers.HandleCommand(&u, b)
	}, th.AnyCommand())

	bh.Handle(func(bot *telego.Bot, u telego.Update) {
		if u.Message.From.Username != "" {
			fmt.Printf("[\033[1m\033[31mDate:%s\033[0m]\033[0m\033[3m\033[36m @%s:\033[32m%s\n\033[0m", time.DateTime, u.Message.From.Username, u.Message.Text)
		} else {
			fmt.Printf("[\033[1m\033[31mDate:%s\033[0m]\033[0m\033[3m\033[36m %s:\033[32m%s\n\033[0m", time.DateTime, u.Message.From.FirstName, u.Message.Text)
		}
	}, th.AnyMessage())

	defer bh.Stop()
	defer bot.StopLongPolling()

	// Start handling updates
	bh.Start()

}
