// Exercise 4: Building Custom io Types
// Learn to create your own Readers and Writers
// Memory trick: "If it Reads, it's a Reader. If it Writes, it's a Writer"

package main

import (
	"fmt"
	"io"
	"strings"
)

// CountingWriter counts bytes as they're written
type CountingWriter struct {
	BytesWritten int
	Destination  io.Writer
}

// Write implements io.Writer interface
func (cw *CountingWriter) Write(data []byte) (int, error) {
	n, err := cw.Destination.Write(data)
	cw.BytesWritten += n
	return n, err
}

// UppercaseReader converts everything to uppercase as it's read
type UpperCaseReader struct {
	Source io.Reader
}

// Read implements io.Reader interface
func (ur *UpperCaseReader) Read(buffer []byte) (int, error) {
	n, err := ur.Source.Read(buffer)
	if n > 0 {
		//Converting to Uppercase: If the character is a lowercase letter, the line buffer[i] = buffer[i] - 'a' + 'A' converts it to uppercase:
		// buffer[i] - 'a' gives the zero-based index of the letter (e.g., for 'a', it gives 0; for 'b', it gives 1, etc.).
		//Adding 'A' (which has an ASCII value of 65) converts this index back to the corresponding uppercase letter.
		for i := range n {
			if buffer[i] >= 'a' && buffer[i] <= 'z' {
				buffer[i] = buffer[i] - 'a' + 'A'
			}
		}
	}

	return n, err
}

// PrefixWriter adds a prefix to each write
type PrefixWriter struct {
	Prefix      string
	Destination io.Writer
	needPrefix  bool
}

func NewPrefixWriter(prefix string, dest io.Writer) *PrefixWriter {
	return &PrefixWriter{
		Prefix:      prefix,
		Destination: dest,
		needPrefix:  true,
	}
}

// Write implements io.Writer interface
func (pw *PrefixWriter) Write(data []byte) (int, error) {
	var totalWritten int

	if pw.needPrefix {
		prefixByte := []byte(pw.Prefix)
		n, err := pw.Destination.Write(prefixByte)
		if err != nil {
			return n, err
		}
		pw.needPrefix = false
	}

	n, err := pw.Destination.Write(data)
	totalWritten += n

	// check whether the next data to be written needs a prefix
	// if needs a prefix if it has a new line character
	if len(data) > 0 && data[len(data)-1] == '\n' {
		pw.needPrefix = true
	}

	return totalWritten, err
}

func main() {
	fmt.Println("=== EXERCISE 4: Custom io Types ===")

	// Test 1: CountingWriter
	fmt.Println("1. Testing CountingWriter:")
	counter := &CountingWriter{Destination: &strings.Builder{}}
	counter.Write([]byte("Hello "))
	counter.Write([]byte("World!"))

	fmt.Printf("Total Bytes Written: %d\n", counter.BytesWritten)

	// Test 2: UppercaseReader
	fmt.Println("\n2. Testing UppercaseReader:")
	originalText := "hello, this should be uppercase!"
	source := strings.NewReader(originalText)

	upperCaseReader := &UpperCaseReader{Source: source}
	result, err := io.ReadAll(upperCaseReader)
	if err != nil {
		fmt.Printf("Error reading from UpperCaseReader: %v\n", err)
	}

	fmt.Printf("Original: %s\n", originalText)
	fmt.Printf("Uppercase: %s\n", string(result))

	// Test 3: PrefixWriter
	fmt.Println("\n3. Testing PrefixWriter:")
	var output strings.Builder
	prefixWriter := NewPrefixWriter("[LOG]", &output)

	prefixWriter.Write([]byte("First Message\n"))
	prefixWriter.Write([]byte("Second Message\n"))
	prefixWriter.Write([]byte("Third Message"))

	fmt.Printf("Output \n%s\n", output.String())

	// Test 4: Combine them all!
	fmt.Println("\n4. Combining custom types:")

	// Create a pipeline: UppercaseReader -> PrefixWriter -> CountingWriter
	text := "hello world\nthis is amazing\n"

	reader := &UpperCaseReader{Source: strings.NewReader(text)}
	var finalOutput strings.Builder
	countingWriter := CountingWriter{Destination: &finalOutput}
	prefixWriter = NewPrefixWriter("> ", &countingWriter)

	io.Copy(prefixWriter, reader)

	fmt.Printf("Pipeline result:\n%s", finalOutput.String())
	fmt.Printf("Total bytes in final output: %d\n", countingWriter.BytesWritten)

}
