package main

import(
	"fmt"
	"math"
)

//Data una sequenza N di interi, stampare il numero di salite da sinistra a destra
//nella sequenza 9 1 3 5 2 0 8 6. Le salite sono due: 1 3 5 e 0 8 quindi stampo 2.
//1 3 e 3 5 non sono salite perchè la prima non finisce in un punto di massimo e la 
//seconda non inizia in un punto di minimo
func main(){
	var current,prev,succ int
	prev = math.MaxInt32
	fmt.Scan(&current)
	count := 0

	for {
		_, err:= fmt.Scan(&succ)
		if err != nil {
			break
		}
		//se si verifica la seguente condizione significa che vi è stato un punto di massimo e ora
		//sono in un punto di minimo, quindi posso contare una salite(appena terminata) e cercare il prossimo picco
		if current < prev && current < succ {
			count++
		}
		prev = current
		current = succ
	}
	fmt.Println(count)
}
