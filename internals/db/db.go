package db

import (
	"database/sql"
	"errors"
	"log"
	"math/rand"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var CHARSET string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

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

	// create table
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

func CreateId(url string) string {
	id := generateId()

	db, err := sql.Open("sqlite3", "tmp/shortener.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO urls(id, url) values(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	
	_, err = stmt.Exec(id, url)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Created id %s with url %s", id, url)
	return id
}

func CreateCustomId(id, url string) (string, error) {
	idExists := checkIdUsed(id)

	if idExists {
		return "", errors.New("Id already exists")
	}

	// TODO: finish
}

func GetUrl(id string) string {
	db, err := sql.Open("sqlite3", "tmp/shortener.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM urls WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var url string
	for rows.Next() {
		rows.Scan(&id, &url)
	}

	return url
}

/* ======= Private Methods =========== */
func checkIdUsed(id string) bool {
	db, err := sql.Open("sqlite3", "tmp/shortener.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM urls WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	return rows.Next()
}

func generateId() string {
	for ;; {
		id := make([]byte, 6)
		for i := 0; i < 6; i++ {
			id[i] = CHARSET[rand.Intn(len(CHARSET))]
		}

		if !checkIdUsed(string(id)) {
			return string(id)
		}
	}
}