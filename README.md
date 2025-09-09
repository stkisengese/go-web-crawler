# Go Web Crawler

A simple and concurrent web crawler written in Go.

## Description

This project is a web crawler that takes a URL as input and crawls all the pages within the same domain. It is designed to be fast and efficient by using goroutines to crawl pages concurrently.

The crawler will output a report of the crawled pages and the number of internal links pointing to each page. The results can also be exported to a CSV file for further analysis.

## Features

- Concurrent crawling using goroutines
- Control the maximum concurrency
- Limit the number of pages to crawl
- Export crawl results to a CSV file
- Export detailed analysis to a CSV file

## Getting Started

### Prerequisites

- Go 1.21 or higher

### Installation

1. Clone the repository:

```bash
git clone https://github.com/stkisengese/go-web-crawler.git
```

2. Navigate to the project directory:

```bash
cd go-web-crawler
```

3. Build the project:

```bash
go build
```

## Usage

To run the web crawler, use the following command:

```bash
./go-web-crawler <URL> <maxConcurrency> <maxPages> [options]
```

**Arguments:**

- `URL`: The starting URL to crawl (e.g., `https://example.com`)
- `maxConcurrency`: The maximum number of concurrent requests (e.g., `5`)
- `maxPages`: The maximum number of pages to crawl (e.g., `100`)

**Options:**

- `--csv FILE`: Export results to a CSV file.
- `--detailed-csv FILE`: Export detailed analysis to a CSV file.

**Example:**

```bash
./go-web-crawler https://example.com 5 100 --csv results.csv
```

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.