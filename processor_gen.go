package main

import (
	"fmt"
	"os"
)

func procGen() {
	// so everything does clean itself but only if I restart the program first for some reason even though the cleanup is in the function

	// Define the file path and the text to be written before the arrays
	filePath := "res/output.html"
	text := "<!doctype html>\n<html>\n<head>\n<title>res</title>\n</head>\n<body>\n<h1>Search results for: {{.First}}</h1>\n"
	goBack := `<a href="http://localhost:6060/goback">[Go back to search]</a>\n`
	//gotoImg := `<a href="/images.html>[Display Images]</a>`
	fileclean, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	fileclean.Close()

	// Open the file in append mode
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Write the initial text
	_, err = file.WriteString(text)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = file.WriteString(goBack)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Find the maximum length of the two arrays
	maxLen := max(len(resultsArr), len(descriptions))

	// Write the array elements
	for i := 0; i < maxLen; i++ {
		if i < len(resultsArr) {
			_, err = file.WriteString(resultsArr[i] + "\n")
			if err != nil {
				fmt.Println(err)
				return
			}
		}
		if i < len(descriptions) {
			_, err = file.WriteString(descriptions[i] + "\n")
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
