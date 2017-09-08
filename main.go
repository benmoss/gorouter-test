package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Host: %q", html.EscapeString(r.Host))
	})

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
