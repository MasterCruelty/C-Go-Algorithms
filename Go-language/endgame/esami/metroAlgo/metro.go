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
		"strconv"
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
	var staz []stazione
	metro := &rete{adj:make(map[string][]*stazione)}

	//leggo il contenuto del file
	for scanner.Scan() {
		//formatto il testo letto
		linea := scanner.Text()
		stazioniLinea := strings.Split(linea,": ")
		fmt.Sscanf(stazioniLinea[0],"%s %s",temp,num)
		numLinea,_ := strconv.Atoi(strings.TrimPrefix(stazioniLinea[0], "Linea "))
		stazioni := strings.Split(stazioniLinea[1],"; ")
		//creo una slice contenente tutte le stazioni in ordine una dopo l'altra
		for _,v := range stazioni{
			stazione := &stazione{nome: v,linea: numLinea}
			staz = append(staz,*stazione)
		}
	}
	//identifico le stazioni di interscambio
	for i:= 0; i < len(staz); i++{
		for j:= 0;j < len(staz); j++{
			if(i != j && staz[i].nome == staz[j].nome && staz[i].linea != staz[j].linea) {
				staz[j].interscambio = true
				staz[i].interscambio = true
				metro.adj[staz[j].nome] = append(metro.adj[staz[j].nome],&staz[i])
				metro.adj[staz[i].nome] = append(metro.adj[staz[i].nome],&staz[j])
			}
		}
	}
	//popolo le slice degli adiacenti
	for i, stazione := range staz{
		if i > 0{
			metro.adj[stazione.nome] = append(metro.adj[stazione.nome],&staz[i-1])
		}
		if i < len(staz)-1{
			metro.adj[stazione.nome] = append(metro.adj[stazione.nome],&staz[i+1])
		}
	}
	return *metro
}


func main(){
	metro := leggiDati("linee.txt")
	adj_cadorna := metro.adj["Cadorna"]
	fmt.Println(adj_cadorna[0])
}