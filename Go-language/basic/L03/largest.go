package main

import "fmt"

func main() {
	var numbers := make([]int,0){}
	n := -1
	for n != 0 {
		fmt.Scan(&n)
		numbers = append(numbers,n)
	}
	fmt.Println(largest(numbers))

}


//questa funzione ha lo scopo di trovare il numero più grande in un vettore
//domande:

/*
     1) Come deve essere completato il caso base?
     risposta: if n == 1*/

//Assumiamo che il vettore ora contenga: [ 1,2,5,7,-2,10,9,21,3,8].

/*
    2)Durante l'esecuzione della chiamata largest(numbers), consideriamo
    la chiamata ricorsiva che termina per prima. Qual è l'argomento passato 
    in questa chiamata e quale valore restituisce la funzione?

    risposta: numbers[n-1], la funzione restituisce max(numbers[n-2],largest[numbers[n-2])
*/

/*
   3) Con quali argomenti viene eseguita per la prima volta la funzione max?
   
   risposta: numbers[n-1],largest(numbers[n-1])
*/

/*
  4) E con quali argomenti viene eseguita l'ultima volta la funzione max?

  risposta: numbers[0],largest(numbers[0])
*/
func largest(numbers []int) int{
	n := len(numbers)
	if ---  {
		return numbers[0]
	}
	return max(numbers[n-1], largest(----))
}
