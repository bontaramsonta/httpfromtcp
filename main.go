package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

func getLinesChannel(f io.ReadCloser) <-chan string {
	lines := make(chan string)
	currentLine := ""

	// read all bytes from file: 8 bytes at a time
	go func() {
		for {
			b := make([]byte, 8, 8)
			_, err := f.Read(b)
			if err != nil {
				if errors.Is(err, io.EOF) {
					break
				}
				log.Fatal("Error reading file:", err)
			}

			for _, char := range string(b) {
				if char == '\n' {
					lines <- currentLine
					currentLine = ""
				} else {
					currentLine += string(char)
				}
			}
		}
		close(lines)
	}()

	return lines
}

func main() {
	file, err := os.Open("./messages.txt")
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	lines := getLinesChannel(file)

	for line := range lines {
		fmt.Printf("read: %s\n", line)
	}

}
