package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func HelloWorldHandlerGet(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
	if r.Method != "GET" {
		log.Printf("Method %s not allowed", r.Method)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "Hello World!")
}