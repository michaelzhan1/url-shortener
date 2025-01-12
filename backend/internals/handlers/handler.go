package handlers

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/michaelzhan1/url-shortener/internals/db"
)

func NewUrlHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		log.Printf("Method %s not allowed", r.Method)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	tempUrl, _ := url.Parse(r.URL.String())
	params, _ := url.ParseQuery(tempUrl.RawQuery)

	if params["url"] == nil {
		http.Error(w, "Missing url parameter", http.StatusBadRequest)
		return
	}

	urlStr, _ := url.QueryUnescape(params["url"][0])

	id := db.CreateId(urlStr)

	fmt.Fprintf(w, "%s", id)
}

func NewCustomUrlHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		log.Printf("Method %s not allowed", r.Method)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	tempUrl, _ := url.Parse(r.URL.String())
	params, _ := url.ParseQuery(tempUrl.RawQuery)
	
	if params["url"] == nil || params["id"] == nil {
		http.Error(w, "Missing id or url parameter", http.StatusBadRequest)
		return
	}
	
	urlStr, _ := url.QueryUnescape(params["url"][0])
	id, _ := url.QueryUnescape(params["id"][0])
	
	id, _ = db.CreateCustomId(id, urlStr)
	
	fmt.Fprintf(w, "%s", id)
}

func UrlGetterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		log.Printf("Method %s not allowed", r.Method)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	tempUrl, _ := url.Parse(r.URL.String())
	params, _ := url.ParseQuery(tempUrl.RawQuery)
	
	if params["id"] == nil {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}
	
	id, _ := url.QueryUnescape(params["id"][0])
	
	url := db.GetUrl(id)
	
	fmt.Fprintf(w, "%s", url)
}

func DoNothingHandler(w http.ResponseWriter, r *http.Request) {
}