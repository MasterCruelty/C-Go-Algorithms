package main

import (
	"fmt"
)

/*
 * Scrivere una funzione con prototipo "func h(a []int, b []int, n int, m int) int"
 * Dato un vettore A di lunghezza n e un vettore B di lunghezza m, 
 * decide se ogni valore in A è strettamente più piccolo di ogni valore in B.
 * Se A = {1,5,5} e B={6,5} allora h(A,B) restituisce 0, mentre per C = {7,6} allora h(A,C) restituisce 1
 * Valutare la complessità: O(n+m)
 * primo ciclo che esegue operazioni elementari n volte dove n è la lunghezza di a
 * secondo ciclo che esegue operazioni elementari m volte dove m è la lunghezza di b
*/

func h(a []int,b []int, n int,m int) int{
	maxA := 0
	for i := 0;i<n;i++{
		if a[i] > maxA{
			maxA = a[i]
		}
	}
	minB := 99999
	for i := 0;i<m;i++{
		if b[i] < minB{
			minB = b[i]
		}
	}

	if maxA < minB{
		return 1
	}else{
		return 0
	}
}

func main(){
	a := []int{1,5,5}
	b := []int{6,5}
	c := []int{7,6}
	result1 := h(a,b,len(a),len(b))
	result2 := h(a,c,len(a),len(c))
	fmt.Println(result1)
	fmt.Println(result2)
}
