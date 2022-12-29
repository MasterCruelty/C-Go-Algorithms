package main

import "fmt"

/*
 * DESCRIZIONE DEL GIOCO:
 * Il gioco Scale e Serpenti si gioca su una griglia rettangolare che contiene numeri, scale e serpenti, come nella figura presente in questa directory.
 * r e c sono il numero di righe e colonne nella griglia. n = r * c è il numero totale di caselle. 
 * la griglia contiene i numeri da 1 a n, uno per casella.
 * Si parte mettendo la pedina nella casella 1, si lancia un dado e si sposta avanti la pedina in base al numero ottenuto.
 * Le scale e i serpenti determinano dei salti:
 * quando si raggiunge una casella che contiene la base di una scala, si deve salire la scala fino a raggiungere la casella in cui si trova la cima della scala.
 * quando si raggiunge una casella con la bocca di un serpente, si deve scendere fino alla casella in cui si trova la coda del serpente.
 * L'obiettivo è raggiungere la casella numerata con n. Se con un lancio si supera il numero n, la partita è persa.
 *
 *
 * ESEMPIO:
 * nella figura scale_serpenti.png, se si lancia il dado tre volte ottenendo 2,6,2 si procede come segue:
 * Da casella 1, avendo lanciato 2, si raggiunge casella 3 che contiene base scala, quindi si sale fino a casella 22.
 * Da casella 22, avendo lanciato 6, si raggiunge casella 28.
 * Da casella 28, avendo lanciato 2, si raggiunge 30.
 * Si vince la partita anche con le sequenze di lanci (2,2,6) oppure (2,5,2,6,2) oppure (5,6,6,6,6)
 *
 *
 * DESCRIZIONE DEL PROBLEMA:
 * Si vuole stabilire il numero minimo di mosse(lanci di dado) necessarie per vincere la partita.
 *
 *
 *
 * MODELLAZIONE:
 * Il gioco in termini di grafi può essere descritto come il trovare il minor numero di nodi visitati per raggiungere il nodo numero n a partire dal nodo 1.
 * Ogni casella rappresenta un nodo, il grafo è connesso. Ogni casella ha un arco che la porta al nodo successivo. La caselle dotate di scale e/o serpente,
 * hanno un'ulteriore arco verso un nodo di numero più alto o più basso. Il grafo inoltre non è pesato, non vi è un peso per raggiungere un nodo da un altro.
 *
 * Ai fini del problema non è necessario memorizzare tutte le informazioni relative alla griglia di gioco. 
 * Per calcolare solo le mosse, la griglia può essere descritta tramite 
 * la sequenza contenente il numero n e la sequenza di interi a1,a2,...,an che indicano gli eventuali salti.
 * Più precisamente: 
 * Se la casella numerata con i non contiene nè la base di una scale nè la bocca di un serpente, la casella ai è pari a 0.
 * Altrimenti ai è il numero della casella in cui saltare, una volta raggiunta la casella numerata con i.
 * 
 * Esempio descrizione sintetica sulla griglia proposta scale_serpenti.png
 * {0,0,22,0,8,0,0,0,0,0,26,0,0,0,0,0,4,0,7,29,9,0,0,0,0,0,1,0,0,0}
 * 
 *
 *
 * PROGETTAZIONE:
 * 1. Progettare u algoritmo che sia in grado di calcolare il numero minimo di lanci necessari per vincere la partita.
 * 2. Modificare l'algoritmo in modo che stampi anche una sequenza di lanci di lunghezza minima che consente di vincere la partita.
 * 3. Modificare l'algoritmo in modo che si possa specificare la casella di partenza.
 * 4. Modificare l'algoritmo in modo da evitare le mosse che usano scale e serpenti. 
 *
 *
 * ESEMPIO:
 * Se la griglia di gioco è quella rappresentata nella figura, il numero minimo di lanci per raggiungere 30 da 1 è 3(ad esempio con i lanci 2,6,2).
 * Senza usare scale nè serpenti, il numero minimo di lanci è 5(ad esempio con i lanci 5,6,6,6,6).
*/

type board struct {
	n int
	jumps map[int]int //jump[i] è la casella di destinazione se ci si trova in i
}

func setBoard() board {
	var r,c int
	fmt.Scan(&r,&c)
	n := r * c

	jumps := make(map[int]int)
	for {
		var start,end int
		_,err := fmt.Scan(&start,&end)
		if err != nil {
			break
		}
		jumps[start] = end
	}
	fmt.Println(jumps)
	return board{n,jumps}
}

//stampa del numero di mosse necessarie
func move(b board,start int,dice int) (end int) {
	if el, ok := b.jumps[start+dice]; ok {
		return el
	} else{
		return start + dice
	}
}

/*
 * Calcola il numero minimo di lanci necessari per arrivare alla casella di valore v partendo dalla casella di valore v0
 * Modello l'evoluzione del gioco con un grafo(i nodi rappresentano le configurazioni, gli archi le mosse)
 * Uso una visita in ampiezza del grafo(serve una coda)
 * Restituisco 0 se v e v0 coincidono; Restituisco -1 se partendo da v0 non si può raggiungere v
 * Se seq è pari a 1, stampa anche una sequenza di lanci di dadi di lunghezza minima(se c'è)
*/
func bfs(b board,v0,v int,seq bool) {
	if v0 == v {
		fmt.Println(0)
		return
	}
	prec := make(map[int]int) // per ogni casella, indica la casella precedente
	dice := make(map[int]int) // per ogni casella, indica conc he dado ci si è arrivati
	aux := make(map[int]int)  // caselle già visitate, con distanza da v0

	fmt.Println("\naux",aux)

	coda := []int{v0} // all'inizio la coda contiene solo la casella di partenza
	curr := v0
	aux[v0] = 0
	dice[v0] = -1
	fmt.Println("coda",coda,"\naux",aux)

	for len(coda) > 0 {
		curr = coda[0]
		coda = coda[1:]
		fmt.Println("\tposizione",curr)

		//metti nella coda tutti i vicini non ancora visitati
		for i := 1; i <= 6;i++ {
			end := move(b,curr,i)
			if aux[end] >= 1 {
				continue
			}

			if _, ok := aux[end]; !ok {
				coda = append(coda,end)
			}

			fmt.Println("end",end,"curr",curr,"i",i)
			prec[end] = curr
			dice[end] = i
			aux[end] = aux[curr] + 1

			if end == b.n {
				fmt.Println("bastano",aux[curr]+1,"mosse")
				if seq {
					printSeq(end,v0,prec,dice)
				}
				return
			}
		}
		fmt.Println("coda",coda,"\naux",aux)
	}
}

//stampa della sequenza di lanci
func printSeq(v,v0 int,prec map[int]int,dice map[int]int) {
	for v != v0 {
		fmt.Println(v,"venendo da",prec[v],"con tiro",dice[v])
		v = prec[v]
	}
}

func main() {
	b:= setBoard()
	fmt.Println(b.n,b.jumps)

	fmt.Println(move(b,3,3))
	bfs(b,14,30,true)
	bfs(b,1,30,true)
}
