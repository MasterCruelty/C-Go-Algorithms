package main

import (
	"fmt"
)

//inserisco N interi e stampo quelli che sono inferiori rispetto al precedente e al successivo
func main() {
	const N = 5
	numeri := make([]int,N)

	for i := 0; i < N; i++ {
		fmt.Scan(&numeri[i])
	}

	for i:=1;i < N-1;i++{
		if numeri[i] < numeri[i-1] && numeri[i] < numeri[i+1] {
			fmt.Println("Giorno ",i, ": ",numeri[i])
		}
	}
}
