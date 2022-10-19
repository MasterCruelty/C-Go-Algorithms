package main

// import delle librerie necessarie
//bufio e os per lo scanner, fmt per la print
import (
	"bufio"
	. "fmt"
	"os"
	)

func main(){

	var count int
	scanner := bufio.NewScanner(os.Stdin) //definisco lo scanner
	scanner.Scan()
	s := scanner.Text() //scan della parola

	var lettera rune
	Scanf("%c", &lettera) //scan della lettera da cercare

	for i, _ := range s {
		if rune(s[i]) == lettera {
			count++
		}
	}
	Println("il carattere", string(lettera), "compare", count, "volte")
}
