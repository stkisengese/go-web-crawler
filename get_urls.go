package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)


func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	// Parse the html body into a string
	doc, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML: %w", err)
	}

	// Parse the base URL link for solving relative URLs
	base, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("invalid base URL: %w, err")
	}
	return traverse(doc, base), nil
}
