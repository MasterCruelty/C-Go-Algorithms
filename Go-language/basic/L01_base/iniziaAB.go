package main

import(
	"fmt"
	"strings"
)

func main(){
	const N = 10
	stringhe := make([]string,N)

	for i:= 0;i < N;i++ {
		fmt.Scan(&stringhe[i])
	}

	iniziaAB(stringhe)
}

//dato un vettore di stringhe, per ogni stringa controllo se inizia con a o con b, stampo il numero di quelle che iniziano con a e b
func iniziaAB(stringhe []string) {
	countA := 0
	countB := 0
	for i:= range stringhe{
		stringhe[i] = strings.ToLower(stringhe[i])
		if (stringhe[i])[0] == 'a'{
			countA++
		} else if (stringhe[i])[0] == 'b' {
			countB++
		}
	}
	fmt.Println("Le stringhe che iniziano per a: ", countA)
	fmt.Println("Le stringhe che iniziano per b: ",countB)
}
