package dbQueries

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func IsUser(id int) bool {
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT * FROM Users WHERE Id = ?)", id).Scan(&exists)
	if err != nil {
		fmt.Printf("Can't find user with id = %v: %v", id, err)
		return false
	}

	return exists
}

func GameInfo(id int) map[string]interface{} {
	var game map[string]interface{}
	query := "SELECT * FROM Games WHERE Id = ?"
	rows, err := db.Query(query, id)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		fmt.Println(err)
	}

	values := make([]interface{}, len(columns))
	valuePointers := make([]interface{}, len(columns))
	for i := range columns {
		valuePointers[i] = &values[i]
	}

	if rows.Next() {
		if err := rows.Scan(valuePointers...); err != nil {
			fmt.Println(err)
		}

		game = make(map[string]interface{})
		for i, colName := range columns {
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				game[colName] = string(b)
			} else {
				game[colName] = val
			}
		}
	}
	return game
}

func UserInfo(id int) {
	var user map[string]interface{}
	query := "SELECT * FROM Users WHERE Id = ?"
	rows, err := db.Query(query, id)
	if err != nil {
		fmt.Println(err)
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	columns, err := rows.Columns()
	if err != nil {
		fmt.Println(err)
	}

	values := make([]interface{}, len(columns))
	valuePointers := make([]interface{}, len(columns))
	for i := range columns {
		valuePointers[i] = &values[i]
	}

	if rows.Next() {
		if err := rows.Scan(valuePointers...); err != nil {
			fmt.Println(err)
		}

		user = make(map[string]interface{})
		for i, colName := range columns {
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				user[colName] = string(b)
			} else {
				user[colName] = val
			}
		}
	}

	fmt.Println(user)
}

func IsGame(id int) bool {
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT * FROM Games WHERE Id = ?)", id).Scan(&exists)
	if err != nil {
		fmt.Printf("Can't find user with id = %v: %v", id, err)
		return false
	}

	return exists
}

func IsData(userId int, gameId int) bool {
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM Games_Users WHERE (User_id, Game_id) = (?,?))", userId, gameId).Scan(&exists)
	if err != nil {
		fmt.Println(err)
		return true
	}

	return exists
}

func OnSale(gameId int) bool {
	row, err := db.Query("SELECT OnSale FROM Games WHERE Id = ?", gameId)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(row)

	defer row.Close()

	onSale := make([]int, 0)
	for row.Next() {
		var val int
		if err := row.Scan(&val); err != nil {
			log.Fatal(err)
		}
		onSale = append(onSale, val)
	}
	if err := row.Err(); err != nil {
		log.Fatal(err)
	}

	on := onSale[0]

	if on == 0 {
		return false
	} else {
		return true
	}
}

func List(userId int) map[int]int {
	var games map[int]int
	query := "SELECT id, Game_id, User_id FROM Games_Users WHERE User_id = ?"
	rows, err := db.Query(query, userId)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, gameID, userID int
		err := rows.Scan(&id, &gameID, &userID)
		if err != nil {
			fmt.Println(err)
		}

		if games == nil {
			games = make(map[int]int)
		}

		games[id] = gameID
	}

	return games
}

func GetGamesGroup(lastId int, limit int) []map[string]interface{} {
	var games []map[string]interface{}
	var game map[string]interface{}
	query := "SELECT * FROM Games WHERE Id>? LIMIT ?"
	rows, err := db.Query(query, lastId, limit)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		fmt.Println(err)
	}

	values := make([]interface{}, len(columns))
	valuePointers := make([]interface{}, len(columns))
	for i := range columns {
		valuePointers[i] = &values[i]
	}

	for rows.Next() {
		if err := rows.Scan(valuePointers...); err != nil {
			fmt.Println(err)
		}

		game = make(map[string]interface{})
		for i, colName := range columns {
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				game[colName] = string(b)
			} else {
				game[colName] = val
			}
		}
		games = append(games, game)
	}
	return games
}

func GetUserGroup(lastId int, limit int, gameId int) []map[string]interface{} {
	var users []map[string]interface{}
	var user map[string]interface{}
	query := "SELECT * FROM Games_Users WHERE Game_id=? AND User_id>? LIMIT ?"
	rows, err := db.Query(query, gameId, lastId, limit)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		fmt.Println(err)
	}

	values := make([]interface{}, len(columns))
	valuePointers := make([]interface{}, len(columns))
	for i := range columns {
		valuePointers[i] = &values[i]
	}

	for rows.Next() {
		if err := rows.Scan(valuePointers...); err != nil {
			fmt.Println(err)
		}

		user = make(map[string]interface{})
		for i, colName := range columns {
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				user[colName] = string(b)
			} else {
				user[colName] = val
			}
		}
		users = append(users, user)
	}
	return users
}

func GetUserCurr(userId int) string {
	var curr string
	query := "SELECT Country_currency FROM Users WHERE Id=?"
	row, err := db.Query(query, userId)
	if err != nil {
		fmt.Println(err)
	}
	defer row.Close()
	if row.Next() {
		if err := row.Scan(&curr); err != nil {
			fmt.Println(err)
		}
	}
	return curr
}
