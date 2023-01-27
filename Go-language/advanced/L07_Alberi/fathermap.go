package main

import (
	"fmt"
)

/*
 * Per implementare una tabella dei padri è possibile utilizzare una mappa, una struttura che consente di associare un valore a una chiave.
 * In questo caso la chiave potrebbe essere l'elemento figlio e il valore potrebbe essere il padre.
 * Per aggiungere un elemento alla mappa si può usare la sintassi fathers["child"] = "parent"
 * Per accedere al valore associato a una chiave si può usare la sintassi parent := fathers["child"].
 * per iterare sulla mappa è sufficiente il seguente ciclo for:
 * for child, parent := range fathers {
 *	 fmt.Println(child," ha come padre ",parent)
 * }
 * utilizzo un albero con tabella dei padri per risolvere https://adventofcode.com/2019/day/6 
*/

type fathers map[string]string

var orbits = make(fathers)

func main() {
	orbits["B"] = "COM"
	orbits["C"] = "B"
	orbits["D"] = "C"
	orbits["E"] = "D"
	orbits["F"] = "E"
	orbits["G"] = "B"
	orbits["H"] = "G"
	orbits["I"] = "D"
	orbits["J"] = "E"
	orbits["K"] = "J"
	orbits["L"] = "K"

	// chiamiamo la funzione per calcolare le orbite indirette
	totalOrbits := 0
	for item := range orbits {
	        totalOrbits += countOrbits(item)
	}

	fmt.Println("Numero totale di orbite:", totalOrbits)
}

func countOrbits(item string) int {
    count := 0
    for item != "COM" {
        item = orbits[item]
        count++
    }
    return count
}

