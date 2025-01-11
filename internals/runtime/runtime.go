package runtime

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/michaelzhan1/url-shortener/internals/db"
	"github.com/michaelzhan1/url-shortener/internals/handlers"
)

// change to SetupServer() to return a server object, then call start and close later
func SetupServer() *http.Server {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	db.CreateDb()
	port := ":" + os.Getenv("PORT")
	log.Println("Starting server on", port)

	mux := http.NewServeMux()
	mux.HandleFunc("/new", handlers.NewUrlHandler)
	mux.HandleFunc("/new/custom", handlers.NewCustomUrlHandler)
	mux.HandleFunc("/get", handlers.UrlGetterHandler)
	mux.HandleFunc("/favicon.ico", handlers.DoNothingHandler)

	server := &http.Server{
		Addr:    port,
		Handler: mux,
	}

	return server
}