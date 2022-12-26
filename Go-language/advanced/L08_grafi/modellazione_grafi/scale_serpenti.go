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
