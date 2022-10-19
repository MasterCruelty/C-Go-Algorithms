


package main

import (
	. "fmt"
	"strings"
	)

func main(){

	var word1, word2 string
	Print("Inserisci prima parola: ")
	Scan(&word1)
	Print("Inserisci seconda parola: ")
	Scan(&word2)
	word1 = strings.ToUpper(word1)
	word2 = strings.ToUpper(word2)
	Println(word1,word2)
	var map1 map[string]int = map[string]int{}
	var map2 map[string]int = map[string]int{}

	for _, v:= range word1{
		vv := string(v)
		if v > 64 && v < 91 {
			map1[vv]++
		}
	}
	for _, v:= range word2{
		vv := string(v)
		if v > 64 && v < 91 {
			map2[vv]++
		}
	}
	Println(map1,map2)

	if len(map1) != len(map2) {
		Println("Le due parole non sono anagrammi")
		return
	}

	for key1, _ := range map1{
		if map1[key1] != map2[key1] {
			Println("Le due parole non sono anagrammi")
			return
		}
	}
	Println("Le due parole sono anagrammi")
}

