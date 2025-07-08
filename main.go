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
var temp *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("sites/*.html"))
	//tpl2 = template.Must(template.ParseGlob("*.html"))
}

// Initiating the webserver
func main() {

	http.HandleFunc("/", index)
	// if there is a form or something in html then the line below will make it handle it with a certain function we assigned to it
	// NOTE THAT THE FIX FOR MY PROBLEM WAS TO MAKE SURE THAT output.html file already exists from the beginning
	http.HandleFunc("/output", output)
	http.ListenAndServe(":6060", nil)
}

// Connecting the html folder to our webserver
func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}

// This one should handle the html form
func output(w http.ResponseWriter, r *http.Request) {
	// okay so it turns out that the whole thing only works if the output file is here from the beginning
	// making the script remove all of Its contents and replacing them with the results for a new search term wont really work at all since the function
	// reads the file from Its beginning state anyways
	fname = r.FormValue("phrase")
	fmt.Println(fname)
	mainParser(fname)
	procGen()
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
	//output := mainParser("a")

	//okay so now when the parser works as it should what I want to do is to make Its content readable like in a regular browser

	query := struct {
		First string
	}{
		First: "a",
	}
	fmt.Println(query)
	// so the ExecuteTemplate function brings me to the html file I want to go to

	temp = template.Must(template.ParseGlob("res/*.html"))

	temp.ExecuteTemplate(w, "output.html", nil)

}
