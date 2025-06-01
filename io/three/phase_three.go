// Exercise 3: The Power of io.Copy
// io.Copy(dst Writer, src Reader) - The most important function in Go!
// Memory trick: "Copy FROM Reader TO Writer" = Copy(TO, FROM)

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"syscall"
)

// Task 1: Create a function that copies from any Reader to any Writer
func CopyFromAnyReaderToAnyWriter(reader io.Reader, writer io.Writer) {
	len := 64
	buffer := make([]byte, len)
	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Printf("Error: %v\n", err)
			break
		}

		writer.Write(buffer[:n])
	}

}

// Task 3: Build a simple "tee" program that writes to both file and stdout
type Tee struct {
	Writers []io.Writer
}

func NewTee(writers ...io.Writer) *Tee {
	return &Tee{
		Writers: writers,
	}
}

func (t *Tee) Write(data []byte) (int, error) {
	var n int
	var err error
	for i, w := range t.Writers {
		n, err = w.Write(data)
		if err != nil {
			fmt.Printf("Error writing to writer: %d", i)
		}
	}

	return n, err
}

func main() {
	fmt.Println("=== EXERCISE 3: The Magic of io.Copy ===")

	// Step 1: Basic copy from string to buffer
	source := strings.NewReader("Hello, Copy World!")
	var destination bytes.Buffer

	bytesWritten, err := io.Copy(&destination, source)
	if err != nil {
		fmt.Printf("Error Copying data: %v\n", err)
		return
	}

	fmt.Printf("Copied %d bytes: %q\n", bytesWritten, destination.String())

	// Step 2: Copy from file to stdout
	testFile, err := os.Create("source.txt")
	if err != nil {
		fmt.Printf("Error Creating file: %v\n", err)
		return
	}
	// testFile.WriteString("This is file content!\nLine 2\nLine 3")

	// let me do this with Write to show that I understand it
	testFile.Write([]byte("This is file content!\nLine 2\nLine 3"))
	testFile.Close()

	file, err := os.Open("source.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	defer file.Close()
	fmt.Println("\nFile Contents")
	io.Copy(os.Stdout, file)

	// Step 3: Chain multiple operations
	fmt.Println("\n\n=== CHAINING OPERATIONS ===")
	// Create a pipe: Reader -> Processing -> Writer
	input := strings.NewReader("HELLO WORLD")
	var pipe bytes.Buffer
	var out bytes.Buffer

	// Copy to pipe buffer
	io.Copy(&pipe, input)

	// Process the data (convert to lowercase)
	processed := strings.ToLower(pipe.String())
	processedRead := strings.NewReader(processed)

	io.Copy(&out, processedRead)

	fmt.Printf("Original: HELLO WORLD\n")
	fmt.Printf("Processed: %s\n", out.String())

	// os.Remove("source.txt")

	testData := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.\n\n"

	reader := strings.NewReader(testData)

	var writerOutput *bytes.Buffer = &bytes.Buffer{}
	CopyFromAnyReaderToAnyWriter(reader, writerOutput)
	fmt.Println(writerOutput.String())

	// teeFile, err := os.Create("tee.txt")
	// if err != nil {
	// 	fmt.Printf("Error reading file: %v\n", err)
	// 	return
	// }
	// defer teeFile.Close()

	// tee := NewTee(os.Stdout, teeFile)
	// data := "Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae\n\n"

	// tee.Write([]byte(data))

	// Task 2: Try copying from stdin to a file (hint: os.Stdin is a Reader)
	file, err = os.Create("file.txt")
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
	}
	defer file.Close()

	tee := NewTee(os.Stdout, file)

	fmt.Println("Enter any text, Press Ctrl+C to exit.")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.Contains(line, "\n") {
			line = line + "\n"
		}

		if strings.Contains(line, syscall.SIGINT.String()) {
			break
		}
		tee.Write([]byte(line))

		// file.Write([]byte(line))

	}

}
