package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/michaelzhan1/url-shortener/internals/db"
	"github.com/michaelzhan1/url-shortener/internals/handlers"
)

func main() {
	db.CreateDb()
	testId := db.CreateId("https://www.google.com")
	log.Println("Test id:", testId)
	testUrl := db.GetUrl(testId)
	log.Println("Test url:", testUrl)
	port := 8080
	portStr := fmt.Sprintf(":%d", port)
	log.Println("Starting server on", portStr)

	http.HandleFunc("/", handlers.HelloWorldHandlerGet)
	http.ListenAndServe(portStr, nil)

	if err := http.ListenAndServe(portStr, nil); err != nil {
		log.Fatal(err)
	}
}