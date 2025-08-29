package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}

func generateShortCode() string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 6)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	for i := range b {
		b[i] = letters[int(b[i])%len(letters)]
	}
	return string(b)
}

func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ShortenRequest
	body, _ := io.ReadAll(r.Body)
	json.Unmarshal(body, &req)

	if req.URL == "" {
		http.Error(w, "URL required", http.StatusBadRequest)
		return
	}

	short := generateShortCode()
	_, err := db.Exec("INSERT INTO urls (short, original) VALUES (?, ?)", short, req.URL)
	if err != nil {
		http.Error(w, "Failed to save URL", http.StatusInternalServerError)
		return
	}

	res := ShortenResponse{
		ShortURL: fmt.Sprintf("http://localhost:8080/%s", short),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Path[1:]
	if code == "" {
		http.ServeFile(w, r, "../frontend/index.html")
		return
	}

	var original string
	err := db.QueryRow("SELECT original FROM urls WHERE short = ?", code).Scan(&original)
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, original, http.StatusFound)
}
