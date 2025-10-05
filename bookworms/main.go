package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	bookworms, err := loadBookworms(".\\testdata\\bookworms.json")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error loading bookworms. %v", err)
		os.Exit(1)
	}

	for _, bookworm := range bookworms {
		fmt.Println(bookworm)
		fmt.Println(strings.Repeat("=", 50))
	}

}
