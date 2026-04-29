package handlers

import (
	"html/template"
	"net/http"
)

func (h *Handler) GetValue(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/getvalue.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	panels, err := h.Store.GetPanels(r.Context())
	if err != nil {
		http.Error(w, "Database error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, panels)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
