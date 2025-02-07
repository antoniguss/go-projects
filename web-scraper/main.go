package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

func scrapePath(rootPath string) {
	rootDomain, err := getDomain(rootPath)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Get(rootPath)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	links := parseResponse(resp)

	for _, link := range links {
		domain, err := getDomain(link)
		if err != nil {
			log.Fatal(err)
		}

		if domain != rootDomain {
			log.Println("skipping")
		}
		fmt.Printf("link: %v\n", link)
	}
}

func main() {
	ch := make(chan int)

	var wg sync.WaitGroup

	for i := range 100000 {
		wg.Add(1)
		go power(2, i, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		fmt.Println(result)
	}
	// rootPath := "https://www.wikipedia.org/"
	// scrapePath(rootPath)
}

// Calculates a^b
func power(a, b int, ch chan<- int, wg *sync.WaitGroup) int {
	defer wg.Done()
	result := int(math.Pow(float64(a), float64(b)))

	ch <- result
	return result
}

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
