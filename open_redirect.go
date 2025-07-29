package main

import (
	"net/http"
)

func handleGet(w http.ResponseWriter, r *http.Request) {
	location := r.URL.Query().Get("url")
	http.Redirect(w, r, location, http.StatusFound) // Noncompliant
}
