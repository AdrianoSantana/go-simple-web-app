package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const PORT_NUMBER string = ":5005"

// home is the homepage handler
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

// about is the aboutpage handler
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	tmplPath := fmt.Sprintf("./templates/%s", tmpl)
	parseTemplate, _ := template.ParseFiles(tmplPath)

	err := parseTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Starting application on port %s\n", PORT_NUMBER)
	http.ListenAndServe(PORT_NUMBER, nil)
}
