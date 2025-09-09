package crawler

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

// GetURLsFromHTML extracts and returns all URLs from the provided HTML body.
func GetURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	// Parse the html body into a string
	doc, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML: %w", err)
	}

	// Parse the base URL link for solving relative URLs
	base, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("invalid base URL: %w", err)
	}

	// Traverse the HTML node tree and extract URLs
	return traverse(doc, base), nil
}

// Helper function to traverse the HTML node tree
func traverse(node *html.Node, base *url.URL) []string {
	var urls []string

	// Traverse the HTML node tree to find anchor tags
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, attr := range node.Attr {
			if attr.Key == "href" {

				// Parse the href value
				link, err := url.Parse(attr.Val)
				if err == nil {

					// Resolve relative URLs
					absLink := base.ResolveReference(link)
					urls = append(urls, absLink.String())
				}
			}
		}
	}

	// Recursively traverse child nodes
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		childURLs := traverse(child, base)
		urls = append(urls, childURLs...)
	}
	return urls
}