package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mattkelly/url-shortener-go/db"
	"github.com/mattkelly/url-shortener-go/shorten"
)

// TODO globals are lame
var tmpl *template.Template

func baseHandler(w http.ResponseWriter, r *http.Request) {
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
}

func lookupLongUrl(slug string) (string, error) {
	return db.Get(slug)
}

func shortUrlHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug := vars["slug"]

	longUrl, err := lookupLongUrl(slug)
	if err != nil {
		// TODO display error message to user
		panic(err)
	}

	http.Redirect(w, r, longUrl, http.StatusMovedPermanently)
}

func main() {
	db.Init()

	tmpl = template.Must(template.ParseFiles("url_form.html"))

	r := mux.NewRouter()

	r.HandleFunc("/", baseHandler)
	r.HandleFunc("/{slug:[a-zA-Z0-9]+}", shortUrlHandler)

	http.ListenAndServe(":3000", r)
}
