package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

func main() {
	fmt.Println("Hello, World!")

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

		fmt.Printf("tokenType: %v\n", tokenType)
		fmt.Printf("token: %v\n", token)
	}
}
