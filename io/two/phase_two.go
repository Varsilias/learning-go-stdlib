// Exercise 2: Understanding io.Writer
// The Writer interface has ONE method: Write([]byte) (int, error)
// Memory trick: "Writers WRITE bytes FROM your slice"
// You can write to Standard Output, to a Buffer(you can read too), to File
// this means that stdout, buffer and file implement the io.Writer interface
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

// Task 4. Create your own writer that counts bytes
type CounterWriter struct {
	Count  int
	writer io.Writer
}

func NewCounterWriter() *CounterWriter {
	return &CounterWriter{
		Count:  0,
		writer: bufio.NewWriter(nil),
	}
}

func (cw *CounterWriter) Write(data []byte) (int, error) {
	n, err := cw.writer.Write(data)
	cw.Count += n

	return n, err
}

func main() {
	fmt.Println("=== EXERCISE 2: io.Writer Fundamentals ===")

	// When I convert values (like int, bool, float) to string and then to []byte,
	// the number of bytes seems to equal the number of characters.

	// ChatGPT said
	/*In UTF-8:
	- ASCII characters (which includes digits 0–9, letters a–z, symbols like +, .) are each 1 byte.
	- Non-ASCII characters (like emojis, Chinese, accented letters) may take 2–4 bytes.
	*/

	// write string
	// data := []byte("Hello, Go Writer!")

	// write int
	// data := []byte(strconv.Itoa(25000))

	// write bool
	// data := []byte(fmt.Sprintf("%t", true))

	// write float
	// data := []byte(fmt.Sprintf("%.2f", 9876.54))

	//Write Emoji
	data := []byte("😀") // this write 4 bytes because it is an emoji

	// Writer 1: Write to stdout (terminal)
	fmt.Println("Writing to standard output(stdout): ")
	n, err := os.Stdout.Write(data)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("\nWrote %d bytes to standard output(stdout)\n", n)

	// Writer 2: Write to a buffer (in-memory)
	var buffer bytes.Buffer
	n, err = buffer.Write(data)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("\nWrote %d bytes to buffer: %q\n", n, buffer.String())

	// Writer 3: Write to a file
	file, err := os.Create("test_output.txt")
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
	}
	defer file.Close()
	n, err = file.Write(data)
	if err != nil {
		fmt.Printf("Error writing to file(%s): %v\n", file.Name(), err)
	}
	fmt.Printf("\nWrote %d bytes to file(%s)\n", n, file.Name())

	// Step 4: treating all writers the same!
	writers := []io.Writer{os.Stdout, &buffer, file}
	for i, w := range writers {
		message := fmt.Sprintf("Hello %d: Hello from writer!\n", i+1)
		w.Write([]byte(message))
	}

	cw := NewCounterWriter()

	log.Print(cw)

	cw.Write([]byte("😀"))
	cw.Write([]byte("Hello, Go Writer!"))
	cw.Write([]byte(fmt.Sprintf("%.2f", 9876.54)))
	cw.Write([]byte(strconv.Itoa(25000)))

	fmt.Printf("CountWriter Has %d bytes written into it\n", cw.Count)

}
