// Level 2: Real System Programming Challenges
// Now that you understand io fundamentals, build actual system tools

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	file := flag.String("file", "", "Path to the file that needs to be watched")

	// Register short flag
	flag.StringVar(file, "f", "", "Short for --file")

	flag.Parse()
	if *file == "" {
		log.Fatalf("No file path provided, please provide a value for --flag")
		os.Exit(1)
	}

	filePath, err := filepath.Abs(*file)
	if err != nil {
		log.Fatalf("Error: could not determine file path - %v\n", err)
		os.Exit(1)
	}

	fileInfo, err := os.Stat(filePath)
	if err != nil {
		log.Fatalf("An error occured when checking file - %v\n", err)
		os.Exit(1)
	}

	if fileInfo.IsDir() {
		log.Fatalf("%s is a directory \n", fileInfo.Name())
		os.Exit(1)
	}

	size := fileInfo.Size()

	fileReader, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file")
		os.Exit(1)
	}
	defer fileReader.Close()

	_, err = fileReader.Seek(2, io.SeekEnd)
	if err != nil {
		log.Fatalf("Error seeking file")
		os.Exit(1)
	}

	for {
		fileInfo, err := os.Stat(filePath)
		if err != nil {
			log.Fatalf("An error occured when checking file - %v\n", err)
			os.Exit(1)
		}
		if fileInfo.Size() < size {
			fmt.Println("File has been truncated, resetting seek offset")
			fileReader.Seek(0, io.SeekEnd)
		}
		scanner := bufio.NewScanner(fileReader)

		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}

		if err := scanner.Err(); err != nil && err != io.EOF {
			log.Fatalf("Error reading file content: %v\n", err)
		}

		time.Sleep(500 * time.Millisecond)

	}

}
