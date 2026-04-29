package handlers

import (
	"net/http"
	"simpleserver/internal/storage"
)

type Handler struct {
	Store *storage.Storage
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	description := r.FormValue("description")
	if title == "" || description == "" {
		http.Error(w, "Title and description are required", http.StatusBadRequest)
		return
	}
	err := h.Store.CreatePanel(r.Context(), title, description)
	if err != nil {
		http.Error(w, "Database error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
