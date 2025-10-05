package main

import (
	"fmt"
	"os"
)

func main() {
	bookworms, err := loadBookworms(".\\testdata\\bookworms.json")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error loading bookworms. %v", err)
		os.Exit(1)
	}

	// find the common books of bookworms
	commonBooks := findCommonBooks(bookworms)
	fmt.Println(commonBooks)
}
