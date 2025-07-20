package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
	"math/rand"
	"net/url"
	"github.com/PuerkitoBio/goquery"
)

var googleDomains = map[string]string{
	"com": "https://www.google.com/search?q=",
}

type SearchResultG struct{
	ResultRank int
	ResultURL string
	ResultTitle string
	ResultDesc string
}

var userAgentsG = []string{

}

func randomUserAgentG() string{
	rand.Seed(time.Now().Unix())
	randNumG := rand.Int() % len(userAgentsG)
	return userAgentsG[randNumG]
}

func buildGoogleUrls(searchTerm, countryCode string, languageCode string, pages, count int)([]string, error){
	toScrape := []string{}
	searchTerm = strings.Trim(searchTerm, " ")
	searchTerm = string.Replace(searchTerm, " ", "+", -1)
	if googleBase, found := googleDomains[countryCode]; found{
		for i := 0, i<pages ; i += 1{
			start := i*count
			scrapeURL := fmt.Sprintf("%s%s&num=%d&hl=%s&start=%d&filter=0",googleBase, searchTerm, count, languageCode, start)
		}
	}else {
		err := fmt.Errorf("%s not supported", countryCode)
		return nil, err
	}
	return toScrape, nil
}

func GoogleScrape(searchTerm, countryCode, languageCode string, pages int)([]SearchResultG, err){
	results := []SearchResultG{}
	resultCounter := 0
	googlePages, err := buildGoogleUrls(searchTerm, countryCode, languageCode, pages, count)
	if err != nil {
		return nil, err
	}
}

func google_main() {
	res, err := GoogleScrape("metro", "com", "en", 1, 30)
	if err == nil{
		for _, res := range res{
			fmt.Println(res)
		}
	}
}