package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func templateExecution(w http.ResponseWriter, parseFile string, data interface{}) {
	html, err := template.ParseFiles(parseFile)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "./ui/html/500.html", http.StatusInternalServerError)
		return
	}
	err = html.Execute(w, data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "./ui/html/500.html", http.StatusInternalServerError)
		return
	}
}
