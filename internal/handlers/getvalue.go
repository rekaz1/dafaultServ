package handlers

import (
	"html/template"
	"net/http"
)

func GetValue(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/getvalue.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
