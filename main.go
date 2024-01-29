package main

import (
	"net/http"
	"text/template"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	r.Get("/", welcome)
	http.ListenAndServe(":3000", r)
}

func welcome(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.New("").ParseFiles("templates/index.html")
	tmpl.ExecuteTemplate(w, "Base", nil)
}
