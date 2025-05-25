package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// renderTemplate renders go template and writes to an http.ResponseWriter
func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}
