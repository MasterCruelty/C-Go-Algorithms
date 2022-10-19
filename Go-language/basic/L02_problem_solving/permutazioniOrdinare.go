package main

import  "fmt"


type Permutazioni struct{
	key int
	name string
}

//dato un vettore di struct non ordinato da 1 a N.
//nella posizione 0 l'elemento di chiave N, nella posizione 1 l'elemento di chiave N-1 e così via.
func main(){
	permutazioni := []Permutazioni{
		{key: 6,name: "Francesco"},
		{key: 1,name: "Andrea"},
		{key: 5,name: "Elisa"},
		{key: 2,name: "Beatrice"},
		{key: 3,name: "Carlo"},
		{key: 4,name: "Dino"},
		{key: 7,name: "Giorgia"},
		{key: 9,name: "Irene"},
		{key: 8,name: "Henry"},
	}
	fmt.Println("vettore non ordinato: ",permutazioni)
	for i := 0;i < len(permutazioni); i++ {
		key := permutazioni[i].key
		permutazioni[i],permutazioni[len(permutazioni)-key] = permutazioni[len(permutazioni) - key],permutazioni[i]
	}
	fmt.Println("\nvettore  ordinato: ",permutazioni)
}
