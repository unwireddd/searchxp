package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// Function to get a list of image files in a directory
func getImageFiles(directory string) ([]string, error) {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		return nil, err
	}

	var imageFiles []string
	for _, file := range files {
		if !file.IsDir() {
			filePath := filepath.Join(directory, file.Name())
			imageFiles = append(imageFiles, filePath)
		}
	}

	return imageFiles, nil
}

// Function to generate HTML content
func generateHTML(imageFiles []string) string {
	htmlContent := `
    <!DOCTYPE html>
    <html>
    <head>
        <title>Image Gallery</title>
    </head>
    <body>
    `
	for _, file := range imageFiles {
		htmlContent += fmt.Sprintf(`<img src="%s" alt="Image">`, filepath.Base(file)) + "\n"
	}
	htmlContent += `
    </body>
    </html>
    `
	return htmlContent
}

func htmlgen() {
	if len(os.Args) != 3 {
		log.Fatal("Usage: ", os.Args[0], " <directory> <output_file>")
	}

	directory := os.Args[1]
	outputFile := os.Args[2]

	imageFiles, err := getImageFiles(directory)
	if err != nil {
		log.Fatal(err)
	}

	htmlContent := generateHTML(imageFiles)

	err = ioutil.WriteFile(outputFile, []byte(htmlContent), 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("HTML file generated successfully.")
}
