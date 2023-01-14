/* MODELLAZIONE
NODI: svincoli delle gallerie
ARCHI: gallerie

c'è un arco tra 2 svincoli se esiste una galleria che li colleghi, e il peso su ogni arco indica la luminosità della galleria

il grafo è non orientato, connesso e pesato 
CORREZIONE: il grafo non necessariamente è connesso perché non è detto Harmony possa raggiungere Sarah.
	    Il grafo nel caso si adotti la strategia di harmony è orientato perché non torna indietro tra un arco e l'altro.

si tratta di un problema di esistenza di un cammino, scegliendo sempre e solo l'arco con peso minore
CORREZIONE: il problema non è di esistenza di un cammino, ma di trovare il cammino di peso minimo tra 2 nodi(se esiste)
	    la strategia usata è greedy e viene scelto sempre l'arco di peso minore senza tornare indietro sui propri passi.

#####fine modellazione.#######

ALGORITMO:  c'è una lista di nodi visitati, un dato che indica il n di archi attraversati, e un dato che indica in che nodo siamo 
	
	se sono in S ho finito
	se non sono in S: scelgo l'arco con peso minore e mi muovo nel nodo collegato con quell'arco,
	  	si aggiunge 1 alla variabile del numero di archi attraversati
	se il nodo è nella lista dei visitati: non esiste un cammino che rispetti la strategia
	se il nodo non è nella lista si inserisce il nodo nella lista e si ripete l'algoritmo

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type arcoMigliore []int

//pos 0 = vicino migliore
//pos 1 = luminosità

type Grafo = []arcoMigliore //indice = nodo di partenza

func main() {
	var numGallerie int
	var nodiVisitati []int
	var posizione int

	g, h, s := CreaGrafo()

	posizione = h

	for {
		if posizione == s {
			fmt.Println(numGallerie)
			return
		}
		posizione = g[posizione-1][0]
		numGallerie++
		if !(CercaInSlice(nodiVisitati, posizione)) {
			nodiVisitati = append(nodiVisitati, posizione)
		} else {
			fmt.Println(-1)
			return
		}
	}
}

func CreaGrafo() (g Grafo, h int, s int) { // restituisce il grafo con solo le gallerie sceglibili da harmony, il nodo di partenza dell'algoritmo (Harmony) e il nodo d'arrivo(Sarah)
	var num []int

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}

		campi := strings.Split(scanner.Text(), " ")
		for _, cont := range campi {
			n, _ := strconv.Atoi(cont)
			num = append(num, n)
		}
		if len(num) == 4 { //solo prima riga input
			g = make(Grafo, num[0])
			h = num[2]
			s = num[3]
			num = nil
			continue
		}

		if (g[num[0]-1] == nil) || (g[num[0]-1][1] > num[2]) {
			g[num[0]-1] = arcoMigliore{num[1], num[2]}
		}
		if (g[num[1]-1] == nil) || (g[num[1]-1][1] > num[2]) {
			g[num[1]-1] = arcoMigliore{num[0], num[2]}
		}

		num = nil
	}
	//fmt.Println(g)
	return g, h, s
}

func CercaInSlice(s []int, n int) bool {
	for _, num := range s {
		if num == n {
			return true
		}
	}
	return false
}

/*
Traccia per l'analisi:

Considerate l’esercizio “Sunnydale” della scheda su “modellazione con grafi” e lo svolgimento proposto in questo programma.
Rileggete le domande nella sezione 1.2 e valutate se il commento fornisce risposte corrette, chiare e complete
alle quattro domande. In caso contrario, come le correggereste/completereste?

Considerate ora il codice:
1. Il grafo che viene manipolato dal programma corrisponde a quanto descritto nel commento?
    1a. Se sì, descrivete come è implementato il grafo (è utile fare riferimento ai nomi dei tipi e delle variabili usate, e alle righe di codice rilevanti)
    Risposta lato 1a. ==> Il grado nel codice è implementato come una slice di slice di adiacenza(vedi arcoMigliore che è una slice di int mentre il tipo grafo 
    è una slice di arcoMigliore)

    
    2a.(esclusa) Se no, descrivete quale grafo è effettivamente implementato nel codice descrivete come è implementato (è utile fare riferimento ai nomi dei tipi e delle variabili usate, e alle righe di codice rilevanti)
    

    2b. Considerate la definizione del tipo arcoMigliore all riga 31 e verificate come viene usato nelle righe 85-89. Descrivete a parole questo tipo: cosa rappresentano gli elementi della slice? E’ una definizione appropriata? Perché? La modifichereste? Come?
    Risposta 2b: gli elementi della slice arcoMigliore rappresentano i nodi con minor luminosità vicini a ogni nodo del grafo. Dunque partendo da H e guardando
    arcoMigliore rispetto ad H so già dove spostarmi usando la strategia di Harmony.


3. Considerate l’if nelle righe 77-83: è possibile portarlo fuori dal ciclo? Come? E’ opportuno farlo? Perché?
Risposta 3. Si, è possibile portarlo fuori dall'if usando lo scanner sulle prime 4 psizioni della slice num, dopo dentro il ciclo non cambia niente.
	    È opportuno farlo per questioni di leggibilità del codice invece di avere un costrutto if dentro un ciclo che viene eseguito una singola volta.


4. Considerate la slice nodiVisitati definita alla riga 40. Cosa rappresenta e come viene usata? E’ una scelta appropriata? In caso negativo, come la modifichereste?
Risposta 4. La slice nodiVisitati rappresenta la lista dei nodi che sono già stati incontrati durante il percorso che harmony dovrebbe fare per raggiungere Sarah.
	    È appropriato perché se mi accorgo di essere in un nodo già visitato significa che sto tornando indietro, quindi so direttamente che non esiste
	    un cammino per raggiungere Sarah con la strategia di Harmony.


*/
