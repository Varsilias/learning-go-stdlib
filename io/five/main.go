// Exercise 5: Advanced io Patterns for System Programming
// Real-world usage patterns  constantly seen in system tools

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

// Pattern 1: Buffered I/O for performance
func demonstrateBufferedIO() {
	fmt.Println("=== PATTERN 1: Buffered I/O ===")

	// testData := strings.Repeat("This is a line of text\n", 1000)

	// Unbuffered writing (slow)
	start := time.Now()
	var unbuffered strings.Builder
	for i := 0; i < 100; i++ {
		unbuffered.WriteString("Line " + fmt.Sprintf("%d", i) + "\n")
	}

	unbufferedTime := time.Since(start)

	bufferedStart := time.Now()
	var bufferedOutput strings.Builder
	bufferdWriter := bufio.NewWriter(&bufferedOutput)
	for i := 0; i < 100; i++ {
		bufferdWriter.WriteString("Line " + fmt.Sprintf("%d", i) + "\n")
	}
	bufferdWriter.Flush()
	bufferedTime := time.Since(bufferedStart)

	fmt.Printf("Unbuffered time: %v\n", unbufferedTime)
	fmt.Printf("Buffered time: %v\n", bufferedTime)
	fmt.Printf("Speedup: %.2fx\n", float64(unbufferedTime)/float64(bufferedTime))

	fmt.Println("🎯 LESSON: Always use bufio for frequent small reads/writes")
}

// Pattern 2: Reading line by line (common in system tools)
func demonstrateLineReading() {
	fmt.Println("\n=== PATTERN 2: Line-by-Line Reading ===")

	// Simulate a log file
	logData := `2024-01-01 10:00:01 INFO Starting application
	2024-01-01 10:00:02 DEBUG Loading configuration
	2024-01-01 10:00:03 WARN Configuration file not found, using defaults
	2024-01-01 10:00:04 INFO Application started successfully
	2024-01-01 10:00:05 ERROR Failed to connect to database`

	reader := strings.NewReader(logData)
	scanner := bufio.NewScanner(reader)

	fmt.Println("Processing log file:")

	lineNumber := 1
	errorCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("Line %d: %s\n", lineNumber, line)

		// Count errors (typical system tool operation)
		if strings.Contains(line, "ERROR") {
			errorCount++
		}
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading: %v\n", err)
	}

	fmt.Printf("Total errors found: %d\n", errorCount)
	fmt.Println("🎯 LESSON: bufio.Scanner is perfect for line-by-line processing")
}

// Pattern 3: Streaming processing (don't load everything in memory)
func demonstrateStreaming() {
	fmt.Println("\n=== PATTERN 3: Streaming Processing ===")

	largeData := strings.Repeat("This is streaming data that we process chunk by chunk", 100)
	reader := strings.NewReader(largeData)

	// Process in chunks without loading all into memory
	buffer := make([]byte, 64)
	totalBytes := 0
	chunkCount := 0

	fmt.Println("Streaming processing:")
	for {
		n, err := reader.Read(buffer)
		if n > 0 {
			chunkCount++
			totalBytes += n

			// Process the chunk (example: count spaces)
			spaces := 0
			for i := 0; i < n; i++ {
				if buffer[i] == ' ' {
					spaces++
				}
			}

			fmt.Printf("Chunk %d: %d bytes, %d spaces\n", chunkCount, n, spaces)
		}

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}

	fmt.Printf("Total: %d bytes in %d chunks\n", totalBytes, chunkCount)
	fmt.Println("🎯 LESSON: Stream processing handles files larger than memory")
}

// Pattern 4: Multi-destination writing (logging to file + console)
type MultiWriter struct {
	writers []io.Writer
}

func NewMultiWriter(writers ...io.Writer) *MultiWriter {
	return &MultiWriter{
		writers: writers,
	}
}

func (mw *MultiWriter) Write(data []byte) (int, error) {
	for _, writer := range mw.writers {
		n, err := writer.Write(data)
		if err != nil {
			return n, err
		}
	}

	return len(data), nil
}

func demonstrateMultiWriter() {
	fmt.Println("\n=== PATTERN 4: Multi-destination Writing ===")

	file, err := os.Create("app.log")
	if err != nil {
		fmt.Printf("Error creating log file: %v\n", err)
		return
	}

	defer file.Close()

	// new multi-writer that writes to both console and file
	multiWriter := NewMultiWriter(os.Stdout, file)

	// Now anything written to multiWriter goes to both destinations
	fmt.Fprint(multiWriter, "This message goes to both console and file\n")
	fmt.Fprint(multiWriter, "Perfect for logging systems!\n")

	fmt.Println("🎯 LESSON: MultiWriter pattern essential for logging systems")

	// os.Remove("app.log")
}

// Pattern 5: Pipeline processing
func demonstratePipeline() {
	fmt.Println("\n=== PATTERN 5: Pipeline Processing ===")

	// we create a processing pipeline: Input -> Uppercase -> Add timestamps -> Output
	input := "hello world\nthis is a test\nfinal line"

	//Stage 1: Input
	stage1 := strings.NewReader(input)

	// Stage 2: Convert to uppercase
	var stage2Buffer strings.Builder
	scanner := bufio.NewScanner(stage1)
	for scanner.Scan() {
		line := strings.ToUpper(scanner.Text())
		stage2Buffer.WriteString(line + "\n")
	}

	// Stage 3: Add timestamps
	stage3Reader := strings.NewReader(stage2Buffer.String())
	var finalOutput strings.Builder

	scanner = bufio.NewScanner(stage3Reader)
	for scanner.Scan() {
		timestamp := time.Now().Format("15:04:05")
		finalOutput.WriteString(fmt.Sprintf("[%s] %s\n", timestamp, scanner.Text()))
	}

	fmt.Println("Pipeline result:")
	fmt.Println(finalOutput.String())

	fmt.Println("🎯 LESSON: Pipelines process data in stages, very memory efficient")

}

func main() {
	// demonstrateBufferedIO()
	// demonstrateLineReading()
	// demonstrateStreaming()
	// demonstrateMultiWriter()
	demonstratePipeline()

	// fmt.Println("\n" + strings.Repeat("=", 50))
	// fmt.Println("🏆 MASTERY CHECKLIST:")
	// fmt.Println("✅ io.Reader: Read([]byte) (int, error)")
	// fmt.Println("✅ io.Writer: Write([]byte) (int, error)")
	// fmt.Println("✅ io.Copy: Universal data transfer")
	// fmt.Println("✅ Buffered I/O: Performance optimization")
	// fmt.Println("✅ Streaming: Memory-efficient processing")
	// fmt.Println("✅ Pipelines: Composable data processing")
	// fmt.Println("\n🎯 NEXT STEPS:")
	// fmt.Println("1. Practice with os.File (implements both Reader and Writer)")
	// fmt.Println("2. Explore net.Conn (network connections are Reader/Writer)")
	// fmt.Println("3. Study compress/gzip (wraps Readers/Writers)")
	// fmt.Println("4. Build a log analyzer using these patterns")
}
