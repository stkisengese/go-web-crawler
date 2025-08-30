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

	// Get normalize version of the current URL
	normURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("error normalizing URL %s: %v\n", rawCurrentURL, err)
		return
	}

	// Check if we've already visited this page
	if count, exists := pages[normURL]; exists {
		pages[normURL] = count + 1
		return
	}

	// Mark this page as visited
	pages[normURL] = 1
	fmt.Printf("Crawling: %s\n", rawCurrentURL)

	// Fetch the HTML of the page
	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error getting HTML from %s: %v\n", rawCurrentURL, err)
		return
	}

	// Extract all the links from the HTML
	urls, err := getURLsFromHTML(htmlBody, rawBaseURL)
	if err != nil {
		fmt.Printf("Error extracting URLS from %s: %v\n", rawCurrentURL, err)
	}

	// Recursively crawl each of the links
}
