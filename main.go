package main

import (
	"fmt"
	"net/http"
	"html/template"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("sites/*.html"))
}

// Initiating the webserver
func main() {
	http.HandleFunc("/", index)
	// if there is a form or something in html then the line below will make it handle it with a certain function we assigned to it
	http.HandleFunc("/process", processor)
	http.ListenAndServe(":6060", nil)
}

// Connecting the html folder to our webserver
func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}

// This one should handle the html form
func processor(w http.ResponseWriter, r *http.Request) {
	fname := r.FormValue("phrase")
	fmt.Println(fname)
	// Saving the variable contents into my text file

	file, err := os.Create("output.txt")
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }
    defer file.Close() // Ensure the file is closed after writing

    // Write the content to the file
    _, err = file.WriteString(fname)
    if err != nil {
        fmt.Println("Error writing to file:", err)
	}
	// yeah that works so now I can display the content of it in my html file using javascript (tomorrow)

	// so the ExecuteTemplate function brings me to the html file I want to go to
	tpl.ExecuteTemplate(w, "processor.html", fname)
}