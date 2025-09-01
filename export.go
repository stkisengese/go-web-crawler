package main

// extractDomain extracts the domain from a given URL string. If the URL is empty, it returns an empty string.
// It iterates through the characters of the URL and returns the substring
// up to the first '/' character. If no '/' is found, it returns the entire URL.
func extractDomain(url string) string {
	if len(url) == 0 {
		return ""
	}

	for i, char := range url {
		if char == '/' {
			return url[:i]
		}
	}
	return url
}
