package main

import (
	"fmt"
)

//inserisco N interi terminando con 0, stampo il numero di coppie di numeri uguali consecutivi
func main() {
	n := 1
	save := 0
	var count int
	for n != 0 {
		fmt.Scan(&n)
		if n == save {
			count++
		}
		save = n
	}

	fmt.Println("coppie uguali: ",count)
}
