package main

import (
	"fmt"
	"strings"
)

func main() {
	const N = 5
	parole := make([]string,N)

	for i := 0;i < N;i++{
		fmt.Scan(&parole[i])
	}

	primeVocali(parole)
}

//dato un vettore di stringhe, controllo per ogni stringa quale sia la prima vocale contenuta e la stampo, altrimenti stampo che non ne contiene
func primeVocali(parole []string) {
	var check int
	for i:= range parole {
		parola := strings.ToLower(parole[i])
		check = 0
		for y:= 0; y <  len(parola);y++ {
			switch rune(parola[y]) {
				case 'a':
					fmt.Println(parola, ": ", string('a'))
					check = 1
					break
				case 'e':
					fmt.Println(parola, ": ", string('e'))
					check = 1
					break
				case 'i':
					fmt.Println(parola, ": ", string('i'))
					check = 1
					break
				case 'o':
					fmt.Println(parola, ": ", string('o'))
					check = 1
					break
				case 'u':
					fmt.Println(parola, ": ", string('u'))
					check = 1
					break
			}
			if check == 1{
				break
			}
		}
		if check == 0{
			fmt.Println("La parola ", parola, " non contiene vocali")
		}
	}
}
