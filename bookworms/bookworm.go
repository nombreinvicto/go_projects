package main

import "os"

// Bookworm Go struct must match the JSON
// Fields need to be exposed as the decoder
// we are about to write need to write to the fields
type Bookworm struct {
	Name  string `json:"name"`
	Books []Book `json:"books"`
}

type Book struct {
	Author string `json:"author"`
	Title  string `json:"title"`
}

// loadBookworms reads the file and returns the list of bookworms
// and their beloved books, found therein
func loadBookworms(filePath string) ([]Bookworm, error) {

	// if success, return *os.File and nill for error
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// os.File implements io.Reader and io.Closer interfaces
	// remember io.Reader enables reading a stream of data into a slice of bytes

	return nil, nil
}
