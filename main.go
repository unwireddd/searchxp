package main

import (
	"net/http"
	"html/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("sites/*.gohtml"))
}

// Initiating the webserver
func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":6060", nil)
}

// Connecting the html folder to our webserver
func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}