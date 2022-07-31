package handlers

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"groupie/internal/logic"
)

// Home function is a main page of the site, which displays list of artist
// their picture and link for more information.
func Home(w http.ResponseWriter, r *http.Request) {
	status := CheckMethod(r, "/", http.MethodGet)
	if status != 200 {
		ErrorHandler(w, status)
		return
	}
	entries := logic.GetArtist()
	Execute(w, "ui/templates/index.html", entries)
}

// Artist function is responsive for artist's info.
func Artist(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/artist/")
	status := CheckMethod(r, "/artist/"+id, http.MethodGet)
	if status != 200 {
		ErrorHandler(w, status)
		log.Printf("Error %d - User tried to use not allowed method - %s \n", http.StatusMethodNotAllowed, r.Method)
		return
	}
	_, err := strconv.Atoi(id)
	if err != nil {
		ErrorHandler(w, http.StatusNotFound)
		return
	}

	entries := logic.GetMembers(id)
	if entries.ID == 0 {
		ErrorHandler(w, http.StatusNotFound)
		log.Printf("Error %d - Page (http://localhost:8080%s) not found", http.StatusNotFound, r.URL.Path)
		return
	}
	Execute(w, "ui/templates/artist.html", entries)
}

// Relation function is responsive for displaying information about location and date artist had.
func Relation(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/relations/")
	status := CheckMethod(r, "/relations/"+id, http.MethodGet)
	if status != 200 {
		ErrorHandler(w, status)
		log.Printf("Error %d - User tried to use not allowed method - %s \n", http.StatusMethodNotAllowed, r.Method)
		return
	}
	_, err := strconv.Atoi(id)
	if err != nil {
		ErrorHandler(w, http.StatusNotFound)
		log.Printf("Error %d - User tried to use not allowed method - %s \n", http.StatusNotFound, r.Method)
		return
	}
	entries := logic.Relation(id)
	if entries.ID == 0 {
		ErrorHandler(w, http.StatusBadRequest)
		log.Printf("Error %d - Page (http://localhost:8080%s) not found", http.StatusNotFound, r.URL.Path)
		return
	}
	Execute(w, "ui/templates/concert.html", entries)
}

// func SearchBar(w http.ResponseWriter, r *http.Request) {
// 	status := CheckMethod(r, "/", http.MethodGet)
// 	if status != 200 {
// 		ErrorHandler(w, status)
// 		return
// 	}
// 	// search := r.FormValue("send")
// 	entries := logic.SearchBar(search)
// 	Execute(w, "ui/templates/index.html", entries)
// }
