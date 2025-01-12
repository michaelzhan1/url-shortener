package db

import (
	"database/sql"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestCreateDb(t *testing.T) {
    defer cleanup()

    CreateDb()

    // check that db is created
    if _, err := os.Stat("tmp/shortener.db"); os.IsNotExist(err) {
        t.Error("Database not created")
    }

    // check that db has the right columns
    db, err := sql.Open("sqlite3", "tmp/shortener.db")
    if err != nil {
        t.Error(err)
    }
    defer db.Close()

    rows, err := db.Query("PRAGMA table_info(urls)")
    if err != nil {
        t.Error(err)
    }
    defer rows.Close()

    var columns []string
    for rows.Next() {
        var cid int
        var name string
        var _type string
        var notnull int
        var dflt_value interface{}
        var pk int

        rows.Scan(&cid, &name, &_type, &notnull, &dflt_value, &pk)
        columns = append(columns, name)
    }

    if len(columns) != 2 {
        t.Error("Wrong columns")
    }

    if columns[0] != "id" || columns[1] != "url" {
        t.Error("Wrong columns")
    }
}

func TestCreateId(t *testing.T) {
    setup()
    defer cleanup()

    url := "http://example.com"
    id := CreateId(url)

    if id == "" {
        t.Error("Id not created")
    }

    db, err := sql.Open("sqlite3", "tmp/shortener.db")
    if err != nil {
        t.Error(err)
    }
    defer db.Close()

    rows, err := db.Query("SELECT * FROM urls WHERE id = ?", id)
    if err != nil {
        t.Error(err)
    }
    defer rows.Close()

    for rows.Next() {
        var _id string
        var _url string
        rows.Scan(&_id, &_url)

        if _url != url {
            t.Error("Wrong url")
        }
    }
}

func TestRemoveId(t *testing.T) {
    setup()
    defer cleanup()

    url := "http://example.com"
    id := "12345"

    db, err := sql.Open("sqlite3", "tmp/shortener.db")
    if err != nil {
        t.Error(err)
    }
    defer db.Close()

    _, err = db.Exec("INSERT INTO urls (id, url) VALUES (?, ?)", id, url)
    if err != nil {
        t.Error(err)
    }

    RemoveId(id)

    rows, err := db.Query("SELECT * FROM urls WHERE id = ?", id)
    if err != nil {
        t.Error(err)
    }
    defer rows.Close()

    if rows.Next() {
        t.Error("Id not removed")
    }
}

func TestCreateCustomId(t *testing.T) {
    setup()
    defer cleanup()

    url := "http://example.com"
    id := "12345"

    _id, err := CreateCustomId(id, url)
    if err != nil {
        t.Error(err)
    }

    if _id != id {
        t.Error("Wrong id")
    }

    db, err := sql.Open("sqlite3", "tmp/shortener.db")
    if err != nil {
        t.Error(err)
    }
    defer db.Close()

    rows, err := db.Query("SELECT * FROM urls WHERE id = ?", id)
    if err != nil {
        t.Error(err)
    }
    defer rows.Close()

    for rows.Next() {
        var _id string
        var _url string
        rows.Scan(&_id, &_url)

        if _url != url {
            t.Error("Wrong url")
        }
    }
}

func TestGetUrl(t *testing.T) {
    setup()
    defer cleanup()

    url := "http://example.com"
    id := "12345"

    db, err := sql.Open("sqlite3", "tmp/shortener.db")
    if err != nil {
        t.Error(err)
    }
    defer db.Close()

    _, err = db.Exec("INSERT INTO urls (id, url) VALUES (?, ?)", id, url)
    if err != nil {
        t.Error(err)
    }

    _url := GetUrl(id)
    if _url != url {
        t.Error("Wrong url")
    }
}

func TestCheckIdUsed(t *testing.T) {
    setup()
    defer cleanup()

    url := "http://example.com"
    id := "12345"

    CreateCustomId(id, url)

    if !checkIdUsed(id) {
        t.Error("Already-used Id not flagged")
    }
}

func setup() {
    CreateDb()
}

func cleanup() {
    os.Remove("tmp/shortener.db")
    os.Remove("tmp")
}