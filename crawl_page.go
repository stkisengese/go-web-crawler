package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	// Check URL is on the same domain as the base URL
	if baseURL, err := url.Parse(rawBaseURL); err != nil {
		fmt.Printf("error parsing base URL %s: %v\n", baseURL, err)
		return
	} else if currentURL, err := url.Parse(rawCurrentURL); err != nil {
		fmt.Printf("error parsing current URL %s: %v\n", currentURL, err)
		return
	} else if baseURL.Host != currentURL.Host {
		return
	}

	// Check if we've already visited this page
	// Mark this page as visited
	// Fetch the HTML of the page
	// Extract all the links from the HTML
	// Recursively crawl each of the links
}
