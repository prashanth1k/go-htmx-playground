package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	P "example.com/go-htmx-playground/pages"
)

type Book struct {
	Title  string
	Author string
}

func main() {
	log.Println("Starting server at http://localhost:8080")

	http.HandleFunc("/hello", P.HelloHandler)

	http.Handle("/help/", http.StripPrefix("/help/", http.FileServer(http.Dir("./help"))))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	http.HandleFunc("/book-add", func(w http.ResponseWriter, r *http.Request) {

		title := r.PostFormValue("title")
		author := r.PostFormValue("author")
		log.Println("html request received.. " + title + " " + author)

		tmpl := template.Must(template.ParseFiles("pages/index.html"))
		tmpl.ExecuteTemplate(w, "book-item", Book{Title: title, Author: author})
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("/ request received")
		path := "pages/index.html"

		if r.URL.Path != "/" {
			path = "pages" + strings.TrimSuffix(r.URL.Path, "/") + ".html"
		}
		books := map[string][]Book{
			"Books": {
				{Title: "The Great Gatsby", Author: "F. Scott Fitzgerald"},
				{Title: "To Kill a Mockingbird", Author: "Harper Lee"},
				{Title: "1984", Author: "George Orwell"},
			},
		}
		tmpl := template.Must(template.ParseFiles(path))
		tmpl.Execute(w, books)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
