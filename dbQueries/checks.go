package dbQueries

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func IsUser(id int) bool {
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT * FROM Users WHERE Id = ?)", id).Scan(&exists)
	if err != nil {
		fmt.Printf("Can't find user with id = %v: %v", id, err)
		return false
	}
	if !exists {
		return exists
	}

	var user map[string]interface{}
	query := "SELECT * FROM Users WHERE Id = ?"
	rows, err := db.Query(query, id)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		fmt.Println(err)
		return false
	}

	values := make([]interface{}, len(columns))
	valuePointers := make([]interface{}, len(columns))
	for i := range columns {
		valuePointers[i] = &values[i]
	}

	if rows.Next() {
		if err := rows.Scan(valuePointers...); err != nil {
			fmt.Println(err)
			return false
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

	return exists
}
