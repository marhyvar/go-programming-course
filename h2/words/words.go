package main

import (
	"os"
	"log"
	"fmt"
	"bufio"
	"flag"
)

func scanHowManyWords() {
	var filename string
	counter := 0
	flag.StringVar(&filename, "filename", "test.txt", "a file to scan")
	flag.Parse()
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		counter++
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println("Words counted: " , counter)
}

func main() {
	scanHowManyWords()
}
