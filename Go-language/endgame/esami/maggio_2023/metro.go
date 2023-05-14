package main


/*
 * Punto 1.
 * La rete metropolitana è modellabile tramite grafo come struttura dati.
 * Ogni stazione rappresente un nodo del grafo. Incluse le stazioni di interscambio, 
 * che permettono il cambio di linea differenziando per esempio con Cadorna 1 il nodo della linea 1,
 * e Cadorna 2 il nodo della linea 2, collegate tramite un arco.
 * I tratti tra una stazione e un'altra lo possiamo considerare come un cammino. I pesi sugli archi sono fissi a 1.
 * Ne risulta un grafo pesato(seppur con peso fisso), non orientato e fortemente connesso.
 * Da ogni stazione posso raggiungere tutte le altre e posso anche tornare indietro senza vincolo di direzione.
 * 
 * Punto 2.
 * Avendo a disposizione il grafo con il nodo di partenza e il nodo di arrivo, è possibile calcolare il tempo minimo,
 * ovvero la lunghezza del cammino minimo, tra P e A utilizzando una visita in ampiezza del grafo.
 * Dunque utilizziamo P come nodo d'origine e visitiamo i nodi adiacenti a P. Successivamente verranno ispezionati i nodi adiacenti e
 * visitati i loro "vicini" fino a raggiungere il nodo A. Al termine dell'esecuzione della visita vado a confrontare le lunghezze e restituisco
 * la lunghezza minima del cammino tra P e A.
 *
 * Punto 3.
 * l'algoritmo progettato e utilizzabile per il problema è quello descritto nel punto 2 prendendo in ingresso il grafo e i nodi di partenza e arrivo.
 * Il grafo può essere implementato in modo efficiente tramite una mappa che ha come chiave il nome di una stazione e come valore l'elenco delle stazioni 
 * adiacenti.
*/

import (
		"fmt"
		"bufio"
		"os"
		"strings"
)

type stazione struct{
	nome string
	linea int
	interscambio bool
}

type rete struct{
	adj map[string][]*stazione
}

/*
 * Funzione che legge i dati del grafo della metropolitana come file di testo 
 * e restituisce il grafo come tipo rete popolandolo.
*/
func leggiDati(nomeFile string) rete{
	myFile, err := os.Open(nomeFile)
	if err != nil{
		fmt.Println("file non trovato")
		os.Exit(0)
	}
	defer myFile.Close()

	scanner := bufio.NewScanner(myFile)
	var temp, num string
	interscambio := make(map[string]int)
	var staz []stazione
	metro := &rete{adj:nil}

	for scanner.Scan() {
		linea := scanner.Text()
		stazioniLinea := strings.Split(linea,":")
		fmt.Sscanf(stazioniLinea[0],"%s %s",temp,num)
		numLinea, _ := strconv.Atoi(num)
		stazioni := strings.Split(stazioniLinea[1],"; ")

		for _,v := range stazioni{
			stazione := &stazione{nome: v,linea: numLinea,interscambio:false}
			staz = append(staz,*stazione)
		}
	}
	for _, j:= range staz{
		for _, i := range staz{
			if(j.nome == i.nome && (j.linea != i.linea)) {
				j.interscambio = true
				i.interscambio = true
				//inserire direttamente le stazioni come adiacenti
			}
		}
	}
	//ciclo per popolare la rete e inserire per ogni stazione i suoi vicini e restituire
	return metro
}
