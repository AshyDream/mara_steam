package dbQueries

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mymmrac/telego"
)

func AddNewUser(u telego.Update, curr string) {
	id := u.CallbackQuery.Message.Chat.ID
	username := u.CallbackQuery.Message.From.Username
	if username == "" {
		username = u.CallbackQuery.Message.From.FirstName
	}
	_, err := db.Query("INSERT INTO Users (Id, Username, Country_currency) VALUES (?,?,?)", id, username, curr)
	if err != nil {
		fmt.Println(err)
	}
	UserInfo(int(id))
}

func UpdateCurr(u telego.Update, curr string) {
	id := u.CallbackQuery.Message.Chat.ID
	_, err := db.Query("UPDATE Users SET Country_currency = ? WHERE Id = ?", curr, id)
	if err != nil {
		fmt.Println(err)
	}
	UserInfo(int(id))
}

func AddNewGame(u telego.Update, id int, title string) bool {
	_, err := db.Query("INSERT INTO Games (Id, Title) VALUES (?,?)", id, title)
	if err != nil {
		fmt.Println(err)
	}
	return UserToGame(u, id)
}

func UserToGame(u telego.Update, gameId int) bool {
	userId := u.Message.Chat.ID
	if IsData(int(userId), gameId) {
		return false
	}
	_, err := db.Query("INSERT INTO Games_Users (Game_Id, User_Id) VALUES (?,?)", gameId, userId)
	if err != nil {
		fmt.Println(err, "\n writers 46")
		return false
	}
	return true
}

func OnSaleChanger(onSale int, gameId int) {
	var n int
	if onSale == 1 {
		n = 0
	} else {
		n = 1
	}
	_, err := db.Query("UPDATE  Games SET OnSale = ? WHERE Id = ?", n, gameId)
	if err != nil {
		fmt.Println(err, "\n writers 61")
		return
	}
}

func DeleteUserGame(u telego.Update, gameId int) bool {
	userId := u.CallbackQuery.Message.Chat.ID
	if !IsData(int(userId), gameId) {
		return false
	}
	_, err := db.Query("DELETE FROM Games_Users Where (Game_id, User_id) = (?,?)", gameId, userId)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
