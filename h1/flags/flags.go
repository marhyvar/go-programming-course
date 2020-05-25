package main

import (
	"fmt"
	"flag"
)

// This program takes a word as a user input (flag) and prints how many letters the word has
func main() {
	var word string
	flag.StringVar(&word, "word", "foo", "any word you like")
	flag.Parse()
	fmt.Printf("The word has %d letters.\n", len(word))
}