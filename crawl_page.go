package main

import (
	"fmt"
	"net/url"
)

// CrawlPage recursively crawls a website starting from the raw currentURL
// rawBaseURL: the root domain we're crawling (stays constant)
// rawCurrentURL: the current page we're crawling (changes with each recursive call)
// pages: map to track how many times we've seen each normalized URL
func (cfg *config) crawlPage(rawCurrentURL string) {
	// Acquire a slot in the concurrency control channel
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl // Release the slot when done
		cfg.wg.Done()            // Mark this goroutine as done when it exits
	}()

	// Check if we've reached the max number of pages
	if cfg.getPageCount() >= cfg.maxPages {
		return
	}

	// Check URL is on the same domain as the base URL
	if currentURL, err := url.Parse(rawCurrentURL); err != nil {
		fmt.Printf("error parsing current URL %s: %v\n", currentURL, err)
		return
	} else if cfg.baseURL.Host != currentURL.Host {
		return
	}

	// Get normalize version of the current URL
	normURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("error normalizing URL %s: %v\n", rawCurrentURL, err)
		return
	}

	// Check if we've already visited this page
	isFirstVisit := cfg.addPageVisit(normURL)
	if !isFirstVisit {
		return // Already visited, no need to crawl again
	}
	fmt.Printf("Crawling: %s\n", rawCurrentURL)

	// Fetch the HTML of the page
	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error getting HTML from %s: %v\n", rawCurrentURL, err)
		return
	}

	// Extract all the links from the HTML
	urls, err := getURLsFromHTML(htmlBody, cfg.baseURL.String())
	if err != nil {
		fmt.Printf("Error extracting URLS from %s: %v\n", rawCurrentURL, err)
		return
	}

	// Recursively crawl each of the links
	for _, foundURL := range urls {
		// Check maxPages reached before spawning new goroutines
		if cfg.getPageCount() >= cfg.maxPages {
			break
		}
		cfg.wg.Add(1)
		go cfg.crawlPage(foundURL)
	}
}
