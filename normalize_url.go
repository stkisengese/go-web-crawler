package main

import (
	"net/url"
	"strings"
)

func normalizeURL(rawURL string) (string, error) {
	parsed, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}
	host := parsed.Hostname()
	path := strings.TrimSuffix(parsed.EscapedPath(), "/")
	if path == "" {
		return host, nil
	}
	return host + path, nil
}
