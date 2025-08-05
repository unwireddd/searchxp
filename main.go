package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
)

// reminder that if I want to use the same variable in different files I need to declare it before using it in any fundtion so Its global

// okay do what I need to do now is to make the output file clean itself before adding new data, adding a goBack function for now and make an image scraper

var fname string
var whichPage int

//var whichEngine string

var tpl *template.Template
var temp *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("sites/*.html"))
	//tpl2 = template.Must(template.ParseGlob("*.html"))
}

// Initiating the webserver
func main() {

	defer os.RemoveAll("/home/metro/searchxp/dataset")
	http.HandleFunc("/", index)
	// if there is a form or something in html then the line below will make it handle it with a certain function we assigned to it
	// NOTE THAT THE FIX FOR MY PROBLEM WAS TO MAKE SURE THAT output.html file already exists from the beginning
	http.HandleFunc("/output", output)
	http.HandleFunc("/displayImages", displayImages)
	http.HandleFunc("/goback", goback)
	http.HandleFunc("/spageNext", spageNext)
	http.HandleFunc("/yandexNext", spageNext)
	http.HandleFunc("/metaNext", spageNext)
	http.ListenAndServe(":6060", nil)
}

// Connecting the html folder to our webserver
func index(w http.ResponseWriter, r *http.Request) {
	whichPage = 1
	tpl.ExecuteTemplate(w, "index.html", nil)
}

