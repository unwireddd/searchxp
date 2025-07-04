package main

import (
	/*"fmt"
	"net/http"
	"strings"
	"time"
	"math/rand"
	"net/url"
	"github.com/PuerkitoBio/goquery"*/
)

bingDomains = map[string]string{
	"com":""
}

type SearchResult struct{
	ResultRank int
	ResultURL string
	ResultTitle string
	ResultDesc string
}

var userAgents = []string{
	"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_8_9; en-US) AppleWebKit/600.23 (KHTML, like Gecko) Chrome/49.0.3283.150 Safari/534",
	"Mozilla/5.0 (Windows; Windows NT 6.0;; en-US) AppleWebKit/535.42 (KHTML, like Gecko) Chrome/47.0.2764.275 Safari/603",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 11_3_6; like Mac OS X) AppleWebKit/533.47 (KHTML, like Gecko)  Chrome/47.0.3302.201 Mobile Safari/602.7",
	"Mozilla/5.0 (Windows; U; Windows NT 6.1; WOW64) Gecko/20100101 Firefox/62.2",
	"Mozilla/5.0 (Windows NT 10.4;; en-US) AppleWebKit/601.3 (KHTML, like Gecko) Chrome/47.0.2810.350 Safari/534.7 Edge/11.94885",
	"Mozilla/5.0 (Linux; Linux i686 x86_64; en-US) Gecko/20100101 Firefox/64.7",
}

func randomUserAgent() string{

}

func buildBingUrls(){

}

func scrapeClientRequest(){

}
func BingScrape(searchTerm, country string, pages, count)([]SearchResult, error){
	results := []SearchResult{}
	bingPages, err := buildBingUrls(searchTerm, country, pages, count)	
}

func bingResultParser(){

}