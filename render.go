package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	tmplPath := fmt.Sprintf("./templates/%s", tmpl)
	parseTemplate, _ := template.ParseFiles(tmplPath)

	err := parseTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}
