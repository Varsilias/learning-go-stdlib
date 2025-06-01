// Exercise 1: Understanding io.Reader
// The Reader interface has ONE method: Read([]byte) (int, error)
// Memory trick: "Readers READ bytes INTO a slice"

package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	fmt.Println("=== EXERCISE 1: io.Reader Fundamentals ===")

	// text := "Hello, Go Reader!"
	text := "The sleeping dog jumps over the lazy sleeping lion"
	reader := strings.NewReader(text)

	// Reading data the hard way
	buffer := make([]byte, 3)

	fmt.Println("Reading in 5-byte chunks:")

	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			fmt.Println("Reached end of data (EOF)")
			break
		}

		if err != nil {
			fmt.Printf("Error: %v/n", err)
		}

		fmt.Printf("Read %d bytes: %q\n", n, string(buffer[:n]))
	}

	// Reset Reader and read all
	reader = strings.NewReader(text)
	allData, err := io.ReadAll(reader)
	if err != nil {
		fmt.Printf("Error reading all: %v\n", err)
	}

	fmt.Printf("\nRead all at once: %q\n", allData)
}
