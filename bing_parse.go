package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/kluctl/go-embed-python/python"
	//"net/http"
	//"strings"
)

var resp_bing string

func test_python_bindings() {
	// so this actually works but I need to find out on how to make it handle requests
	// note that 15 mins more tomorrow

	fname_bing := strings.ReplaceAll(fname, " ", "+")
	python_code := fmt.Sprintf(`
	import requests
	from bs4 import BeautifulSoup

	l=[]
	o={}
	for i in range(0,100,10):

    target_url=f"https://www.bing.com/search?q=%s&rdr=1".format(i+1)

    print(target_url)

    resp=requests.get(target_url)

    soup = BeautifulSoup(resp.text, 'html.parser')

    completeData = soup.find_all("li",{"class":"b_algo"})

    for i in range(0, len(completeData)):
         o["Title"]=completeData[i].find("a").text
         o["link"]=completeData[i].find("a").get("href")
         o["Description"]=completeData[i].find("div",
       {"class":"b_caption"}).text
         o["Position"]=i+1
         l.append(o)
         o={}

print(l)`, fname_bing)

	// ok so this seems to be working and it just prints out the output of my python code
	ep, err := python.NewEmbeddedPython("example")
	if err != nil {
		panic(err)
	}

	cmd, err := ep.PythonCmd("-c", python_code)
	if err != nil {
		panic(err)
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()

	if err != nil {
		panic(err)
	}
	/*

		fname_bing := strings.ReplaceAll(fname, " ", "+")

		for i := 0; i < 100; i += 10 {
			target_url := fmt.Sprintf("https://www.bing.com/search?q=%s&rdr=1", fname_bing)
			req, err := http.Get(target_url)
			if err != nil {
				fmt.Println(err)
			}
			for {
				fmt.Println()
			}

			/*resp, err := soup.Get(target_url)
			if err != nil {
				os.Exit(1)
			}
			resp_bing = soup.HTMLParse(resp).HTML()

			fmt.Println(fname_bing)
			fmt.Println(target_url)
			fmt.Println(resp_bing)
		}
	*/

}

// if something goes really bad turned out that I can just embed python scripts inside of my golang code using https://github.com/kluctl/go-embed-python
// so now we can tell that the fname package can be read here which lets us proceed to rewriting the parser in go
// note that I may also want to either change the IDE or solve the extension problem somehow
// so my tood for tomorrow is to fix the broken vscode extension and finish rewriting that python script
// broken extension got fixed so I can proceed to writing parsers
