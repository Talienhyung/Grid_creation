package main

import (
	GridCreator "GridCreator/function"
	"net/http"
)

func main() {
	GridCreator.Root()
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static", fs))
	http.ListenAndServe(":8080", nil)
}
