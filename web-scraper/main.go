package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	path := "https://webscraper.io/test-sites/e-commerce/allinone"
	resp, err := http.Get(path)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	links := parseResponse(resp)

	fmt.Printf("links: %v\n", links)
}

func parseResponse(resp *http.Response) (links []string) {
	linkSet := make(map[string]struct{})
	tokenizer := html.NewTokenizer(resp.Body)

	baseURL, err := url.Parse(resp.Request.URL.String())
	if err != nil {
		log.Fatal(err)
	}

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

		// Check if the href is a relative URL
		if strings.HasPrefix(href, "/") {
			href = baseURL.Scheme + "://" + baseURL.Host + href
		} else if !strings.HasPrefix(href, "http") {
			absoluteURL, err := baseURL.Parse(href)
			if err != nil {
				log.Printf("error parsing href %s: %v", href, err)
				continue
			}
			href = absoluteURL.String()
		}

		linkSet[href] = struct{}{}
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
