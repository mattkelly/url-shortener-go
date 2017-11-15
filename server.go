package main

import (
	"html/template"
	"net/http"
)

func main() {
	tmpl := template.Must(template.ParseFiles("url_form.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		longUrlFromForm := r.FormValue("long-url")

		tmpl.Execute(w, struct {
			Success bool
			LongUrl string
		}{true, longUrlFromForm})
	})

	http.ListenAndServe(":8080", nil)
}
