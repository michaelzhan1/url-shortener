package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func CreateDb() {
	log.Print("Creating db")

	// check tmp folder
	if _, err := os.Stat("tmp"); os.IsNotExist(err) {
		os.Mkdir("tmp", 0755)
	}

	// create db
	if _, err := os.Stat("tmp/shortener.db"); os.IsExist(err) {
		log.Print("Database already exists")
		return
	}

	file, err := os.Create("tmp/shortener.db")
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	db, err := sql.Open("sqlite3", "tmp/shortener.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// create table with id cou
	sqlStmt := `
		create table urls (
			"id" text not null primary key,
			"url" text not null
		)
	`

	log.Print("Creating urls table")
	stmt, err := db.Prepare(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}
	stmt.Exec()

	log.Print("Database created")
}

func Add(a, b int) int {
	return a + b
}