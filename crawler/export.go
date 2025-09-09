package crawler

import "strings"

// ExtractDomain extracts the domain from a given URL string. If the URL is empty, it returns an empty string.
// It iterates through the characters of the URL and returns the substring
// up to the first '/' character. If no '/' is found, it returns the entire URL.
func ExtractDomain(url string) string {
	if len(url) == 0 {
		return ""
	}

	if slashIndex := strings.Index(url, "/"); slashIndex != -1 {
		return url[:slashIndex]
	}

	return url
}

// CalculateURLDepth check the number of nested pages on the domain
func CalculateURLDepth(url string) int {
	var depth = 0
	for _, char := range url {
		if char == '/' {
			depth++
		}
	}
	return depth
}

// DeterminePageType is a helper function to determine page type based on URL patterns
func DeterminePageType(url string) string {
	if url == ExtractDomain(url) {
		return "Homepage"
	}

	// Check for common page types
	pageTypeMappings := map[string]string{
		"/blog":     "Blog",
		"/post":     "Blog",
		"/about":    "About",
		"/contact":  "Contact",
		"/product":  "Product",
		"/item":     "Product",
		"/category": "Category",
		"/cat":      "Category",
		"/news":     "News",
		"/service":  "Service",
		"/help":     "Support",
		"/support":  "Support",
	}

	for path, pageType := range pageTypeMappings {
		if strings.Contains(url, path) {
			return pageType
		}
	}
	return "Content"
}