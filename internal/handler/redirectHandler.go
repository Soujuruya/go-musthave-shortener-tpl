package handler

import (
	"net/http"
	"strings"
)

func RedirectUrlHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/") // берём всё после "/"

	originalURL, ok := urlService.GetOriginal(id)
	if !ok {
		http.Error(w, "not found", http.StatusBadRequest)
		return
	}

	// Устанавливаем редирект
	http.Redirect(w, r, originalURL, http.StatusTemporaryRedirect)
}
