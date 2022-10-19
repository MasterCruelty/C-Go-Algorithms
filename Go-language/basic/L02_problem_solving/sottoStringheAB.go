package main

import (
	"bufio"
	"fmt"
	"os"
)
//data una stringa di N caratteri nell'alfabeto {a,b,c}, stampare il numero di sottostringhe che iniziano con a e finiscono con b.
//le sottostringhe possono sovrapporsi.
func main(){
	countA := 0
	count := 0

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	//conto tutte le a che sono l'inizio di una possibile sottostringa 'ab' 
	//quando trovo una b incremento il contatore di sottostringhe con il numero di a trovate finora
	for _,char := range scanner.Text() {
		if char == 'a' {
			countA++
		}
		if char == 'b'{
			count += countA
		}
	}
	fmt.Println(count)
}
