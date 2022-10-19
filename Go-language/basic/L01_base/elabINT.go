package main

import (
	"fmt"
	"math"
)

func main(){
	const N = 10
	numeri := make([]int,N)

	for i:= 0;i < N;i++ {
		fmt.Scan(&numeri[i])
	}

	fmt.Println("Strano prodotto: ",stranoProdotto(numeri))
	pariDispari(numeri)
	fmt.Println("minDispari: ",minDispari(numeri))
}

//strano prodotto, trovo due numeri dentro il vettore che sono maggiori di 7 e multipli di 4, ne restituisco il prodotto
func stranoProdotto(numeri []int) int {
	var n1,n2 int
	for i:= 0;i < len(numeri);i++{
		if numeri[i] > 7 && numeri[i] % 4 == 0 {
			if n1 == 0{
				n1 = numeri[i]
				continue
			}else if n2 == 0{
				n2 = numeri[i]
			}
		}
	}
	return n1 * n2
}

//pari dispari, stampo per ogni elemento se è pari o dispari
func pariDispari(numeri []int)  {
	for i:= 0; i < len(numeri); i ++{
		if numeri[i] % 2 == 0{
			fmt.Println(numeri[i], ": pari")
		}else{
			fmt.Println(numeri[i], ": dispari")
		}
	}
}

//min dispari, trovo il numero dispari più piccolo nel vettore e lo restituisco
func minDispari(numeri []int) int{
	min := math.MaxInt32
	for i := 0;i < len(numeri); i++{
		if numeri[i] % 2 != 0 && numeri[i] < min{
			min = numeri[i]
		}
	}
	return min
}
