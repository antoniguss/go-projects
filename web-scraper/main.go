package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/html"
)

func main() {
	fmt.Println("Starting")

	rootPath := "https://webscraper.io/test-sites/e-commerce/allinone"
	urls := make(chan string)
	results := make(chan []string)

	numWorkers := 5
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for range numWorkers {
		go worker(urls, results, &wg)
	}

	urls <- rootPath
	close(urls)

	wg.Wait()
	close(results)

	for result := range results {
		for _, url := range result {
			urls <- url
		}
	}

	fmt.Println("Scraping completed")
}

func worker(urls <-chan string, results chan<- []string, wg *sync.WaitGroup) {
	defer wg.Done()
	rateLimiter := time.Tick(100 * time.Millisecond)

	for url := range urls {
		<-rateLimiter
		links, err := scrapePath(url)
		if err != nil {
			fmt.Println("TODO: Handle error in worker")
		}
		results <- links
	}
}

func scrapePath(rootPath string) (links []string, err error) {
	var rootDomain string
	rootDomain, err = getDomain(rootPath)
	if err != nil {
		return
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(rootPath)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	foundLinks := parseResponse(resp)

	for _, link := range foundLinks {
		domain, err := getDomain(link)
		if err != nil {
			return nil, err
		}

		if domain == rootDomain {
			links = append(links, link)
		}
	}

	return links, nil
}

// Calculates a^b
func power(a, b int, ch chan<- int, wg *sync.WaitGroup) int {
	defer wg.Done()
	result := int(math.Pow(float64(a), float64(b)))

	ch <- result
	return result
}

// Returns the domain from the url, returns an error if the url can't be parsed
// e.g. https://en.wikipedia.org/wiki/Go_(programming_language) returns wikipedia.org
func getDomain(path string) (domain string, err error) {
	var url *url.URL
	url, err = url.Parse(path)
	if err != nil {
		return
	}

	hostname := strings.Split(url.Hostname(), ".")

	domain = hostname[len(hostname)-2] + "." + hostname[len(hostname)-1]
	return domain, nil
}

// Returns a list of absolute links found in the given HTTP response, without duplicates
func parseResponse(resp *http.Response) (links []string) {
	linkSet := make(map[string]struct{})
	tokenizer := html.NewTokenizer(resp.Body)

	baseURL := resp.Request.URL

	for {
		tokenType := tokenizer.Next()
		token := tokenizer.Token()
		if tokenType == html.ErrorToken {
			break
		}

		if token.Data != "a" {
			continue
		}

		ok, href := getHref(token)
		if !ok {
			continue
		}

		hrefURL, err := url.Parse(href)
		if err != nil {
			log.Fatal(err)
		}

		var absoluteURL *url.URL
		if hrefURL.IsAbs() {
			absoluteURL = hrefURL
		} else {
			absoluteURL = baseURL.ResolveReference(hrefURL)
		}

		if absoluteURL.Scheme == "http" || absoluteURL.Scheme == "https" {
			linkSet[absoluteURL.String()] = struct{}{}
		} else {
			log.Println("Skipping non-http(s) link:", href)
		}

	}

	links = make([]string, 0, len(linkSet))
	for k := range linkSet {
		links = append(links, k)
	}

	return links
}

func getHref(token html.Token) (ok bool, href string) {
	for _, attr := range token.Attr {
		if attr.Key == "href" {
			return true, attr.Val
		}
	}
	return
}
