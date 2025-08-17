package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type URL struct {
	Id         string    `json:"id"`
	OriginURL  string    `json:"origin_url"`
	ShortURL   string    `json:"short_url"`
	CreateDate time.Time `json:"create_date"`
}

var DB = make(map[string]URL)

func genShortURL(url string) string {
	hasher := md5.New()
	hasher.Write([]byte(url))
	hash := hasher.Sum(nil)
	shortURL := fmt.Sprintf("%x", hash[:6])
	return shortURL
}

func saveURL(orgUrl string) string {
	shortURL := genShortURL(orgUrl)
	urlData := URL{
		Id:         shortURL,
		OriginURL:  orgUrl,
		ShortURL:   shortURL,
		CreateDate: time.Now(),
	}
	DB[shortURL] = urlData
	return shortURL
}

func getURL(shortURL string) (URL, error) {
	key, exists := DB[shortURL]
	if !exists {
		return URL{}, fmt.Errorf("URL not found")
	}
	return key, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the URL Shortener!\n")
}

func shortUrlHandler(w http.ResponseWriter, r *http.Request) {
	type Data struct {
		URL string `json:"url"`
	}
	var data Data
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if data.URL == "" {
		http.Error(w, "URL cannot be empty", http.StatusBadRequest)
		return
	}

	shortURL := saveURL(data.URL)
	type ResData struct {
		ShortURL string `json:"short_url"`
	}
	response := ResData{
		ShortURL: shortURL,
	}

	ResJson, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(ResJson)
}

func getit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/get/"):]
	if id == "" {
		http.Error(w, "Short URL ID is required", http.StatusBadRequest)
		return
	}
	urlData, err := getURL(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	http.Redirect(w, r, urlData.OriginURL, http.StatusSeeOther)
}

func main() {
	fmt.Println("Project: URL Shortener")
	fmt.Println("Sever is starting at PORT 8080")

	http.HandleFunc("/", handler)
	http.HandleFunc("/shorten", shortUrlHandler)
	http.HandleFunc("/get/", getit)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
}
