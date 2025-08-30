package main

import (
	"net/url"
	"sync"
)

// config struct holds shared state for concurrent crawling
type config struct {
	pages              map[string]int  // Track pages and visit counts
	baseURL            *url.URL        // Original base URL for domain checking
	mu                 *sync.Mutex     // Mutex for thread-safe map access
	concurrencyControl chan struct{}   // Buffered channel to control max goroutines
	wg                 *sync.WaitGroup // WaitGroup to wait for all goroutines to finish
}

// AddPageVisit safely adds or updates the visit count for a normalized
// URL and returns whether it was the first visit.
func (cfg *config) addPageVisit(normURL string) (isFirstVisit bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	if _, exists := cfg.pages[normURL]; exists {
		cfg.pages[normURL]++
		return false // Not the first visit
	}

	cfg.pages[normURL] = 1
	return true // First visit
}
