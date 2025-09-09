package crawler

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

// ExportToCSV function exports the crawl result to a CSV file
func ExportToCSV(pages map[string]int, baseURL, filename string) error {
	// Create or truncate the file
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create CSV file: %v", err)
	}
	defer file.Close()

	// Initialize CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV header
	header := []string{"URL", "Link_Depth", "Domain", "Base_URL", "Crawl_Timestamp"}
	if err := writer.Write(header); err != nil {
		return fmt.Errorf("error writing header to CSV: %v", err)
	}

	// sort pages for consistent output
	sortedPages := SortPages(pages)
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	// Write each page's data
	for _, page := range sortedPages {
		record := []string{
			page.URL,
			strconv.Itoa(page.Count),
			ExtractDomain(page.URL),
			baseURL,
			timestamp,
		}
		if err := writer.Write(record); err != nil {
			return fmt.Errorf("error writing record to CSV: %v", err)
		}
	}
	return nil
}

// ExportDetailedCSV exports a detailed analysis of the crawled pages to a CSV file
func ExportDetailedCSV(pages map[string]int, baseURL, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create detailed CSV file: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Enhanced header with more analytical columns
	header := []string{
		"URL",
		"Internal_Link_Count",
		"Relative_Popularity",
		"URL_Depth",
		"Page_Type",
		"Domain",
		"Base_URL",
		"Crawl_Timestamp",
	}
	if err := writer.Write(header); err != nil {
		return fmt.Errorf("failed to write detailed CSV header: %w", err)
	}

	sortedPages := SortPages(pages)
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	// Calculate max links for relative popularity
	maxLinks := 0
	for _, page := range sortedPages {
		if page.Count > maxLinks {
			maxLinks = page.Count
		}
	}

	// Write enhanced data rows
	for _, page := range sortedPages {
		relativePopularity := "0%"
		if maxLinks > 0 {
			percentage := float64(page.Count) / float64(maxLinks) * 100
			relativePopularity = fmt.Sprintf("%.1f%%", percentage)
		}

		record := []string{
			page.URL,
			strconv.Itoa(page.Count),
			relativePopularity,
			strconv.Itoa(CalculateURLDepth(page.URL)),
			DeterminePageType(page.URL),
			ExtractDomain(page.URL),
			baseURL,
			timestamp,
		}
		if err := writer.Write(record); err != nil {
			return fmt.Errorf("failed to write detailed CSV record: %w", err)
		}
	}

	return nil
}