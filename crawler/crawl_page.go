package crawler

import (
	"fmt"
	"net/url"
)

// CrawlPage recursively crawls a website starting from the raw currentURL
func (cfg *Config) CrawlPage(rawCurrentURL string) {
	// Acquire a slot in the concurrency control channel
	cfg.ConcurrencyControl <- struct{}{}
	defer func() {
		<-cfg.ConcurrencyControl // Release the slot when done
		cfg.Wg.Done()            // Mark this goroutine as done when it exits
	}()

	// Check if we've reached the max number of pages
	if cfg.GetPageCount() >= cfg.MaxPages {
		return
	}

	// Check URL is on the same domain as the base URL
	if currentURL, err := url.Parse(rawCurrentURL); err != nil {
		fmt.Printf("error parsing current URL %s: %v\n", currentURL, err)
		return
	} else if cfg.BaseURL.Host != currentURL.Host {
		return
	}

	// Get normalize version of the current URL
	normURL, err := NormalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("error normalizing URL %s: %v\n", rawCurrentURL, err)
		return
	}

	// Check if we've already visited this page
	isFirstVisit := cfg.AddPageVisit(normURL)
	if !isFirstVisit {
		return // Already visited, no need to crawl again
	}
	fmt.Printf("Crawling: %s\n", rawCurrentURL)

	// Fetch the HTML of the page
	htmlBody, err := GetHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error getting HTML from %s: %v\n", rawCurrentURL, err)
		return
	}

	// Extract all the links from the HTML
	urls, err := GetURLsFromHTML(htmlBody, cfg.BaseURL.String())
	if err != nil {
		fmt.Printf("Error extracting URLS from %s: %v\n", rawCurrentURL, err)
		return
	}

	// Recursively crawl each of the links
	for _, foundURL := range urls {
		// Check maxPages reached before spawning new goroutines
		if cfg.GetPageCount() >= cfg.MaxPages {
			break
		}
		cfg.Wg.Add(1)
		go cfg.CrawlPage(foundURL)
	}
}