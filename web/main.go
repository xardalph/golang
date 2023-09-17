// Copyright 2010 The Go Authors. All rights reserved.

// Use of this source code is governed by a BSD-style

// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
)

var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

type Page struct {
	Title string

	Body []byte
}

func (p *Page) save() error {

	filename := p.Title + ".txt"

	return os.WriteFile(filename, p.Body, 0600)

}

func loadPage(title string) (*Page, error) {

	filename := title + ".txt"

	body, err := os.ReadFile(filename)

	if err != nil {

		return nil, err

	}

	return &Page{Title: title, Body: body}, nil

}
func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		fmt.Println("error when saving a document : ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {

	title := r.URL.Path[len("/view/"):]

	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}

	renderTemplate(w, "view", p)

}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9-]+)$")

func getTitle(w http.ResponseWriter, r *http.Request) string {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.Error(w, "invalid Page Title : "+r.URL.Path, http.StatusForbidden)
		return ""
	}
	return m[2] // The title is the second subexpression.
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := getTitle(w, r)
	if title == "" {
		return
	}
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
