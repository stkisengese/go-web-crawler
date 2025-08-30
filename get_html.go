package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	// Make an HTTP GET request to the URL
	res, err := http.Get(rawURL)
	if err != nil {
		return "", fmt.Errorf("failed to fetch URL: %w", err)
	}
	defer res.Body.Close()

	// Check for non-200 status codes
	if res.StatusCode > 399 {
		return "", fmt.Errorf("request failed with status code: %d %s", res.StatusCode, res.Status)
	}

	// Check that the Content-Type is text/html
	contentType := res.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		return "", fmt.Errorf("expected content type 'text/html', got: %s", contentType)
	}

	// Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	//Return the body as a string
	return string(body), nil
}
