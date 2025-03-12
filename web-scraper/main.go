package main

func main() {
	rootPath := "https://webscraper.io/test-sites/e-commerce/allinone"
	scraper, _ := NewScraper(rootPath, 5)

	scraper.Scrape(rootPath)
}
