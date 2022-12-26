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
 * Cosa rappresentano i nodi? ==> I nodi sono rappresentati dalle gallerie.
 * Cosa rappresentano gli archi? ==> Gli archi sono rappresentati dagli svincoli.
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
 * Uso della struttura dati grafo con rappresentazione con liste di adiacenza
 * input: N M S H
 * per ogni M input: A B L
 * A e B indicano lo svincolo(arco) e L il grado di luminosità
 * riutilizzo del sorgente graph_list.go in advanced/L08_grafi/rappresentazione_grafi/graph_list.go
*/
