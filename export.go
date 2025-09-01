package main

import "strings"

// extractDomain extracts the domain from a given URL string. If the URL is empty, it returns an empty string.
// It iterates through the characters of the URL and returns the substring
// up to the first '/' character. If no '/' is found, it returns the entire URL.
func extractDomain(url string) string {
	if len(url) == 0 {
		return ""
	}

	if slashIndex := strings.Index(url, "/"); slashIndex != -1 {
		return url[:slashIndex]
	}

	return url
}

// CalculateURLDepth check the number of nested pages on the domain
func calculateURLDepth(url string) int {
	var depth = 0
	for _, char := range url {
		if char == '/' {
			depth++
		}
	}
	return depth
}
