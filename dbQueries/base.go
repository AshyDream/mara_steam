package dbQueries

import (
	"database/sql"
	"fmt"
	"mara/utils"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbSource = utils.DbSource
	db       *sql.DB
)

func InitDB() error {
	dataSourceName := dbSource
	var err error
	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		return err
	}

	db.SetMaxOpenConns(0)
	db.SetMaxIdleConns(0)
	db.SetConnMaxLifetime(1 * time.Minute)

	err = db.Ping()
	if err != nil {
		return err
	}

	return nil
}

func ShowDB() {
	rows, err := db.Query("SHOW TABLES")
	if err != nil {
		fmt.Printf("Error getting Database structure: %v", err)
		return
	}

	defer rows.Close()

	var tableName string
	for rows.Next() {
		err := rows.Scan(&tableName)
		if err != nil {
			fmt.Printf("Error scanning row: %v", err)
			return
		}
		fmt.Println("\033[34m", tableName)
		showTable(tableName)
	}
}

func showTable(name string) {
	fmt.Println("\033[33m===========================================")
	fmt.Println("\033[33mField  Type NULL Key DefaultVal Extra")
	rows, err := db.Query("DESCRIBE " + name)
	if err != nil {
		fmt.Printf("Error describe table - %v : %v", name, err)
		return
	}

	defer rows.Close()

	var (
		field      string
		typ        string
		null       string
		key        string
		defaultVal sql.NullString
		extra      string
	)

	for rows.Next() {
		err := rows.Scan(&field, &typ, &null, &key, &defaultVal, &extra)
		if err != nil {
			fmt.Printf("Error scanning row: %v", err)
			return
		}
		fmt.Println(field, typ, null, key, defaultVal, extra)
	}
	fmt.Println("\033[33m===========================================\n")
}
