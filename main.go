package main

import (
	"bytes"
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

		// split bytes into fields based on newline character
		fields := bytes.FieldsFunc(b, func(r rune) bool {
			return r == '\n'
		})

		switch len(fields) {
		case 0:
			continue
		case 1:
			currentLine += string(fields[0])
		default:
			// handle multiple line breaks
			for i := 0; i < len(fields)-1; i++ {
				currentLine += string(fields[i])
				lines = append(lines, currentLine)
				currentLine = ""
			}
			currentLine += string(fields[len(fields)-1])
		}
	}

	//! debug
	for _, line := range lines {
		fmt.Printf("read: %s\n", line)
	}
}
