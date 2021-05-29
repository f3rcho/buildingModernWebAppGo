package render

import (
	"fmt"
	"html/template"
	"net/http"
)

// RenderTemplate is the handler to render and parse templates
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error rendering template", err)
		return
	}
}
