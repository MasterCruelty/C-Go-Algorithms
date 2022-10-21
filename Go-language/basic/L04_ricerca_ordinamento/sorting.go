package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var c string
	numeri := []int{}
	fmt.Print("Vuoi inserire i numeri? s/n\n")
	fmt.Scan(&c)
	if c == "n"{
		fmt.Println("Allora verranno generati casualmente")
		numeri = crea_vettore()
	} else{
		numeri = inserisci_vettore()
	}
	fmt.Println("Vettore iniziale: ",numeri)
	fmt.Print("\n\nVettore ordinato con insertion_sort: ")
	insertion_sort(numeri)
	fmt.Print("\n\nVettore ordinato con selection_sort: ")
	selection_sort(numeri)
	fmt.Print("\n\nVettore ordinato con merge_sort: ")
	merge_sort(numeri)
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
	fmt.Println(numeri)
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
	fmt.Println(numeri)
}

func merge_sort(numeri []int) {

}


func crea_vettore() []int {
	numeri := make([]int,25)
	for i:= 0;i < 25;i++{
		numeri[i] = rand.Intn(100)
	}
	return numeri
}

func inserisci_vettore() []int {
	numeri := make([]int,10)
	var n int
	for i:=0;i < 10;i++{
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
