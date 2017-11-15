package main

import (
	"html/template"
	"net/http"

	"github.com/mattkelly/url-shortener-go/shorten"
)

func main() {
	tmpl := template.Must(template.ParseFiles("url_form.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		longUrl := r.FormValue("long-url")

		shortUrl := shorten.Shorten(longUrl)

		tmpl.Execute(w, struct {
			Success  bool
			ShortUrl string
		}{true, shortUrl})
	})

	http.ListenAndServe(":8080", nil)
}
