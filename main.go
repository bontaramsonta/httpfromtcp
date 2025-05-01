package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./messages.txt")
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	var lines []string
	currentLine := ""

	// read all bytes from file: 8 bytes at a time
	for {
		b := make([]byte, 8, 8)
		_, err := file.Read(b)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal("Error reading file:", err)
		}

		for _, char := range b {
			if char == '\n' {
				lines = append(lines, currentLine)
				currentLine = ""
			} else {
				currentLine += string(char)
			}
		}
	}

	//! debug
	for _, line := range lines {
		fmt.Printf("read: %s\n", line)
	}
}
