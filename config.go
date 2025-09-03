package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
)

// config struct holds shared state for concurrent crawling
type config struct {
	pages              map[string]int  // Track pages and visit counts
	baseURL            *url.URL        // Original base URL for domain checking
	mu                 *sync.Mutex     // Mutex for thread-safe map access
	concurrencyControl chan struct{}   // Buffered channel to control max goroutines
	wg                 *sync.WaitGroup // WaitGroup to wait for all goroutines to finish
	maxPages           int             // Optional limit on number of pages to crawl
}

// CrawlArgs struct holds arguments for each crawl operation
type CrawlArgs struct {
	URL            string // URL to crawl
	MaxConcurrency int    // Maximum concurrency for this crawl
	MaxPages       int    // Maximum pages for this crawl
	CSVFile        string // CSV file to save results
	DetailedCSV    string // Detailed CSV file to save results
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

// GetPageCount safety returns the current number of pages crawled
func (cfg *config) getPageCount() int {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	return len(cfg.pages)
}

// printUsage prints usage information
func printUsage() {
	fmt.Println("Usage: ./crawler <URL> <maxConcurrency> <maxPages> [options]")
	fmt.Println()
	fmt.Println("Required arguments:")
	fmt.Println("  URL              Website to crawl (include http:// or https://)")
	fmt.Println("  maxConcurrency   Maximum number of concurrent requests (1-20)")
	fmt.Println("  maxPages         Maximum number of pages to crawl")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  --csv FILE              Export results to CSV file")
	fmt.Println("  --detailed-csv FILE     Export detailed analysis to CSV file")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  ./crawler https://example.com 5 20")
	fmt.Println("  ./crawler https://blog.dev 3 10 --csv results.csv")
	fmt.Println("  ./crawler https://site.com 8 50 --detailed-csv analysis.csv")
	fmt.Println("  go run . https://example.com 5 20 --csv output.csv --detailed-csv detailed.csv")
}

// parseArgs parses command line arguments with support for CSV export
func parseArgs() CrawlArgs {
	args := os.Args[1:]

	if len(args) < 3 {
		printUsage()
		os.Exit(1)
	}

	// Parse required arguments
	url := args[0]
	maxConcurrency, err := strconv.Atoi(args[1])
	if err != nil || maxConcurrency < 1 {
		fmt.Printf("Invalid maxConcurrency '%s': must be a positive integer\n", args[1])
		os.Exit(1)
	}

	maxPages, err := strconv.Atoi(args[2])
	if err != nil || maxPages < 1 {
		fmt.Printf("Invalid maxPages '%s': must be a positive integer\n", args[2])
		os.Exit(1)
	}

	result := CrawlArgs{
		URL:            url,
		MaxConcurrency: maxConcurrency,
		MaxPages:       maxPages,
	}

	// parse optional flags
	for i := 3; i < len(args); i++ {
		arg := args[i]
		switch {
		case strings.HasPrefix(arg, "--csv="):
			result.CSVFile = strings.TrimPrefix(arg, "--csv=")
		case strings.HasPrefix(arg, "--detailed-csv="):
			result.DetailedCSV = strings.TrimPrefix(arg, "--detailed-csv=")
		case arg == "--csv" && i+1 < len(args):
			result.CSVFile = args[i+1]
			i++
		case arg == "--detailed-csv" && i+1 < len(args):
			result.DetailedCSV = args[i+1]
			i++
		default:
			fmt.Printf("Unknown argument: %s\n", arg)
			printUsage()
			os.Exit(1)
		}
	}
	return result
}
