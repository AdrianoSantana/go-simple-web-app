package main

import (
	"fmt"
	"net/http"

	"github.com/AdrianoSantana/go-simple-web-app/pkg/handlers"
)

const PORT_NUMBER string = ":5005"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port %s\n", PORT_NUMBER)
	http.ListenAndServe(PORT_NUMBER, nil)
}