// This one should handle the html form
func output(w http.ResponseWriter, r *http.Request) {

	// okay so it turns out that the whole thing only works if the output file is here from the beginning
	// making the script remove all of Its contents and replacing them with the results for a new search term wont really work at all since the function
	// reads the file from Its beginning state anyways
	whichPage = 1
	fname = r.FormValue("phrase")
	fmt.Println(fname)
	whichEngine := r.FormValue("Engine")
	/*
		isYandex := r.FormValue("Yandex")
		isStartpage := r.FormValue("Startpage")
		isMetasearch := r.FormValue("Metasearch")
		whichEngine := ""
		switch {
		case isYandex == "Yandex":
			whichEngine = "Yandex"
		case isStartpage == "Startpage":
			whichEngine = "Startpage"
		case isMetasearch == "Metasearch":
			whichEngine = "Metasearch"
		default:
			whichEngine = "Err"
		}*/

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

	// TEST USING INIT PARSER
	/*scriptPathInit := "/home/metro/searchxp/helium/parsers_init.py"
	cmdInit := exec.Command("python3", scriptPathInit)

	errSpage := cmdInit.Run()
	if errSpage != nil {
		fmt.Println(errSpage)
		return
	}*/

	// executing the python startpage scraper script

	//startpage

	switch {
	case whichEngine == "Metasearch":
		// TEST USING INIT PARSER
		scriptPathInit := "/home/metro/searchxp/helium/parsers_init.py"
		cmdInit := exec.Command("python3", scriptPathInit)

		errSpage := cmdInit.Run()
		if errSpage != nil {
			fmt.Println(errSpage)
			return
		}

	case whichEngine == "Yandex":
		scriptPathYandex := "/home/metro/searchxp/helium/yandex_standalone.py"
		cmdY := exec.Command("python3", scriptPathYandex)
		errYandex := cmdY.Run()
		if errYandex != nil {
			fmt.Println(errYandex)
			return
		}
	case whichEngine == "Startpage":
		scriptPathSpage := "/home/metro/searchxp/helium/startpage.py"

		// Create a new command to run the Python script
		cmdS := exec.Command("python3", scriptPathSpage)

		errSpage := cmdS.Run()
		if errSpage != nil {
			fmt.Println(errSpage)
			return
		}
	default:
		scriptPathSpage := "/home/metro/searchxp/helium/startpage.py"

		// Create a new command to run the Python script
		cmdS := exec.Command("python3", scriptPathSpage)

		errSpage := cmdS.Run()
		if errSpage != nil {
			fmt.Println(errSpage)
			return
		}
		// yandex
		scriptPathYandex := "/home/metro/searchxp/helium/yandex.py"
		cmdY := exec.Command("python3", scriptPathYandex)
		errYandex := cmdY.Run()
		if errYandex != nil {
			fmt.Println(errYandex)
			return
		}

		//images
		scriptPathImages := "/home/metro/searchxp/helium/images_yandex.py"
		cmdY1 := exec.Command("python3", scriptPathImages)
		errImages := cmdY1.Run()
		if errImages != nil {
			fmt.Println(errImages)
			return
		}
	}

	/*
		// Create a new command to run the Python script
			cmdS := exec.Command("python3", scriptPathSpage)

			errSpage := cmdS.Run()
			if errSpage != nil {
				fmt.Println(errSpage)
				return
			}


		scriptPathSpage := "/home/metro/searchxp/helium/startpage.py"

		// Create a new command to run the Python script
		cmdS := exec.Command("python3", scriptPathSpage)

		errSpage := cmdS.Run()
		if errSpage != nil {
			fmt.Println(errSpage)
			return
		}
		// yandex
		scriptPathYandex := "/home/metro/searchxp/helium/yandex.py"
		cmdY := exec.Command("python3", scriptPathYandex)
		errYandex := cmdY.Run()
		if errYandex != nil {
			fmt.Println(errYandex)
			return
		}

		//images
		scriptPathImages := "/home/metro/searchxp/helium/images_yandex.py"
		cmdY1 := exec.Command("python3", scriptPathImages)
		errImages := cmdY1.Run()
		if errImages != nil {
			fmt.Println(errImages)
			return
		}*/

	// AAA

	// Capture the output of the script
	//var out bytes.Buffer
	//cmd.Stdout = &out

	// Run the command and wait for it to finish

	// end

	//mainParser(fname)
	//procGen()

	// EXECUTING THE IMAGE BULK-DOWNLOAD SCRIPT START

	/*
		pythonCmd := "python3"
		scriptPath := "/home/metro/searchxp/scripts/imgscrape_test.py"
		scriptPath2 := "/home/metro/searchxp/htmlgen.py"

		// Create a new command
		cmd := exec.Command(pythonCmd, scriptPath)

		// Run the command and wait for it to complete
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Error executing script: %s\nOutput: %s\n", err, output)
			return
		}

		fmt.Println("Script executed successfully")
		cmd2 := exec.Command(pythonCmd, scriptPath2)
		time.Sleep(10)
		// Run the command and wait for it to complete
		output2, err := cmd2.CombinedOutput()
		if err != nil {
			fmt.Printf("Error executing script: %s\nOutput: %s\n", err, output2)
			return
		}

		fmt.Println("Script executed successfully")
		// image bulk-downloading works fine
		// so now what I need to do is to make a script that displays all of those downloaded images in a html format

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

		// test executing the python webserver
		dir3 := "/home/metro/searchxp/dataset"
		cmd3 := "python3 -m http.server 8000"

		// Change into the directory
		c := exec.Command("sh", "-c", cmd3)
		c.Dir = dir3

		// Start the command in the background
		err3 := c.Start()
		if err3 != nil {
			log.Fatal(err)
		}

		// Your script continues to run here
		log.Println("Command started in background")

		// end test
	*/

	temp = template.Must(template.ParseGlob("helium/*.html"))

	temp.ExecuteTemplate(w, "res_spage.html", nil)

}

func displayImages(w http.ResponseWriter, r *http.Request) {
	scriptPathInit := "/home/metro/searchxp/helium/images_yandex.py"
	cmdInit := exec.Command("python3", scriptPathInit)

	errSpage := cmdInit.Run()
	if errSpage != nil {
		fmt.Println(errSpage)
		return
	}
	temp = template.Must(template.ParseGlob("helium/*.html"))
	temp.ExecuteTemplate(w, "res_images.html", nil)
}

