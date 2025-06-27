package main

import (
	"os"
	"fmt"
	"strings"
	"github.com/anaskhan96/soup"
)

var resp_bing string

func bing() {
	fname_bing := strings.ReplaceAll(fname, " ", "+")
	target_url := fmt.Sprintf("https://www.bing.com/search?q=%s&rdr=1", fname_bing)
	// target_url=f"https://www.bing.com/search?q={query2}&rdr=1".format(i+1)
	resp, err := soup.Get(target_url)
	if err != nil {
		os.Exit(1)
	}
	resp_bing = soup.HTMLParse(resp).HTML()

		fmt.Println(fname_bing)
		fmt.Println(target_url)
		fmt.Println(resp_bing)

}

 // if something goes really bad turned out that I can just embed python scripts inside of my golang code using https://github.com/kluctl/go-embed-python
 // so now we can tell that the fname package can be read here which lets us proceed to rewriting the parser in go
 // note that I may also want to either change the IDE or solve the extension problem somehow
 // so my tood for tomorrow is to fix the broken vscode extension and finish rewriting that python script