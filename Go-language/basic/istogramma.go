package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
	)

func main() {
	//definisco lo scanner per l'input
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	//s1:= strings.Split(s," ")

	diz := make(map[byte]int)
	//popolo la mappa
	for i := 0;i< len(s);i++ {
		unicode.ToLower(rune(s[i]))
		if rune(s[i]) >= rune('a') && rune(s[i]) <= rune('z'){ //uso 'a' e 'z' invece dei numeri 97 e 122
			diz[s[i]]++
		}
	}
	//conto gli asterischi per ogni lettera della frase
	for i:= 'a'; i <= 'z';i++{
		if diz[byte(i)] > 0{
			fmt.Println(string(i), " ", strings.Repeat("*", diz[byte(i)]))
		}
	}
}