func spageNext(w http.ResponseWriter, r *http.Request) {
	whichPage += 1
	// WRITING THE PAGE VAR TO FILE

	// Convert int to string
	myIntStr := fmt.Sprintf("%d", whichPage)

	// Specify the filename
	filename := "whichPage.txt"

	// Open the file for writing
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Failed to create file: %v\n", err)
		return
	}
	defer file.Close()

	// Write the string to the file
	_, err = file.WriteString(myIntStr) // Added newline for clarity
	if err != nil {
		fmt.Printf("Failed to write to file: %v\n", err)
		return
	}

	// END

	// for how the problem is that this function doesnt even execute on button
	scriptPathInit := "/home/metro/searchxp/helium/startpage_next.py"
	cmdInit := exec.Command("python3", scriptPathInit)
	fmt.Println("AA")

	errSpage := cmdInit.Run()
	if errSpage != nil {
		fmt.Println(errSpage)
		return
	}
	temp = template.Must(template.ParseGlob("helium/*.html"))
	temp.ExecuteTemplate(w, "res_spage.html", nil)
}

func yandexNext(w http.ResponseWriter, r *http.Request) {
	whichPage += 1
	// WRITING THE PAGE VAR TO FILE

	// Convert int to string
	myIntStr := fmt.Sprintf("%d", whichPage)

	// Specify the filename
	filename := "whichPage.txt"

	// Open the file for writing
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Failed to create file: %v\n", err)
		return
	}
	defer file.Close()

	// Write the string to the file
	_, err = file.WriteString(myIntStr) // Added newline for clarity
	if err != nil {
		fmt.Printf("Failed to write to file: %v\n", err)
		return
	}

	// END

	// for how the problem is that this function doesnt even execute on button
	scriptPathInit := "/home/metro/searchxp/helium/yandex_next.py"
	cmdInit := exec.Command("python3", scriptPathInit)
	fmt.Println("AA")

	errSpage := cmdInit.Run()
	if errSpage != nil {
		fmt.Println(errSpage)
		return
	}
	temp = template.Must(template.ParseGlob("helium/*.html"))
	temp.ExecuteTemplate(w, "res_spage.html", nil)
}

func metaNext(w http.ResponseWriter, r *http.Request) {
	whichPage += 1
	// WRITING THE PAGE VAR TO FILE

	// Convert int to string
	myIntStr := fmt.Sprintf("%d", whichPage)

	// Specify the filename
	filename := "whichPage.txt"

	// Open the file for writing
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Failed to create file: %v\n", err)
		return
	}
	defer file.Close()

	// Write the string to the file
	_, err = file.WriteString(myIntStr) // Added newline for clarity
	if err != nil {
		fmt.Printf("Failed to write to file: %v\n", err)
		return
	}

	// END

	// for how the problem is that this function doesnt even execute on button
	scriptPathInit := "/home/metro/searchxp/helium/parsers_init_next.py"
	cmdInit := exec.Command("python3", scriptPathInit)
	fmt.Println("AA")

	errSpage := cmdInit.Run()
	if errSpage != nil {
		fmt.Println(errSpage)
		return
	}
	temp = template.Must(template.ParseGlob("helium/*.html"))
	temp.ExecuteTemplate(w, "res_spage.html", nil)
}

func goback(w http.ResponseWriter, r *http.Request) {
	// ok so basically the only thing this function does is redirecting me to that file
	// ok so right now it doesnt even execute so the problem is rather in my backend
	// I think that may be caused by the fact that the button is not in my index file
	// but the function itself works since I can execute it from the search bar
	fmt.Println("point")
	files, err := ioutil.ReadDir("./")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, file := range files {
		fmt.Println(file.Name())
		// Handle the file
	}
	temp = template.Must(template.ParseGlob("dataset/*.html"))

	temp.ExecuteTemplate(w, "imgdata.html", nil)
}

// so now the image loading script seems to work just fine but the problem is that it just cant find those images before I reload the script or something
// checking this out rn
