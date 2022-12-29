package main

import "fmt"

/*
 * SunnyDale è una città che ospita un elevatissimo numero di vampiri. I vampiri non possono sopportare la luce solare.
 * Unico modo per spostarsi è passare per le gallerie sotterranee.
 * Ma vi sono dei fori nel soffitto delle gallerie, i quali fanno passare raggi solari.
 * Harmony segue una regola semplice e tassativa: ad ogni svincolo sceglie sempre e comunque la galleria meno luminosa.
 * Harmony vuole andare a trovare la sua amica Sarah.
 * 
 * INFORMAZIONI A DISPOSIZIONE:
 * N è il numero degli svincoli(numerati da 1 a N)
 * M è il numero delle gallerie.
 * H è l'indice dello svincolo dove abita Harmony.
 * S è l'indice dello svincolo dove abita Sarah.
 * Per ogni galleria, si sa quali svincoli collega e si sa la sua luminosità, indicata da un numero intero.(Maggiore è il numero, maggiore è la luminosità)
 *
 * PROBLEMA:
 * Sapendo che non esistono due gallerie egualmente luminose, determinare se Harmony sia in grado di raggiungere Sarah, e in caso affermativo dire in quante gallerie.
 *
 *
 *
 * MODELLAZIONE:
 * Cosa rappresentano i nodi? ==> I nodi sono rappresentati dagli svincoli.
 * Cosa rappresentano gli archi? ==> Gli archi sono rappresentati dalle gallerie.
 * Che caratteristiche ha il grafo? ==> Il grafo è un grafo orientato e pesato.(Non connesso siccome non è detto Harmony sia in grado di raggiungere Sarah)
 * 
 * Il problema in termini di grafi può essere formulato come il trovare il cammino di peso minimo dalla galleria di Harmony a quella di Sarah.
 * 
 * 
 * 
 * PROGETTAZIONE ALGORITMO PER RISOLVERE IL PROBLEMA:
 * Seguendo la strategia di Harmony, il problema può essere risolto tramite strategia greedy. Quindi ad ogni svincolo scelgo sempre quello con luminosità minore e
 * non torno mai indietro sulle mie scelte fino a raggiungere Sarah(se possibile).
 * All'inizio ordino i dati sulle gallerie in funzione del grado di luminosità e dopo di che uso la strategia di Harmony.
 * 
 * 
 *
 * IMPLEMENTAZIONE:
 * Uso della struttura dati grafo con rappresentazione con una slice di slice di adiacenza
 * input: N M S H
 * per ogni M input: A B L
 * A e B indicano lo svincolo(arco) e L il grado di luminosità
 *
 *
 *
 * ESEMPI INPUT:
 * 5 6 1 5
 * 1 2 5
 * 2 3 1
 * 3 4 3
 * 4 5 2
 * 5 1 6
 * 1 4 4
 * OUTPUT ASPETTATO:
 * 2
 *
 * ESEMPIO 2:
 * 3 2 2 1
 * 3 1 2
 * 2 3 1
 * OUTPUT ASPETTATO:
 * -1
*/

type arcoUscente struct {
	v int //nodo di destinazione
	l int //luminosità dell'arco
}

type grafo struct {
	n int                     // numeri di nodi
	m int                     // numero di archi
	adiacenti [][]arcoUscente // adiacenti[i] è la slice delle gallerie che partono da i
}

func stampaGrafo(g grafo) {
	fmt.Println("\nIl grafo ha",g.n,"nodi:")
	for i := 1; i <= g.n; i++ {
		fmt.Println(i,":",g.adiacenti[i])
	}
	fmt.Println()
}

// restituisce il numero di gallerie da attraversare oppure -1
// uso di visita in ampiezza per visitare i nodi del grafo
func calcolaCammino(g grafo,from int,to int) int {
	//se non ci sono gallerie che partono dal nodo di partenza, non mi posso muovere e restituisco -1
	if g.adiacenti[from] == nil {
		return -1
	}

	v := from
	count := 0
	//mappa ausiliaria per tenere conto dei nodi visitati durante il percorso
	//Probabilmente non necessaria considerando che Harmony non torna indietro sui suoi passi nella sua strategia
	visited := make(map[int]bool)

	for v != to {
		//setto come visitato il nodo corrente
		visited[v] = true
		fmt.Println(visited)
		min := 1000
		minTo := -1
		//visito i vicini del nodo v e segno il vicino con luminosità minima tra tutti
		for _, galleria := range g.adiacenti[v] {
			fmt.Println("visito vicini di",v,":",galleria,visited[galleria.v],galleria.l)
			if galleria.l < min {
				min = galleria.l
				minTo = galleria.v
			}
		}
		fmt.Println("luminosità minima:",min,"verso il nodo",minTo)
		//controllo se il nodo appena trovato con luminosità minima è stato visitato altrimenti riprendo la visita partendo da quello
		//e incremento il contatore delle gallerie da visitare
		if visited[minTo] {
			break
		} else {
			v = minTo
			count++
		}
	}
	if v == to {
		fmt.Println("numero di gallerie da attraversare:",count)
		return count
	} else {
		fmt.Println(-1)
		return -1
	}
}

//main per testare l'algoritmo di risoluzione del problema
func main() {
	var n int // numero svincoli-nodi da 1 a N
	var m int // numero delle gallerie-archi
	var h int // indice dello svincolo in cui si trova Harmony
	var s int // indice dello svincolo della casa di Sarah

	fmt.Scan(&n,&m,&h,&s)
	fmt.Println(n,m,h,s)

	var g grafo

	g.n = n
	g.m = m
	g.adiacenti = make([][]arcoUscente,n+1) // lascio vuoto il primo vertice
	var a,b,l int
	for i := 0;i < m; i++ {
		fmt.Scan(&a,&b,&l) // arco: da,a,luminosità

		g.adiacenti[a] = append(g.adiacenti[a],arcoUscente{b,l})
		g.adiacenti[b] = append(g.adiacenti[b],arcoUscente{a,l})
	}

	stampaGrafo(g)
	fmt.Println(calcolaCammino(g,h,s))

}
