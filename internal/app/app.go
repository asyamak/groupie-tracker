package app

import (
	"log"
	"net/http"

	"groupie-tracker/internal/handlers"
)

const (
	port = ":8080"
)

func App() error {
	mux := http.NewServeMux()
	h := handlers.NewHandler()
	mux.HandleFunc("/", h.GETHandler)
	mux.HandleFunc("/artist/", h.POSTHandler)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("ui/static"))))
	log.Println("Starting server at port :8080\nhttp://localhost:8080")
	return http.ListenAndServe(port, mux)
}
