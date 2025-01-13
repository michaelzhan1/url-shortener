package integration

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	r "github.com/michaelzhan1/url-shortener/internals/runtime"
)

func TestNewUrl(t *testing.T) {
	// setup
	loadDotEnv()
	defer dbCleanup()

	server := r.SetupServer()
	defer server.Close()

	// run server
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			t.Errorf("Server failed to start: %v", err)
		}
		t.Log("Server closed")
	}()
	time.Sleep(1 * time.Second)

	// make request
	goalUrl := "https://www.example.com"
	resp, err := http.Get(fmt.Sprintf("%s:%s/api/new?url=%s", os.Getenv("HOST"), os.Getenv("PORT"), url.QueryEscape(goalUrl)))
	if err != nil {
		t.Errorf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {	
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	// get id from body
	body, _ := io.ReadAll(resp.Body)
	id := string(body)
	
	// check redirect
	resp, err = http.Get(fmt.Sprintf("%s:%s/%s", os.Getenv("HOST"), os.Getenv("PORT"), id))
	if err != nil {
		t.Errorf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	location := resp.Request.URL.String()
	if location != goalUrl {
		t.Errorf("Expected location %s, got %s", goalUrl, location)
	}
}

func TestCustomUrl(t *testing.T) {
	// setup
	loadDotEnv()
	defer dbCleanup()

	server := r.SetupServer()
	defer server.Close()

	// run server
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			t.Errorf("Server failed to start: %v", err)
		}
		t.Log("Server closed")
	}()

	time.Sleep(1 * time.Second)

	goalId := "abcde"
	goalUrl := "https://www.example.com"
	resp, err := http.Get(fmt.Sprintf("%s:%s/api/new/custom?url=%s&id=%s", os.Getenv("HOST"), os.Getenv("PORT"), url.QueryEscape(goalUrl), goalId))
	if err != nil {
		t.Errorf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {	
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}
	
	// check redirect
	resp, err = http.Get(fmt.Sprintf("%s:%s/%s", os.Getenv("HOST"), os.Getenv("PORT"), goalId))
	if err != nil {
		t.Errorf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	location := resp.Request.URL.String()
	if location != goalUrl {
		t.Errorf("Expected location %s, got %s", goalUrl, location)
	}
}

func loadDotEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	if os.Getenv("HOST") == "" || os.Getenv("PORT") == "" {	
		log.Fatalf("HOST and PORT must be set in .env file")
	}
}

func dbCleanup() {
	os.Remove("tmp/shortener.db")
	os.Remove("tmp")
}