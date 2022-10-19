package main

import "fmt"


//data una permutazione 1..N stampare il numero di volte che è necessario tornare indietro a sinistra per 
//analizzare i numeri in ordine crescente.
//esempio: 4 5 1 3 6 2 ==> avremo bisogno di tornare indietro a sinistra 2 volte.
//analizziamo 1 andando a destra, poi il 2 è in fondo, torno a sinistra una volta per il 3 e poi di nuovo per il 4
//poi proseguo solo a destra per il 5 e il 6.
func main(){
	var n = make([]int,0)
	for {
		var numero int
		_, err := fmt.Scan(&numero)
		if err != nil {
			break
		}
		n = append(n,numero)
	}
	fmt.Println(n)
	fmt.Println("Inversioni necessarie:",raccogli(n))
}


func raccogli(n []int ) int{
	count := 0
	mapN := make (map[int]bool)
	for _, element := range n {
		mapN[element] = true
		if mapN[element+1]{
			count++
		}
	}
	return count

}
