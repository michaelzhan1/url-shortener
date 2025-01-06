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
	db.CreateId("https://www.google.com")
	port := 8080
	portStr := fmt.Sprintf(":%d", port)
	log.Println("Starting server on", portStr)

	http.HandleFunc("/", handlers.HelloWorldHandlerGet)
	http.ListenAndServe(portStr, nil)

	if err := http.ListenAndServe(portStr, nil); err != nil {
		log.Fatal(err)
	}
}