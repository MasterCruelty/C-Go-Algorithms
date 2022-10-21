package main

import (
	"fmt"
	"math/rand"
	"time"
)

//in questo sorgente vi sono 3 algoritmi di ordinamento: insertion, selection e merge.
//I primi due sono cosiddetti algoritmi di ordinamento elementari e hanno una complessità quadratica rispetto al tempo O(n^2).
//mergesort è più avanzato e riesce ad avere una complessità poco più che lineare rispetto al tempo O(n*log_n).
func main() {
	var c string
	numeri := []int{}
	fmt.Print("Vuoi inserire i numeri? s/n\n")
	fmt.Scan(&c)
	if c == "n"{
		fmt.Println("Allora verranno generati casualmente(lunghezza vettore 1000, verranno stampati i primi 30 elementi..)")
		numeri = crea_vettore()
	} else{
		numeri = inserisci_vettore()
	}
	fmt.Println("Vettore iniziale: ",numeri[0:30])
	fmt.Print("\n\nVettore ordinato con insertion_sort: ")
	start := time.Now()
	insertion_sort(numeri)
	insertion := time.Since(start)
	fmt.Println("\nTempo impiegato da insertionsort: ",insertion)
	fmt.Print("\n\nVettore ordinato con selection_sort: ")
	start2 := time.Now()
	selection_sort(numeri)
	selection := time.Since(start2)
	fmt.Println("\nTempo impiegato da selectionsort: ",selection)
	fmt.Print("\n\nVettore ordinato con merge_sort: ")
	start3 := time.Now()
	merge_sorted := merge_sort(numeri)
	fmt.Println(merge_sorted[0:30])
	merge := time.Since(start3)
	fmt.Println("\nTempo impiegato da mergesort: ",merge)
}



func insertion_sort(numeri []int) {
	var key,j int
	for i:= 1;i< len(numeri);i++ {
		key = numeri[i]
		j = i - 1
		//muovo gli elementi dell'array che sono più grandi di key a una posizione più avanti rispetto alla
		//loro posizione attuale
		for j >= 0 && numeri[j] > key {
			numeri[j+1] = numeri[j]
			j--
		}
		numeri[j+1] = key
	}
	fmt.Println(numeri[0:30])
}

func selection_sort(numeri []int) {
	var min = -1
	for i:= 0;i < len(numeri);i++{
		min = i
		for j := i+1;j< len(numeri)-1;j++{
			if numeri[j] < numeri[min]{
				min = j
			}
		}
		scambia(numeri[min],numeri[i])
	}
	fmt.Println(numeri[0:30])
}

func merge_sort(numeri []int) []int {
	var mid int
	if len(numeri) > 1 {
		mid = len(numeri) / 2
		return merge(merge_sort(numeri[:mid]),merge_sort(numeri[mid:]))
	}else{
		return numeri
	}
}

func merge(left,right []int) []int {
	size,i,j := len(left)+len(right),0,0
	result := make([]int,size)
	for k:= 0;k < size;k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			result[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			result[k] = left[i]
			i++
		} else if left[i] < right[j] {
			result[k] = left[i]
			i++
		}else {
			result[k] = right[j]
			j++
		}
	}
	return result
}


func crea_vettore() []int {
	numeri := make([]int,1000)
	for i:= 0;i < 1000;i++{
		numeri[i] = rand.Intn(1000)
	}
	return numeri
}

func inserisci_vettore() []int {
	numeri := make([]int,30)
	var n int
	for i:=0;i < 30;i++{
		fmt.Scan(&n)
		numeri[i] = n
	}
	return numeri
}

func scambia(i,k int) {
	save := i
	i = k
	k = save
}
