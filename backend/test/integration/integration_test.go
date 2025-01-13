package integration

import (
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

// TODO: use httptest
func TestNewUrl(t *testing.T) {
	loadDotEnv()
	defer cleanup()

	server := r.SetupServer()
	defer server.Close()

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			t.Errorf("Server failed to start: %v", err)
		}
		t.Log("Server closed")
	}()

	time.Sleep(1 * time.Second)
	goalUrl := "https://www.example.com"
	httpUrl := os.Getenv("HOST") + ":" + os.Getenv("PORT") + "/api/new?url=" + url.QueryEscape(goalUrl)

	resp, err := http.Get(httpUrl)
	if err != nil {
		t.Errorf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {	
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	// get text from body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Failed to read response body: %v", err)
	}

	id := string(body)
	
	// get url back from db
	httpUrl = os.Getenv("HOST") + ":" + os.Getenv("PORT") + "/" + id
	resp, err = http.Get(httpUrl)
	if err != nil {
		t.Errorf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	// get text from body
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Failed to read response body: %v", err)
	}
	recoveredUrl, _ := url.QueryUnescape(string(body))

	if recoveredUrl != goalUrl {
		t.Errorf("Expected url %s, got %s", goalUrl, recoveredUrl)
	}
}

func TestCustomUrl(t *testing.T) {
	loadDotEnv()
	defer cleanup()

	server := r.SetupServer()
	defer server.Close()

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			t.Errorf("Server failed to start: %v", err)
		}
		t.Log("Server closed")
	}()

	time.Sleep(1 * time.Second)
	goalId := "abcde"
	goalUrl := "https://www.example.com"
	httpUrl := os.Getenv("HOST") + ":" + os.Getenv("PORT") + "/api/new/custom?url=" + url.QueryEscape(goalUrl) + "&id=" + goalId

	resp, err := http.Get(httpUrl)
	if err != nil {
		t.Errorf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {	
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}
	
	// get url back from db
	httpUrl = os.Getenv("HOST") + ":" + os.Getenv("PORT") + "/" + goalId
	resp, err = http.Get(httpUrl)
	if err != nil {
		t.Errorf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	// get text from body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Failed to read response body: %v", err)
	}
	recoveredUrl, _ := url.QueryUnescape(string(body))

	if recoveredUrl != goalUrl {
		t.Errorf("Expected url %s, got %s", goalUrl, recoveredUrl)
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

func cleanup() {
	os.Remove("tmp/shortener.db")
	os.Remove("tmp")
}