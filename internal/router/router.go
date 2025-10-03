package router

import (
	"net/http"

	"go-musthave-shortener-tpl/internal/handler"
)

type Router struct {
	server *http.Server
}

func New() *Router {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			handler.ShortenHandler(w, r)
			return
		}
		if r.Method == http.MethodGet {
			handler.RedirectUrlHandler(w, r)
			return
		}
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	})

	return &Router{server: &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}}

}

func (r *Router) Run() error {
	return r.server.ListenAndServe()
}
