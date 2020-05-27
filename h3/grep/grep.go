package main

import (
	"os"
	"log"
	"fmt"
	"bufio"
	"flag"
	"strings"
)

func search() {
	var filename string
	var word string
	var caseInsensitive string
	flag.StringVar(&filename, "f", "test.txt", "a file to search for words")
	flag.StringVar(&word, "w", "foo", "a search word")
	flag.StringVar(&caseInsensitive, "c", "n", "case-insensitive search - y for yes, n for no")
	flag.Parse()

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var isFound bool
		if caseInsensitive == "y" {
			isFound = strings.Contains(strings.ToLower(line), strings.ToLower(word))
		} else {
			isFound = strings.Contains(line, word)
		}	
		if isFound {
			fmt.Println(line)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	search()
}
