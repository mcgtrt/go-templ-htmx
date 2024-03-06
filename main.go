package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	component := hello("SIEMANKO ZIOMECZKI")
	header := headerTemplate("To jest Szablon mój")

	http.Handle("/", templ.Handler(component))
	http.Handle("/cosik", templ.Handler(header))

	fmt.Println("Listening on:3000")
	http.ListenAndServe(":3000", nil)
}
