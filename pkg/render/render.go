package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	tmplPath := fmt.Sprintf("./templates/%s", tmpl)
	baseTmplPath := fmt.Sprintf("./templates/%s", "base.layout.tmpl")
	parseTemplate, _ := template.ParseFiles(tmplPath, baseTmplPath)

	err := parseTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}
