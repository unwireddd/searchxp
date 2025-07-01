package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/anaskhan96/soup"
	//"github.com/anaskhan96/soup"
)

var l []string
var o = make(map[string]string)

func get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func rewritten(x string) {
	for i := 0; i < 100; i += 10 {
		target_url := fmt.Sprintf("https://www.bing.com/search?q=%s&rdr=1", x)
		fmt.Println(target_url)
		resp, err := soup.Get(target_url)
		if err != nil {
			fmt.Println(err)
		}
		parsed := soup.HTMLParse(resp)
		//links := doc.Find("div", "id", "comicLinks").FindAll("a")
		completeData := parsed.FindAll("li", "class", "b_algo")
		for i := 0; i < len(completeData); i++ {
			o["Title"] = completeData[i].Find("a").Text()
			fmt.Println(completeData)
			fmt.Println("start")
			// so the .html actually gives me some sensible output
			// lets try 1,5h tomorrow
			//fmt.Println(parsed.HTML())
			fmt.Println("end")
			fmt.Println(completeData[i].Find("a").Text())
			// ok so the problem is that the line above is empty for some reason
			fmt.Println(o)
		}

	}

}
