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
 * Il grafo può essere implementato in modo efficiente tramite una lista di adiacenze dove ad ogni stazione è associata la lista delle stazioni vicine.
 * Per quanto riguarda l'implementazione avremo una mappa che ha come chiave il nome di una stazione e come valore l'elenco delle stazioni vicine.
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
	adj map[string][]*stazione //mappa che dato il nome di una stazione, si hanno gli adiacenti
	linee map[int][]*stazione  //mappa che dato il numero di linea, si hanno le stazioni di quella linea in ordine
}

/*
 * Funzione che legge i dati del grafo della metropolitana come file di testo 
 * e restituisce il grafo come tipo rete popolandolo.
 * Valutazione complessità:
 * C'è un primo for che esegue una serie di operazioni elementari e lo fa per ogni riga del file.
 * Al suo interno vi è un altro for innestato che esegue un numero di operazioni elementare che dipende
 * dal numero di stazioni su ogni riga del file. Se identifico con R la riga del file e con S il numero di stazioni.
 * Questa prima parte ha complessità O(R*S).
 * Successivamente ho altri due for innestati, quello esterno esegue S operazioni e quello interno pure.
 * Questa seconda parte ha complessità O(S^2).
 * Infine ho un ultimo ciclo for che esegue S operazioni elementari quindi ha costo O(S).
 * O(R*S) + O(S^2) + O(S) = O(S^2)
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
	metro := &rete{adj:make(map[string][]*stazione),linee:make(map[int][]*stazione)}

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
			metro.linee[numLinea] = append(metro.linee[numLinea],stazione)
		}
	}
	//identifico le stazioni di interscambio e gestisco le adiacenze tra loro
	//Ad esempio Cadorna sulla linea 1 è adiacente a Cadorna sulla linea 2
	for i:= 0; i < len(staz); i++{
		for j:= 0;j < len(staz); j++{
			if(i != j && staz[i].nome == staz[j].nome && staz[i].linea != staz[j].linea) {
				staz[j].interscambio = true
				staz[i].interscambio = true
				staz[j].nome = staz[j].nome + "-" + strconv.Itoa(staz[j].linea)
				staz[i].nome = staz[i].nome + "-" + strconv.Itoa(staz[i].linea)
				metro.adj[staz[j].nome] = append(metro.adj[staz[j].nome],&staz[i])
				metro.adj[staz[i].nome] = append(metro.adj[staz[i].nome],&staz[j])
			}
		}
	}
	//popolo le slice degli adiacenti
	for i, stazione := range staz{
		if i > 0 && stazione.linea == staz[i-1].linea {
			metro.adj[stazione.nome] = append(metro.adj[stazione.nome],&staz[i-1])
		}
		if i < len(staz)-1 && stazione.linea == staz[i+1].linea{
			metro.adj[stazione.nome] = append(metro.adj[stazione.nome],&staz[i+1])
		}
	}
	return *metro
}

/*
 * Restituisce la slice con le stazioni della linea numLinea in ordine
 * Valutazione complessità:
 * O(1) prelevo direttamente dalla mappa le stazioni della linea richiesta
*/
func linea(metro rete,numLinea int) []*stazione{
	stazioni := metro.linee[numLinea]
	return stazioni
}

/*
 * Restituisce le stazioni vicine alla stazione in ingresso
 * Valutazione complessità:
 * l'assegnamento iniziale viene eseguito in O(1).
 * Il ciclo for esegue al più 3 iterazioni per popolare la slice delle vicine.
 * O(S) complessità finale, ma considerando le poche iterazioni è come se fosse tempo costante.
*/
func stazioniVicine(metro rete, s string) []string{
	vicine := metro.adj[s]
	result := make([]string,0)
	for i:= 0;i< len(vicine);i++{
		result = append(result,vicine[i].nome)
	}
	return result
}

/*
 * Restituisce la slice contenente i nomi delle stazioni di interscambio
 * Valutazione complessità:
 * Il primo ciclo scorre gli elementi della mappa metro.adj
 * Il secondo ciclo scorre la slice stazioni
 * Consideriamo R il numero di chiavi nella mappa e S il numero di stazioni per chiave
 * Complessità finale O(R*S)
*/
func interscambio(metro rete) []string{
	interscambiSet := make(map[string]bool)
    interscambi := make([]string, 0)

    for _, stazioni := range metro.adj {
        for _, stazione := range stazioni {
            if stazione.interscambio {
                nomeStazione := stazione.nome
                if !interscambiSet[nomeStazione] {
                    interscambiSet[nomeStazione] = true
                    interscambi = append(interscambi, nomeStazione)
                }
            }
        }
    }

    return interscambi
}

/*
 * Restituisce true se due stazioni sono sulla stessa linea altrimenti false.
 * Valutazione complessità:
 * O(1) sono tutte operazioni elementari.
*/
func stessaLinea(metro rete,s1 string,s2 string) bool{
	stazione1 := metro.adj[s1][len(metro.adj[s1])-1].linea
	stazione2 := metro.adj[s2][len(metro.adj[s2])-1].linea
	if stazione1 == stazione2{
		return true
	}else{
		return false
	}
}

//main di test per le funzioni
func main(){
	metro := leggiDati("linee.txt")

	fmt.Println("Quale stazione vuoi sapere chi è adiacente?")
	stazione := ""
	fmt.Scan(&stazione)
	vicine:= stazioniVicine(metro,stazione)
	fmt.Println(vicine)

	fmt.Println("Di quale linea vuoi sapere le stazioni?")
	line := 0
	fmt.Scan(&line)
	linea(metro,line)
	fmt.Println("Elenco delle stazioni di interscambio:\n")
	interscambi := interscambio(metro)
	fmt.Println(interscambi)
	fmt.Println("Di quali stazioni vuoi controllare se sono sulla stessa linea? ")
	s1,s2 := "",""
	fmt.Scan(&s1,&s2)
	ok := stessaLinea(metro,s1,s2)
	fmt.Println(ok)
}
