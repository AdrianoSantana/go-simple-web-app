package handlers

import (
	"net/http"

	"github.com/AdrianoSantana/go-simple-web-app/pkg/render"
)

// home is the homepage handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

// about is the aboutpage handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
