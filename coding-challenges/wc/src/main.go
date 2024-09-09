package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	fileBytesCount := flag.String("c", "", "count bytes in a file, specify file path")
	fileLinesCount := flag.String("l", "", "count lines in a file, specify file path")
	fileWordsCount := flag.String("w", "", "count words in a file, specify file path")
	fileCharsCount := flag.String("m", "", "count chars in a file, specify file path")
	flag.Parse()

	if *fileBytesCount != "" {
		fileContent, err := os.ReadFile(*fileBytesCount)
		if err != nil {
			log.Fatal(err)
		}
		countBytes(fileContent)
		return
	}
	if *fileLinesCount != "" {
		fileContent, err := os.ReadFile(*fileLinesCount)
		if err != nil {
			log.Fatal(err)
		}
		countLines(fileContent)
		return
	}
	if *fileWordsCount != "" {
		fileContent, err := os.ReadFile(*fileWordsCount)
		if err != nil {
			log.Fatal(err)
		}
		countWords(fileContent)
		return
	}
	if *fileCharsCount != "" {
		fileContent, err := os.ReadFile(*fileCharsCount)
		if err != nil {
			log.Fatal(err)
		}
		countChars(fileContent)
		return
	}
	// Run all the flags if only filename was provided
	if flag.NFlag() == 0 && flag.NArg() == 1 {
		fileContent, err := os.ReadFile(flag.Arg(0))
		if err != nil {
			log.Fatal(err)
		}
		countBytes(fileContent)
		countLines(fileContent)
		countWords(fileContent)
		return
	}
	// Read stdin if no input was provided
	if flag.NFlag() == 0 && flag.NArg() == 0 {
		stdin, err := io.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		countBytes(stdin)
		countLines(stdin)
		countWords(stdin)
		return
	}

}

func countBytes(content []byte) {
	fmt.Println(len(content))
}

func countLines(content []byte) {
	contentStr := string(content)
	lines := strings.Split(contentStr, "\n")
	// Count the number of lines
	lineCount := len(lines)

	if contentStr != "" && contentStr[len(contentStr)-1] == '\n' {
		lineCount--
	}
	fmt.Println(lineCount)
}

func countWords(content []byte) {
	contentStr := string(content)
	words := strings.Fields(contentStr)
	// Count the number of words
	wordCount := len(words)
	fmt.Println(wordCount)
}

func countChars(content []byte) {
	contentStr := string(content)
	charsCount := utf8.RuneCountInString(contentStr)
	fmt.Println(charsCount)
}
