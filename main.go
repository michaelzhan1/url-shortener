package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/michaelzhan1/url-shortener/db"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
	if r.Method != "GET" {
		log.Printf("Method %s not allowed", r.Method)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "Hello World!")
}

func main() {
	fmt.Println(db.Add(1, 2))
	// port := 8080
	// portStr := fmt.Sprintf(":%d", port)
	// log.Println("Starting server on", portStr)
	// http.HandleFunc("/", indexHandler)
	// http.ListenAndServe(portStr, nil)

	// if err := http.ListenAndServe(portStr, nil); err != nil {
	// 	log.Fatal(err)
	// }
}