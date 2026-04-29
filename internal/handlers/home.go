package handlers

import (
	"html/template"
	"net/http"
)

func (h *Handler) Register(mux *http.ServeMux) {
	mux.HandleFunc("/", h.Home)
	mux.HandleFunc("/update", h.Update)
	mux.HandleFunc("/getValue", h.GetValue)
	mux.HandleFunc("/edit", h.CreateEditPage)
	mux.HandleFunc("/saveEdit", h.Edit)

	//авторизация
	mux.HandleFunc("/auth/register", h.RegisterPage)
	mux.HandleFunc("/auth/registerPost", h.RegisterUser)
	mux.HandleFunc("/auth/login", h.LoginPage)
	mux.HandleFunc("/auth/loginPost", h.Login)

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/home.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
