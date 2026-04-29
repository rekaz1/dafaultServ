package handlers

import (
	"net/http"
)

func Update(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	description := r.FormValue("description")
	if title == "" || description == "" {
		http.Error(w, "Title and description are required", http.StatusBadRequest)
		return
	}
	data = append(data, PanelData{ID: len(data) + 1, Title: title, Description: description})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
