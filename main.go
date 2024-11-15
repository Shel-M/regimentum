package main

import (
	"html/template"
	"net/http"
)

type Templ struct {
	name     string
	template *template.Template
}

var templates []Templ

func main() {
	http.HandleFunc("/", index)
	t, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		println("error {}", err.Error())
		return
	}
	templates = append(templates, Templ{"index", t})

	http.ListenAndServe(":3000", nil)
}

func index(w http.ResponseWriter, _ *http.Request) {
	for _, template := range templates {
		if template.name == "index" {
			template.template.Execute(w, "")
			return
		}
	}
}
