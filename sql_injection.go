package main

import (
	"database/sql"
	"fmt"
	"net/http"
)

type User struct {
	ID       int
	Username string
}

func authenticate(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	var user User
	row := db.QueryRow("SELECT * FROM users WHERE username='" + username + "' AND password = '" + password + "'") // Noncompliant
	if err := row.Scan(&user); err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	fmt.Fprintf(w, "Authentication successful")
}
