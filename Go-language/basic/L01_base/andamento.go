package main

import "fmt"

func main() {
	var precedente, attuale int
	fmt.Scan(&attuale)
	for attuale != 0 {
		precedente = attuale
		fmt.Scan(&attuale)
		if attuale >= precedente {
			fmt.Println("+")
		} else {
			fmt.Println("-")
		}
	}
}
