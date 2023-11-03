package dbQueries

import (
	"database/sql"
	"fmt"
	"mara/handlers"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	cfg      = handlers.Cfg()
	dbSource = cfg.DBSource
	db       *sql.DB
)

func InitDB() error {
	dataSourceName := dbSource
	var err error
	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		return err
	}

	// Встановлення параметрів з'єднання
	db.SetMaxOpenConns(5)                  // Максимальна кількість відкритих з'єднань
	db.SetMaxIdleConns(0)                  // Максимальна кількість неактивних з'єднань у пулі
	db.SetConnMaxLifetime(3 * time.Minute) // Максимальний час життя з'єднання

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