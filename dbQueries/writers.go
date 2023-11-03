package dbQueries

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mymmrac/telego"
)

func AddNewUser(u *telego.Update, curr string) {
	id := u.Message.Chat.ID
	username := u.Message.From.Username
	_, err := db.Query("INSERT INTO Users (Id, Username, Country_currency) VALUES (?,?,?)", id, username, curr)
	if err != nil {
		fmt.Println(err)
	}
}
