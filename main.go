package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	fs := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", index)

	log.Print("Starting server...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	log.Print(r.URL.Path)
	if r.URL.Path == "/" {
		http.Redirect(w, r, "/index.html", http.StatusSeeOther)
		return
	}

	index_path := filepath.Join("templates", "index.html")
	template_path := filepath.Join("templates", filepath.Clean(r.URL.Path))
	log.Print(template_path)

	tmpl, _ := template.ParseFiles(index_path, template_path)
	err := tmpl.ExecuteTemplate(w, "layout", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}
