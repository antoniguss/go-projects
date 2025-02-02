package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	fmt.Println("Hello, World!")

	links := make(map[string]struct{})

	path := "https://webscraper.io/test-sites/e-commerce/allinone"
	resp, err := http.Get(path)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	// urlRegex := regexp.MustCompile(
	// 	`(?mi)https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{2,256}\.[a-z]{2,4}\b([-a-zA-Z0-9@:%_\+.~#?&//=]*)`,
	// )

	// fmt.Printf("body: %v\n", string(body))

	tokenizer := html.NewTokenizer(resp.Body)
	for {
		tokenType := tokenizer.Next()
		token := tokenizer.Token()
		if tokenType == html.ErrorToken {
			break
		}

		// fmt.Printf("tokenType: %v\n", tokenType)
		// fmt.Printf("token: %v\n", token)
		if token.Data != "a" {
			continue
		}

		ok, href := getHref(token)

		if !ok {
			continue
		}

		if strings.HasPrefix(href, "/") {
			href = path + href
		}

		links[href] = struct{}{}

	}

	for link := range links {
		fmt.Printf("link: %v\n", link)
	}
}

func getHref(token html.Token) (ok bool, href string) {
	for _, attr := range token.Attr {
		if attr.Key == "href" {
			return true, attr.Val
		}
	}

	return
}
