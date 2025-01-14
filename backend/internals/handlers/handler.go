package handlers

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/michaelzhan1/url-shortener/internals/db"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
}

func NewUrlHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
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
	enableCors(&w)
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
	enableCors(&w)
	if r.Method != "GET" {
		log.Printf("Method %s not allowed", r.Method)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.PathValue("id")
	
	url, _ := url.QueryUnescape(db.GetUrl(id))

	http.Redirect(w, r, url, http.StatusSeeOther)
}