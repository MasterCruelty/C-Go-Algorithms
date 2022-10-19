package main

import (
	"fmt"
	"strings"
	"math"
)

func main() {
	const N = 10
	parole := make([]string, N)

	for i:=0;i < N;i++ {
		fmt.Scan(&parole[i])
	}

	fmt.Println(quanteConA(parole))
	fmt.Println(primaConA(parole))
	fmt.Println(piuCorta(parole))
}

//dato un vettore di stringhe, contare quante stringhe contengano il carattere 'a'
func quanteConA(parole []string) int {
	var result int
	for i:= range parole {
		parola := strings.ToUpper(parole[i])
		for _, y:= range parola {
			if y == 'A' {
				result++
				break
			}
		}
	}
	return result
}

//dato un vettore di stringhe restituisce la prima stringa che contiene il carattere 'a' oppure la stringa vuota
func primaConA(parole []string) string {
	var result string
	for i := range parole{
		result = strings.ToUpper(parole[i])
		for _, y:= range result {
			if y == 'A' {
				return parole[i]
			}
		}
	}
	return "Nessuna stringa con carattere a"
}

//dato un vettore di stringhe restituisce la lunghezza della stringa più corta in termini di byte
func piuCorta(parole []string) int {
	min := math.MaxInt32
	for i:= range parole{
		if len(parole[i]) < min {
			min = len(parole[i])
		}
	}
	return min
}
