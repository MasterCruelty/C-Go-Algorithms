package main

/*
 * Assumendo che gli N dipendenti di Algorè siano indicati con un numero progressivo da 0 a N-1.
 * Modellazione:
 * La struttura dati scelta è di tipo albero. Gli n dipendenti possono essere rappresentati come i nodi dell'albero.
 * Le relazioni tra di essi si rappresentano attraverso il legame di parentela padre/figlio tra i nodi.
 * Ad esempio il subordinato di un dipendente, è un suo nodo figlio, mentre il supervisore di un dipendente, è un suo nodo padre.
 * La radice è il dipendente di massimo livello con tutti gli altri nodi figli come subordinati.
 * 
 * 1. Dato un certo dipendente, stampare l'elenco dei suoi subordinati.
 *    - Dato un certo nodo all'interno dell'albero, stampare l'elenco dei suoi figli fino a raggiungere le foglie.
 * 2. Contare quanti sono i dipendenti che non hanno alcun subordinato.
 *    - Contare il numero di nodi all'interno dell'albero che non hanno alcun figlio, ovvero il numero di nodi foglia.
 * 3. Dato un certo dipendente, individuare chi è il suo supervisore.
 *    - Dato un certo nodo all'interno dell'albero, individuare il nodo padre.
 * 4. Dato un certo dipendente, stampare la lista dei dipendenti che si trovano sopra di lui gerarchicamente, 
 *    partendo dal suo supervisore e risalendo la gerarchia fino a un dipendente di massimo livello.
 *    - Dato un certo nodo all'interno dell'albero, stampare la lista dei nodi padre fino a raggiungere la radice.
 * 5. Stampare l'elenco di tutti i dipendenti -non importa l'ordine-, indicando per ciascuno chi è il suo supervisore(tranne per il massimo livello)
 *    - Stampare l'elenco di tutti i nodi individuando per ciascuno il nodo padre ma escludendo la radice da questa stampa.
 * 6. Stampare l'elenco di tutti i dipendenti, in ordine di livello(prima quelli di livello massimo, poi tutti quelli subordinati a quelli di livello max)
 *    Non importa l'ordine tra i dipendenti di pari livello.
 *    - Essenzialmente è una visita in ampiezza dell'albero
 *
 * Descrivere come è opportuno implementare la struttura dati scelta.
 *
 * Per ciascun compito, progettare e implementare un algoritmo che lo svolge.
 *
 * Spiegate come modifichereste le risposte ai punti precedenti se i dipendenti fossero identificati da un nome invece che un numero progressivo.
*/
