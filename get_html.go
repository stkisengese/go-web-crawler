package main

import (
	"fmt"
	"net/http"
)

func getHTML(rawURL string) (string, error) {
	// Make an HTTP GET request to the URL
	res, err := http.Get(rawURL)
	if err != nil {
		return "", fmt.Errorf("failed to fetch URL: %w", err)
	}
	defer res.Body.Close()
	// Read the response body
	

	// Handle errors and return the response body as a string
	return "", nil
}
