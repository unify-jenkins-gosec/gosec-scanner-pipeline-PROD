package main

import (
	"fmt"
	"github.com/antchfx/xmlquery"
	"net/http"
)

func Authenticate(request *http.Request, doc *xmlquery.Node) bool {
	user := request.FormValue("user")
	password := request.FormValue("password")

	query := fmt.Sprintf("//User[username=\"%s\" and password=\"%s\"]", user, password) // Noncompliant
	userObj, _ := xmlquery.Query(doc, query)

	return userObj != nil
}
