package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type URLShortener struct {
	ID          string `json:"id"`
	OriginalURL string `json:"original_url"`
	ShortURL    string `json:"short_url"`
	CreatedAt   string `json:"created_at"`
}

var urlDB = make(map[string]URLShortener)

func CreateShortURL(originalURL string) string {
	hasher := md5.New()
	hasher.Write([]byte(originalURL))
	data := hasher.Sum(nil)
	hash := hex.EncodeToString(data)
	return hash[:8]
}

func SaveURL(originalURL string) URLShortener {
	shortURL := CreateShortURL(originalURL)
	urlEntry := URLShortener{
		ID:          shortURL,
		OriginalURL: originalURL,
		ShortURL:    shortURL,
		CreatedAt:   time.Now().Format(time.RFC3339), // Example timestamp
	}
	urlDB[shortURL] = urlEntry
	return urlEntry
}

func getURL(shortURL string) (URLShortener, bool) {
	if urlEntry, exists := urlDB[shortURL]; exists {
		return urlEntry, true
	}
	return URLShortener{}, false
}

func main() {
	fmt.Println("URL Shortener Service")
	// OriginalURL := "https://example.com/some/long/url"
	// urlEntry := SaveURL(OriginalURL)
	// fmt.Printf("Original URL: %s\n", urlEntry.OriginalURL)
	// fmt.Printf("Short URL: %s\n", urlEntry.ShortURL)

	// start the server
	fmt.Println("Starting server...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the URL Shortener Service!\n")
		fmt.Fprintf(w, "Use /shorten to create a short URL and /get to retrieve the original URL.\n")
	})

	http.HandleFunc("/shorten", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		var data struct {
			URL string `json:"url"`
		}
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		if data.URL == "" {
			http.Error(w, "Missing URL in request body", http.StatusBadRequest)
			return
		}
		urlEntry := SaveURL(data.URL)
		// fmt.Fprintf(w, "Short URL created: %s\n", urlEntry.ShortURL)
		response := map[string]string{
			"original_url": urlEntry.OriginalURL,
			"short_url":    urlEntry.ShortURL,
			"created_at":   urlEntry.CreatedAt,
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Error encoding response", http.StatusInternalServerError)
			return
		}
	})
	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		shortURL := r.URL.Path[len("/get/"):]

		if shortURL == "" {
			http.Error(w, "Missing short URL parameter", http.StatusBadRequest)
			return
		}
		urlEntry, exists := getURL(shortURL)
		if !exists {
			http.Error(w, "Short URL not found", http.StatusNotFound)
			return
		}
		// fmt.Fprintf(w, "Original URL: %s\n", urlEntry.OriginalURL)
		http.Redirect(w, r, urlEntry.OriginalURL, http.StatusFound)
	})

	error := http.ListenAndServe(":8080", nil)
	if error != nil {
		fmt.Printf("Error starting server: %v\n", error)
	}
}
