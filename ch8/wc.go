package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var charCount, lineCount, wordCount int

	for scanner.Scan() {
		line := scanner.Text()
		// Slightly hacky; add on one character for newline.
		// Would be an issue if e.g. on Windows.
		charCount += len(line) + 1
		lineCount++
		wordCount += len(strings.Fields(line))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("lines:     ", lineCount)
	fmt.Println("characters:", charCount)
	fmt.Println("words:     ", wordCount)
}
