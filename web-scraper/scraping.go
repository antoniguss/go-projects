package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/html"
)

// TODO: Make it work
type Scraper struct {
	visited     map[string]struct{}
	mutex       sync.RWMutex
	wg          sync.WaitGroup
	urlQueue    chan string
	results     []string
	rootDomain  string
	maxRoutines int
	semaphore   chan struct{}
}

func NewScraper(rootURL string, maxConcurrency int) (*Scraper, error) {
	parsedURL, err := url.Parse(rootURL)
	if err != nil {
		return nil, err
	}

	return &Scraper{
		visited:     make(map[string]struct{}),
		urlQueue:    make(chan string, 1000),
		results:     []string{},
		rootDomain:  parsedURL.Hostname(),
		maxRoutines: maxConcurrency,
		semaphore:   make(chan struct{}, maxConcurrency),
	}, nil
}

func (s *Scraper) worker() {
	defer s.wg.Done()

	for pageURL := range s.urlQueue {
		links, err := scrapePath(pageURL)
		if err != nil {
			fmt.Printf("Error scraping %s: %v\n", pageURL, err)
			continue
		}

		// Process found links
		for _, link := range links {
			s.mutex.RLock()
			_, visited := s.visited[link]
			s.mutex.RUnlock()

			if !visited {
				s.mutex.Lock()
				// Check again after acquiring write lock to avoid race condition
				if _, has := s.visited[link]; !has {
					s.visited[link] = struct{}{}
					s.results = append(s.results, link)
					s.mutex.Unlock()

					// Queue this new URL
					s.wg.Add(1)
					s.urlQueue <- link
				} else {
					s.mutex.Unlock()
				}
			}
		}

		<-s.semaphore // Release semaphore slot
	}
}

func (s *Scraper) Scrape(rootURL string) ([]string, error) {
	s.mutex.Lock()
	s.visited[rootURL] = struct{}{}
	s.results = append(s.results, rootURL)
	s.mutex.Unlock()

	for i := 0; i < s.maxRoutines; i++ {
		s.wg.Add(1)
		go s.worker()
	}

	s.semaphore <- struct{}{}
	s.urlQueue <- rootURL

	s.wg.Wait()
	close(s.urlQueue)

	return s.results, nil
}

func processUrl(url string, urlsChan chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Processing:", url)

	rateLimiter := time.Tick(100 * time.Millisecond)
	<-rateLimiter

	links, err := scrapePath(url)
	if err != nil {
		fmt.Println("Error scraping URL:", url, err)
		return
	}

	for _, link := range links {
		urlsChan <- link
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
