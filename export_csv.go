package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

// ExportToCSV function exports the crawl result to a CSV file
func exportToCSV(pages map[string]int, baseURL, filename string) error {
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
	sortedPages := sortPages(pages)
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	// Write each page's data
	for _, page := range sortedPages {
		record := []string{
			page.URL,
			strconv.Itoa(page.Count),
			extractDomain(page.URL),
			baseURL,
			timestamp,
		}
		if err := writer.Write(record); err != nil {
			return fmt.Errorf("error writing record to CSV: %v", err)
		}
	}
	return nil
}
