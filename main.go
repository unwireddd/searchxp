package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

// reminder that if I want to use the same variable in different files I need to declare it before using it in any fundtion so Its global

var fname string

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
	fname = r.FormValue("phrase")
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
	// Variable saved

	// ok so this one passes the variable from my previous form to the html file (seems to be working for now)
	// so now what I think I need to do is to parse the search results from some kind of another search engine, I think I may just try to scrape
	// as much content as possible from various different engines through all the methods I can find and then connect all of them somehow, so basically writing a couple of different parsers for next days
	for {
		rewritten("germany")
	}
	query := struct {
		First string
	}{
		First: resp_bing,
	}
	// so the ExecuteTemplate function brings me to the html file I want to go to

	tpl.ExecuteTemplate(w, "processor.html", query)

}
