package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
- E' possibile utilizzare i grafi per modellare il gioco "scale e serpenti"
- I nodi possono rappresentare la posizione della pedina, mentre un arco che
collega due vertici v1 e v2, se è possibile raggiungere v2 partendo da v1.
- Il grafo sarà orientato, dal momento che non è possibile tornare alla casella
precedente una volta fatta una mossa, a meno che non faccio una mossa e
finisco in una casella con un serpente, che mi riporta alla casella precedente
(ma questo è un caso raro).
- Il grafo generalmente non è connesso, ad eccezione di casi eccezionali. Ad
esempio se in ogni casella c'è un serpente che riporta alla casella precedente,
il grafo risultante sarà connesso.
- Il grafo non è pesato, in quanto ci basta sapere da una casella (nodo), quale
altra casella è possibile raggiungere

Approccio al problema:
- Per memorizzare i dati scelgo di utilizzare una mappa. La preferisco rispetto
ad un array in quanto ci evita di utilizzare più memoria (non dobbiamo salvare
tutti gli 0). Se una cella non ha nè basi di scale, nè bocche di serpente, il suo
numero non sarà presente come chiave nella mappa, altrimenti la mappa avrà una
chiave pari al numero della casella e il suo valore corrisponde al numero della
casella al quale si deve andare se si finisce sopra.
- Decido di sfruttare il concetto di pila. In un ciclo continuo ad iterare
finchè la pila non è vuota. Ad ogni iterazione aggiungo alla pila tutte le
possibili mosse che è possibile fare da una certa casella. Alla prossima
iterazione preleverò l'ultima mossa salvata, e la eseguo.
- Per evitare di andare avanti all'infinito, tengo traccia su quali caselle sono
già stato. In quanto stiamo cercando il percorso più breve, è inutile passare
per una casella già visitata in precedenza. In questo caso non aggiungo nessuna
mossa alla pila.
- Nella pila salvo una struct di tipo T che contiene una slice contenente le mosse
dei dadi effettuate, e un intero che indica a che casella siamo arrivati. Il valore
size(slice_mosse) indica il numero di mosse compiute fino a quel momento.
- Quando l'intero all'interno della struct è pari ad n (con n pari al numero di
caselle) salvo l'intera struct in una slice di struct T. Se la slice (all'interno
della struct) che sto per inserire ha size minore di tutte quelle già salvate,
cancello tutte quelle presenti e salvo quella nuova. Se invece il size della slice
che sto per aggiungere è maggiore del size di quelle presenti non la salvo.
- Quando la pila è vuota mi sono rimasti nell'array tutte le posibili combinazioni
di lanci più brevi per quella tabella di gioco.
*/

type Move struct {
	moves []int
	cell  int
}

func main() {
	check_for_stairs_and_snake := false
	r, c, cells := read_input()
	n := r * c

	stack := make([]Move, 0)
	already_visited := make([]bool, n)
	var finish_moves []int

	for i := 0; i < n; i++ {
		already_visited[i] = false
	}

	stack = append(stack, Move{moves: make([]int, 0), cell: 1})

	for len(stack) != 0 {
		move := stack[0]
		stack = stack[1:]

		if move.cell == n && (len(move.moves) <= len(finish_moves) || len(finish_moves) == 0) {
			finish_moves = move.moves
		}

		value, exists := cells[move.cell]
		if exists && check_for_stairs_and_snake {
			if (!already_visited[value-1]) || (value == n) {
				stack = append(stack, Move{moves: move.moves, cell: value})
				already_visited[value-1] = true
			}
		} else {
			for i := move.cell + 1; (i < move.cell+7) && (i <= n); i++ {
				if (!already_visited[i-1]) || (i == n) {
					stack = append(stack, Move{moves: append(move.moves, i-move.cell), cell: i})
					already_visited[i-1] = true
				}
			}
		}
	}

	fmt.Print("Puoi arrivare alla casella finale con ")
	fmt.Print(len(finish_moves))
	fmt.Println(" mosse e con i seguenti lanci:")

	for i := 0; i < len(finish_moves); i++ {
		fmt.Print(finish_moves[i])
		fmt.Print(" ")
	}

	fmt.Println()
}

func read_input() (r int, c int, cells map[int]int) {
	var c1, c2 int

	file, err := os.Open("input.txt")

	if err == nil {
		scanner := bufio.NewScanner(file)

		if scanner.Scan() {
			line := scanner.Text()

			fmt.Sscanf(line, "%d %d", &r, &c)
			cells = make(map[int]int)
		}

		for scanner.Scan() {
			fmt.Sscanf(scanner.Text(), "%d %d", &c1, &c2)
			cells[c1] = c2
		}
	}

	return
}

/*
Traccia per l'analisi:

Considerate l’esercizio “scale e serpenti” della scheda su “modellazione con grafi” e lo svolgimento proposto in questo programma.

Leggete innanzitutto il commento e valutate quanto segue:
1. La descrizione del grafo è chiara e corretta? In caso contrario, come la correggereste/chiarireste?
risposta 1. La descrizione è chiara e corretta.


2. Il grafo descritto modella appropriatamente la situazione descritta nella traccia dell’esercizio?


3. Il commento contiene una formulazione del problema da risolvere il termini di grafi? Se sì, quale è questa formulazione? Se no, come lo formulereste?
Risposta 3. il problema viene approcciato in termini di pila, non di grafi seppur descritto come tale. Andrebbe formulato come il problema di trovare il 
	    minor numero di nodi visitati per raggiungere il nodo numero n a partire dal nodo 1.


4. Il commento contiene una descrizione dell’algoritmo per risolvere il problema (poi implementato nel programma)? Se sì, dove si trova questa descrizione?
Risposta 4. Si, si trova nella sezione "approccio al problema" dove viene spiegato come affrontare il problema tramite l'uso del concetto di pila.


Considerate ora il codice:
1. Come sono rappresentati nel programma i vertici e gli archi del grafo? Sono rappresentazioni implicite o esplicite?
Risposta 1. il grafo è rappresentato tramite mappa di interi, sono rappresentazioni implicite senza esplicitazione dell'uso di questa struttura dati.


2. Cosa fa il ciclo for di riga 85-89? Provate a descriverlo facendo riferimento al grafo usato per modellare l’evoluzione del gioco.


3. Considerate la variabile “stack”. Che tipo di struttura dati implementa?
Risposta 3. Implementa la struttura dati Pila come descritto nel commento iniziale.


4. A cosa serve il controllo di riga 74? Vi pare necessario? Giustificate la risposta
Risposta 4. Il controllo sembra essere finalizzato al vedere se si ha consumato tutta la board, ma non è necessario poiché un controllo simile è sul ciclo 
	    sulle righe 85-89(?)


5. Se avevate risposto no all’ultima domanda a proposito del commento, date voi una descrizione dell’algoritmo implementato per risolvere il problema (è utile descriverlo facendo riferimento al grafo).
6. Vi sembra che questo algoritmo risolva correttamente il problema? Giustificate la risposta (di nuovo, è utile fare riferimento al grafo).


*/
