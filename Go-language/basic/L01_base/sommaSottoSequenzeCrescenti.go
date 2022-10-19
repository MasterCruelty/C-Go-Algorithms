package main

import(
	"fmt"
	"math"
)

//Data una sequenza N di interi, stampare la somma degli interi in ciascuna sottosequenza crescente.
//nella sequenza 1 2 13 0 7 8 9 -1 0 2 le sottosequenze sono 3 e stampo le loro somme: 16,24 e 1..
//1 3 e 3 5 non sono salite perchè la prima non finisce in un punto di massimo e la 
//seconda non inizia in un punto di minimo
func main(){
	var current,prev,succ int
	prev = math.MaxInt32
	fmt.Scan(&current)
	count := 0
	somma := current
	for {
		_, err:= fmt.Scan(&succ)
		somma += succ
		//fmt.Println("succ: ", succ)
		//fmt.Println("somma: ",somma)
		if err != nil {
			break
		}
		//se si verifica la seguente condizione significa che vi è stato un punto di massimo e ora
		//sono in un punto di minimo, quindi posso contare una salite(appena terminata) e cercare il prossimo picco
		if current < prev && current < succ {
			count++
		}else if current > succ{
			somma -= succ
			fmt.Println(somma)
			somma = 0
		}
		prev = current
		current = succ
	}
	fmt.Println("numero di sotto sequenze crescenti: ",count)
}
