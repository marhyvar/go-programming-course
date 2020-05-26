package main

import (
	"fmt"
	"math/rand"
	"time"
)

// random integer, 0 <= integer < max
func randomInteger(min, max int) int {
	return min + rand.Intn(max - min)
}

// arpoo lottorivin, 7 eri numeroa väliltä 1 - 40
func lottorivi() {
	var numerot []int
	for i :=1; i < 41; i++ {
		numerot = append(numerot, i)
	}	
	
	rand.Seed(time.Now().UnixNano())
	for i :=0; i < 33; i++ {		
    	pituus := len(numerot)
		indeksi := randomInteger(0, pituus)
		numerot = append(numerot[:indeksi], numerot[(indeksi + 1):]...)
	}	
	fmt.Println("Kokeile näitä lotossa. ", numerot)
}

func main() {
	lottorivi()
}