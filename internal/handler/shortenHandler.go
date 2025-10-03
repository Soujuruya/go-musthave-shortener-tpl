package handler

import (
	"io"
	"net/http"
)

func ShortenHandler(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "cant read the body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	originalURL := string(body)

	id := urlService.Shortner(originalURL)

	shortURL := "http://localhost:8080/" + id

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(shortURL))
}
