package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"regexp"
)

func main() {
	fmt.Println("Hello, World!")

	path := "https://webscraper.io/test-sites/e-commerce/allinone"
	resp, err := http.Get(path)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	urlRegex := regexp.MustCompile(
		`(?mi)https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{2,256}\.[a-z]{2,4}\b([-a-zA-Z0-9@:%_\+.~#?&//=]*)`,
	)

	fmt.Printf("body: %v\n", string(body))

	urlsFound := urlRegex.FindAllString(string(body), -1)
	urlSet := make(map[string]struct{})
	for _, url := range urlsFound {
		urlSet[url] = struct{}{}
	}

	for url := range urlSet {
		fmt.Println(url)
	}

	for urlS := range urlSet {
		url, _ := url.Parse(urlS)

		fmt.Printf("url.Hostname(): %v\n", url.Hostname())
	}
}
