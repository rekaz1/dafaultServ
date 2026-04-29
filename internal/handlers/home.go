package handlers

import (
	"html/template"
	"net/http"
)

func Register(mux *http.ServeMux) {
	mux.HandleFunc("/", Home)
	mux.HandleFunc("/update", Update)
	mux.HandleFunc("/getValue", GetValue)

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

}

func Home(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	tmpl, err := template.ParseFiles("templates/home.html")
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

// func PanelPage(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json; charset=utf-8")

// }

type PanelData struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

var data = []PanelData{
	{ID: 1, Title: "Panel 1", Description: "This is the first panel."},
	{ID: 2, Title: "Panel 2", Description: "This is the second panel."},
	{ID: 3, Title: "Panel 3", Description: "This is the third panel."},
}
