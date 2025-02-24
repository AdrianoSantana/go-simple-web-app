package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc(
		"/",
		func(w http.ResponseWriter, r *http.Request) {
			n, err := fmt.Fprintf(w, "<html> <h1> Hello, world! </h1> </html>")
			if err != nil {
				panic(err)
			}

			fmt.Printf("Bytes Written %d\n", n)
		},
	)

	http.ListenAndServe(":5005", nil)
}
