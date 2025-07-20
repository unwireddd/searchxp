package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// the country detection thing is pretty pointless since I havent really proceeded to writing the actual server side stuff so lets just leave it poland for now

var bingDomains = map[string]string{
	"com": "",
	"pl":  "&cc=PL",
}
var bingRes []SearchResult

type SearchResult struct {
	ResultRank  int
	ResultURL   string
	ResultTitle string
	ResultDesc  string
}

// so I have that result struct here which I can convert to html type strings using the string conversion
var resultsArr []string
var descriptions []string

var userAgents = []string{
	"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_8_9; en-US) AppleWebKit/600.23 (KHTML, like Gecko) Chrome/49.0.3283.150 Safari/534",
	"Mozilla/5.0 (Windows; Windows NT 6.0;; en-US) AppleWebKit/535.42 (KHTML, like Gecko) Chrome/47.0.2764.275 Safari/603",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 11_3_6; like Mac OS X) AppleWebKit/533.47 (KHTML, like Gecko)  Chrome/47.0.3302.201 Mobile Safari/602.7",
	"Mozilla/5.0 (Windows; U; Windows NT 6.1; WOW64) Gecko/20100101 Firefox/62.2",
	"Mozilla/5.0 (Windows NT 10.4;; en-US) AppleWebKit/601.3 (KHTML, like Gecko) Chrome/47.0.2810.350 Safari/534.7 Edge/11.94885",
	"Mozilla/5.0 (Linux; Linux i686 x86_64; en-US) Gecko/20100101 Firefox/64.7",
}

func randomUserAgent() string {
	rand.Seed(time.Now().Unix())
	randNum := rand.Int() % len(userAgents)
	return userAgents[randNum]
}

func buildBingUrls(searchTerm, country string, pages, count int) ([]string, error) {
	toScrape := []string{}
	searchTerm = strings.Trim(searchTerm, " ")
	searchTerm = strings.Replace(searchTerm, " ", "+", -1)
	// here it should enter the country
	if countryCode, found := bingDomains[country]; found {
		for i := 0; i < pages; i++ {
			first := firstParameter(i, count)
			scrapeURL := fmt.Sprintf("https://bing.com/search?q=%sfirst=%d&count=%d%s", searchTerm, first, count, countryCode)
			toScrape = append(toScrape, scrapeURL)
		}
	} else {
		fmt.Println("Country not found")

	}
	return toScrape, nil
}

func firstParameter(number, count int) int {
	if number == 0 {
		return number + 1
	}
	return number*count + 1
}

func getScrapeClient(proxyString interface{}) *http.Client {
	switch V := proxyString.(type) {
	case string:
		//proxyUrl, _ := url.Parse(V)
		//return &http.Client(Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)})
		fmt.Println(V)
	default:
		return &http.Client{}
	}
	return nil
}

func scrapeClientRequest(searchURL string, proxyString interface{}) (*http.Response, error) {
	baseClient := getScrapeClient(proxyString)
	req, err := http.NewRequest("GET", searchURL, nil)
	req.Header.Set("User-Agent", randomUserAgent())
	res, err := baseClient.Do(req)
	if res.StatusCode != 200 {
		err, _ := fmt.Println("BANNED")
		fmt.Println(err)
	}
	if err != nil {
		fmt.Println(err)
	}
	return res, nil
}

func BingScrape(searchTerm, country string, pages, count, backoff int) ([]SearchResult, error) {
	results := []SearchResult{}
	bingPages, err := buildBingUrls(searchTerm, country, pages, count)
	if err != nil {
		fmt.Println(err)
	}

	for _, page := range bingPages {
		rank := len(results)
		res, err := scrapeClientRequest(page, nil)
		if err != nil {
			fmt.Println(err)
		}
		data, err := bingResultParser(res, rank)
		if err != nil {
			fmt.Println(err)
		}
		for _, result := range data {
			results = append(results, result)
		}
		time.Sleep(time.Duration(backoff) * time.Second)
	}
	return results, nil
}

func bingResultParser(response *http.Response, rank int) ([]SearchResult, error) {
	doc, err := goquery.NewDocumentFromResponse(response)
	if err != nil {
		fmt.Println(err)
	}
	results := []SearchResult{}
	sel := doc.Find("li.b_algo")
	rank++

	for i := range sel.Nodes {
		item := sel.Eq(i)
		linkTag := item.Find("a")
		link, _ := linkTag.Attr("href")
		titleTag := item.Find("h2")
		descTag := item.Find("div.b_caption p")
		desc := descTag.Text()
		title := titleTag.Text()
		link = strings.Trim(link, " ")
		if link != "" && link != "#" && !strings.HasPrefix(link, "/") {
			result := SearchResult{
				rank,
				link,
				title,
				desc,
			}
			resultArr := fmt.Sprintf(`<a href="%s">%s</a>`, link, title)
			resultsArr = append(resultsArr, resultArr)
			descriptions = append(descriptions, fmt.Sprintf("<p>%s</p>", desc))
			results = append(results, result)
		}
	}
	return results, nil
}

func mainParser(query string) []string {
	res, err := BingScrape(query, "pl", 2, 30, 30)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println("A")
	fmt.Println(res)
	bingRes = res
	procGen()
	return resultsArr
	// ok so I have my parsed arrays with html strings and now what I want to do is to inject them into my html file somehow
	// I can just write a html file and then make a function that puts all of them in it

}
