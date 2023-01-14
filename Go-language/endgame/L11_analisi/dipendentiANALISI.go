// DIPENDENTI

/*
Uso la struttura dati albero.
Non sapendo quanti subordinati ha un dipendente, ho usato una mappa per i subordinati di una struttura dipendente (struttura con nome, e puntatore al superiore)
e una lista contenente tutti i dipendenti in elenco.
*/

package main

import "fmt"

type dipendente struct {
	name      string
	val       int
	superiore *dipendente
}

func main() {

	var subordinati map[dipendente][]dipendente = make(map[dipendente][]dipendente)
	//var superiori map[dipendente]dipendente = make(map[dipendente]dipendente)
	var lista []dipendente

	A := &dipendente{"A", 1, nil}
	B := &dipendente{"B", 3, A}
	C := &dipendente{"C", 12, A}
	D := &dipendente{"D", 30, A}
	E := &dipendente{"E", 9, C}
	F := &dipendente{"F", 87, D}
	G := &dipendente{"G", 54, D}

	subordinati[*A] = append(subordinati[*A], *B, *C, *D)
	subordinati[*C] = append(subordinati[*C], *E)
	subordinati[*D] = append(subordinati[*D], *F, *G)

	lista = append(lista, *A, *B, *C, *D, *E, *F, *G)

	stampaSubordinati(B, subordinati)
	fmt.Println()
	fmt.Println(quantiSenzaSubordinati(lista, subordinati))

}

func stampaSubordinati(dipendente *dipendente, subordinati map[dipendente][]dipendente) {
	for _, subordinato := range subordinati[*dipendente] {
		fmt.Print(subordinato.name, " ")
	}
}

func quantiSenzaSubordinati(lista []dipendente, subordinati map[dipendente][]dipendente) int {
	count := 0
	for _, dipendente := range lista {
		_, haSub := subordinati[dipendente]
		if !haSub {
			count++
		}
	}
	return count
}

func supervisore(dipendente dipendente) {
	fmt.Println(dipendente.superiore)
}

/* Traccia per l'analisi:
Considerate l’esercizio “Dipendenti” della scheda di modellazione con alberi
e lo svolgimento proposto in questo programma.

1. Il comento descrive come si può modellare la situazione (domanda 1, parte 1)?
	* Se sì, la descrizione è corretta, chiara, completa?
	* Se no, considerate quanto scritto nel codice e completate il commento rispondendo alla domanda 1, parte 1.

2. Il comento descrive come è stata implementata la struttura dati scelta (domanda 2)?
	1. Se sì, valutate se la descrizione è coerente con quanto scritto nel codice
	ed eventualmente correggete/chiarite tale descrizione.
    2. Se no, considerate quanto scritto nel codice e descrivete come è stata implementata la struttura dati scelta.

3. La struttura dati scelta modella appropriatamente la situazione descritta nella traccia dell’esercizio?
(se vi è incoerenza tra quanto scritto nel commento e quanto implementato nel codice, fate riferimento al codice).

4. Valutate le scelte relative all’implementazione della struttura dati:
le considerate appropriate per potrer svolgere i compiti in maniera efficiente?
Altrimenti, come le correggereste/migliorereste?

*/
