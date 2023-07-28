package main

import (
	"fmt"
)
/*
 * Si consideri una griglia rettangolare formata da n righe e m colonne. Le caselle possono essere nere o bianche.
 * Percorso: una qualunque sequenza di caselle c0,..,ck della griglia con le seguenti proprietà:
 *           - le caselle della sequenza sono tutte bianche.
 *           - per ogni i=1,..,k la casella ci è subito alla destra di ci-1 oppure subito sotto a ci-1.
 * Il percorso c0,..,ck parte dalla casella c0 e arriva nella casella ck.
 * Inoltre chiamiamo completo un percorso che parte dalla casella in alto a sinistra e arriva nella casella in basso a destra.
 * Si vuole calcolare, data una griglia qualsiasi, quanti sono i percorsi completi.
 *
 ***************************
 * Domande:
 * Considerare la funzione f così definita per calcolare il numero di percorsi completi:
 * - data una casella c qualsiasi, f(c) è pari al numero di percorsi che partono nella casella c e finiscono nella casella in basso a destra.
 * 
 * 1. Se c è una casella nera, allora f(c) = 0 ===> non esistono percorsi perché parto da una casella nera che è vietata.
 * 2. Se c è la casella in basso a destra ed è bianca, allora f(c) = 1 ==> sono già in basso a destra, è l'unico cammino.
 * 3. Se c è una casella qualsiasi dell'ultima riga, si può scrivere una relazione di ricorrenza che fa riferimento alla casella a destra di c:
 *    - f(c) = f(c+1) dove con c+1 si intende la casella subito a destra  
 * 4. Se invece c è una casella qualsiasi dell'ultima colonna come si può scrivere f(c)?
 *    - f(c) = f(c+riga), posso solo andare in basso 
 * 5. Se c è una casella bianca non di bordo, si può scrivere una relazione di ricorrenza che fa riferimento alla casella sotto e alla casella a destra di c:
 *   - f(c) = f(c+1) + f(c+riga)
 * 6. Per quale casella c si ha che f(c) coincide il numero di cammini completi?
 *    Coincide quando c è la casella è quella in alto a sinistra.
 ******************************************************************
 * Progettazione-descrizione camminiCompleti:
 * La funzione `camminiCompleti` prende in input una griglia rettangolare rappresentata come una matrice di dimensioni `n` righe e `m` colonne. 
 * Ogni elemento della matrice può essere 0 o 1, dove 0 rappresenta una casella bianca e 1 rappresenta una casella nera. 
 * La casella in alto a sinistra è indicizzata come (0, 0) e la casella in basso a destra come (n-1, m-1).
 * L'algoritmo restituisce il numero di cammini completi dalla casella in alto a sinistra alla casella in basso a destra nella griglia.
 *
 * L'algoritmo utilizza un approccio di programmazione dinamica per calcolare il numero di cammini completi in modo efficiente, 
 * evitando di calcolare più volte gli stessi percorsi.
 * Passi:
 * 1. Inizializzazione:
 *     - Creare una matrice ausiliaria `F` di dimensioni `n` righe e `m` colonne, inizializzata con tutti gli elementi a zero. 
 *       La matrice `F` conterrà i valori intermedi dei cammini completi.
 * 2. Calcolo dei valori intermedi:
 *    - Iterare sulla griglia dalla casella in basso a destra fino alla casella in alto a sinistra.
 *    - Per ogni casella `c` considerata:
 *    - Se `c` è una casella nera o la casella in basso a destra, assegnare `F[c] = 0`, poiché non sono possibili cammini in queste caselle.
 *    - Se `c` è una casella dell'ultima riga, assegnare `F[c] = 1`, poiché esiste un solo cammino che arriva alla casella in basso a destra partendo da questa casella.
 *    - Se `c` è una casella dell'ultima colonna, assegnare `F[c] = F[c+riga]`, poiché non è possibile spostarsi lateralmente verso destra.
 *    - Se `c` è una casella bianca non di bordo, assegnare `F[c] = F[c+1] + F[c+riga]`, dove `c+1` rappresenta la casella a destra di `c` 
 *      e `c+riga` rappresenta la casella sotto `c`. Questo calcolo tiene conto dei cammini che possono arrivare a `c` dalla destra e dal basso.
 * 3. Restituzione del risultato:
 *    - Restituire il valore `F[0][0]`, che rappresenta il numero di cammini completi dalla casella in alto a sinistra alla casella in basso a destra.
 *
 * Attraverso questo approccio di programmazione dinamica, l'algoritmo evita di calcolare ripetutamente lo stesso percorso e calcola in modo efficiente 
 * il numero di cammini completi nella griglia data. La complessità dell'algoritmo è O(n * m), 
 * dove `n` è il numero di righe della griglia e 'm' il numero di colonne della griglia.
 *  
*/

func camminiCompleti(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])

	// Inizializza la matrice F
	F := make([][]int, rows)
	for i := range F {
		F[i] = make([]int, cols)
	}

	// Calcola il numero di cammini completi per ogni casella
	for i := rows - 1; i >= 0; i-- {
		for j := cols - 1; j >= 0; j-- {
			if grid[i][j] == 1 {
				// Casella nera, non ci sono cammini possibili
				F[i][j] = 0
			} else if i == rows-1 && j == cols-1 {
				// Casella in basso a destra
				F[i][j] = 1
			} else if i == rows-1 {
				// Casella dell'ultima riga
				F[i][j] = F[i][j+1]
			} else if j == cols-1 {
				// Casella dell'ultima colonna
				F[i][j] = F[i+1][j]
			} else {
				// Casella bianca non di bordo
				F[i][j] = F[i][j+1] + F[i+1][j]
			}
		}
	}

	// Restituisci il numero di cammini completi dalla casella in alto a sinistra
	return F[0][0]
}


func main() {
	var n, m int
	fmt.Scan(&n, &m)
	M := make([][]int, n)
	for i := 0; i < n; i++ {
		M[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scan(&M[i][j])
		}
	}
	fmt.Println(camminiCompleti(M))
}
