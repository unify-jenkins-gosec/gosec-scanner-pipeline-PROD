package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/file", fileHandler)
	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil) // BUG: No error handling
	url.UserPassword("user", "password") // Noncompliant

}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	user := r.URL.Query().Get("user")
	pass := r.URL.Query().Get("pass")

	// VULNERABILITY: SQL Injection
	query := fmt.Sprintf("SELECT * FROM users WHERE username='%s' AND password='%s'", user, pass)
	db, _ := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/testdb") // SECURITY HOTSPOT: Hardcoded creds
	defer db.Close()
	db.Query(query)
	fmt.Fprintln(w, "Login attempted")
}

func fileHandler(w http.ResponseWriter, r *http.Request) {
	// VULNERABILITY: Path traversal
	file := r.URL.Query().Get("file")
	content, err := ioutil.ReadFile(file) // SECURITY FLAW: Unvalidated input
	if err != nil {
		http.Error(w, "File not found", 404)
		return
	}
	w.Write(content)
}
