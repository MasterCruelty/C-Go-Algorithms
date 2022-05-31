typedef struct{
    Node head;
} *List;

confrontare questo tipo di lista con list_with_tail.

Per quali delle seguenti operazioni si ha un tempo migliore con list_with_tail piuttosto che con List?

A)Restituisci 1 se la lista contiene un dato elemento [NO]
B) cancella l'ultimo elemento della lista [SI] ===> Poichè avendo un nodo puntatore alla coda della lista è veloce trovare l'ultimo elemento.
C) aggiungi un elemtno all'inizio della lista [NO]
D) aggiungi un elemtno alla fine della lista [SI] ===> per lo stesso motivo dell'opzione B non mi è necessario scorrere tutta la lista per inserire in coda alla lista.

Lista semplice o bidirezionale??
Individuare quale dei seguenti esperimenti consente di stabilire se l'implementazione è una lista semplice o una bidirezionale sapendo solo la seguente interfaccia:
* size() restituisce il numero di elementi nella lista
* conatins(e) restituisce 1 se "e" è nella lista
* removeAtStart() cancella l'elemento all'inizio della lista
* removeAtEnd() cancella l'elemento alla fine della lista
* AddAtStart(e) inserisce l'elemento "e" all'inizio della lista
* addAtEnd(e) inserisce l'elemento "e" alla fine della lista

A) Realizzo una mia implementazione della lista semplice e confronto i tempi di esecuzione su tutte le funzioni. 
   Se i tempi della mia implementazione coincidono con quelli della implementazione data per tutte le funzioni, 
   allora si tratta di una lista semplice, altrimenti di una lista bidirezionale.
   [VALIDO] se ho costruito in modo corretto la mia lista semplice, se i tempi coincidono è ragionevole pensare che anche 
            l'implementazione sconosciuta sia una lista semplice.
   
B) Eseguo N operazioni addAtEnd seguite da N operazioni addAtStart. Se le operazioni addAtEnd ci impiegano
   molto di più di quelle addAtStart, allora si tratta di una lista semplice, altrimenti una lista bidirezionale.
   [VALIDO] perché in una lista semplice l'inserimento in testa cosa molto meno dell'inserimento in coda.

C) Eseguo N operazioni addAtEnd seguite da N operazioni removeAtEnd. Se le operazioni removeAtEnd ci impiegano
   molto di più di quelle addAtEnd, allora si tratta di una lista semplice, altrimenti di una lista bidirezionale.
   [NON VALIDO] perché in una lista semplice sia le addAtEnd che le removeAtEnd hanno tempo uguale.

D) Creo due istanze della lista e confronto il tempo di esecuzione della prima rispetto alla seconda, per tutte le operazioni.
   Se i tempi sono simili per tutte le operazioni, allora si tratta di una lista semplice, altrimenti di una lista bidirezionale.
   [NON VALIDO] se creo due istanze della lista i tempi saranno per forza simili essendo due liste dello stesso tipo.

E) Nessuno degli esperimenti precedenti mi consente di rispondere alla domanda. Bisogna esaminare il codice per
   vedere se l’implementazione usa un riferimento prev al nodo precedente.
   [NON VALIDO] almeno uno degli esperimenti precedenti può rispondere alla domanda
   
---------------------------------------------------------------------------------------------------------------------------------------------
Esercizio 1.5 ANNULLA L'ULTIMA OPERAZIONE

Bisogna implementare una funzione "annulla" undo, che annulla l'ultima operazione svolta da un applicativo: si vuole salvare le azioni dell'utente in modo da poterle
annullare nell'ordine inverso. Se l'utente esegue le azioni a,b e c, la funzionalità "annulla" andrebbe ad annullare prima c, poi b e infine a.

La struttura dati usata è una lista concatenata semplice. Tra le seguenti opzioni, come usereste la lista per memorizzare le azioni dell'utente, 
al fine di avere le migliori prestazioni per la feature "annulla"??

A) Aggiungo alla fine e tolgo dall’inizio
B) Aggiungo all’inizio e tolgo dalla fine
C) Aggiungo e tolgo dalla fine
D) Aggiungo e tolgo dall’inizio.

Scelgo l'opzione C perché "annulla" funziona in LIFO annullando dall'ultimo al primo.
