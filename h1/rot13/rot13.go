package main

import (
	"fmt"
	"flag"
)

// function that replaces a unicode a-zA-Z character by 13 places
func transform(x rune) rune {
	if x >= 97 && x <= 122 {
		//lowercase a - z, m = 109 (halfway)
		if x > 109 { 
			return x - 13
		} else {
			return x + 13
		}
	} else if x >= 65 && x <= 90 {
		// uppercase A - Z, M = 77 (halfway)
		if x > 77 {
			return x - 13
		} else {
			return x + 13
		}
	}
	return x
}

// function that prints out a ROT-13 ciphered text for text given by user
func main() {
	var word string
	flag.StringVar(&word, "word", "foo", "a word to be rotated by 13 places")
	flag.Parse()
	letters := []rune(word)
	
	for i, letter := range letters {
		letters[i] = transform(letter)
	}

	crypted := string(letters)
	fmt.Println(crypted)
}
