package main

import "fmt"


/*
 * Complessità lineare O(n) con n numeri degli elementi nell'array
 * Ogni elemento viene controllato una volta sola
*/
func lenCrescente(a []int) []int {
	//alloco array che restituirò alla fine
	var result []int
	//se l'array è vuoto, lo restituisco direttamente
	if len(a) == 0 {
		return result
	}
	//mi segno il primo elemento dell'array per confrontarlo nella prima iterazione
	//il ciclo partirà da 1
	last := a[0]
	current := make([]int,0)
	/*
	 * In ogni iterazione confronto se l'elemento corrente è minore di last
	 * se è minore, current si re-inizializza da capo appendendo l'elemento corrente.
	 *
	 * Altrimenti a current si appende l'elemento corrente, poi viene controllato se
	 * la lunghezza di current è maggiore di quella di result, in caso result lo sovrascrivo.
	 *
	 * Ad ogni iterazione aggiorno il valore di last con l'elemento di indice i.
	 * Ogni elemento viene analizzato una volta sola
	*/
	for i:=1;i<len(a);i++{
		if a[i] < last{
			current = make([]int,0,1)
			current = append(current,a[i])
		}else{
			current = append(current,a[i])
			if len(current) > len(result) {
				result = current
			}
		}
		last = a[i]
	}
	return result
}

//main di test
func main() {
	fmt.Println(lenCrescente([]int{12,3,5,7,8,4,6,1,2,3,4,5,6,7,8}))
}
