package models

import (
	"database/sql"
	"fmt"

	"os"

	_ "github.com/mattn/go-sqlite3"
)

// defaultDatabaseFile is default database file
const defaultDatabaseFile = "./data/database.db"

// InitDatabase is init database connection
func InitDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", defaultDatabaseFile)
	if err != nil {
		fmt.Println(err)
	}
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS installation (is_already_installed boolean, access_token text);
	CREATE TABLE IF NOT EXISTS container (id integer not null primary key autoincrement, container_id text, container_name text, image_name text, author text);
	CREATE TABLE IF NOT EXISTS network (id integer not null primary key autoincrement, network_id text, network_name text, author text);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		fmt.Println(err)
	}
	return db
}

// DeleteDatabaseFile is remove database file
func DeleteDatabaseFile() {
	err := os.Remove(defaultDatabaseFile)
	if err != nil {
		fmt.Println(err)
	}
}
