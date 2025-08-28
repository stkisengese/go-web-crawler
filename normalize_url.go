package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(rawURL string) (string, error) {
	parsed, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("couldn't parse URL: %w", err)
	}
	host := parsed.Hostname()
	path := strings.TrimSuffix(parsed.EscapedPath(), "/")
	if path == "" {
		return strings.ToLower(host), nil
	}
	return strings.ToLower(host + path), nil
}
