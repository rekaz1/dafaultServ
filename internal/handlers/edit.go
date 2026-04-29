package handlers

import (
	"html/template"
	"net/http"
	"strconv"
)

func (h *Handler) CreateEditPage(w http.ResponseWriter, r *http.Request) {
	idString := r.FormValue("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	panel, err := h.Store.GetPanelByID(r.Context(), id)
	if err != nil {
		http.Error(w, "Panel not found", http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("templates/edit.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, panel)
}

func (h *Handler) Edit(w http.ResponseWriter, r *http.Request) {

	// Получаем ID из формы
	idString := r.FormValue("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	// Получаем панель по ID
	panel, err := h.Store.GetPanelByID(r.Context(), id)
	if err != nil {
		http.Error(w, "Panel not found", http.StatusNotFound)
		return
	}

	err = h.Store.EditPanel(r.Context(), panel.ID, r.FormValue("title"), r.FormValue("description"))
	if err != nil {
		http.Error(w, "Database error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/getValue", http.StatusSeeOther)

}
