package main

import (
	"fmt"
	"io"
	"net/http"
)

func getHTML(rawURL string) (string, error) {
	// Make an HTTP GET request to the URL
	res, err := http.Get(rawURL)
	if err != nil {
		return "", fmt.Errorf("failed to fetch URL: %w", err)
	}
	defer res.Body.Close()

	// Check for non-200 status codes
	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("request faild with status code: %d %s", res.StatusCode, res.Status)
	}

	// Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	//Return the body as a string
	return string(body), nil
}
