package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	//"github.com/anaskhan96/soup"
)

// maybe it would be good to just search for a more sensible script that does that

var l []string
var o = make(map[string]string)
var bing_html string

func get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func rewritten(x string) {
	fmt.Println("A")

}
