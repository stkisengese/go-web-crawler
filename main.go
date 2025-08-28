package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(os.Args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	BASE_URL := os.Args[1]
	fmt.Println("starting crawl of:", BASE_URL)

	if html, err := getHTML(BASE_URL); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	} else {
		fmt.Println(html)
	}
}
