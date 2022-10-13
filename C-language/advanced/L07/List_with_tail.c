#include <stdlib.h>

//list_with_tail indica una lista concatenata di interi dotata di due riferimenti al primo e all'ultimo elemento della lista.
//Quando la lista è vuota sia head che tail sono NULL. new_node alloca lo spazio per un nuovo nodo e inizializza il valore con l'argomento passato.
struct node {
    int info;
    struct node *next;
};

typedef struct node *Node;

typedef struct {
    Node head;
    Node tail;
} *List_with_tail;

Node new_node(inte e){
    Node new = malloc(sizeof(struct node));
    new -> info = e;
    new -> next = NULL;
    return new;
}

void addAtEnd(List_with_tail l, int e){
    if(l -> tail == NULL) {
        l -> tail = new_node(e);
        l -> head = l -> tail;
    }
    else{
        /*soluzioni sbagliate con commenti
        //In questo caso sto assegnando al campo next direttamente l'intero e senza allocare lo spazio per il nodo, risulta ovviamente non funzionante e non corretto.
        l -> tail -> next = e;
        l -> tail = e;
        ----------------------------
        //In questo caso alloco correttamente lo spazio per il nuovo nodo e lo assegno in coda alla lista, ma non collego la vecchia coda a quella nuova, dunque il nuovo nodo
        //risulta staccato dalla lista.
        Node temp = new_node(e);
        l -> tail -> next = temp;
        ----------------------------
        //In questo caso alloco correttamente lo spazio per il nuovo nodo e lo assegno alla coda della lista sovrascrivendo la coda già esistente senza salvarla.
        Node temp = new_node(e);
        l -> tail = temp;
        ----------------------------
        //In questo caso la soluzione risulta funzionante ma non è corretta dal punto di vista delle performance.
        //Sto scorrendo inutilmente tutta la lista prima di trovare la coda ma io avendo un puntatore dedicato alla coda posso evitare di scorrerla.
        Node temp = -> head;
        while(temp -> next != NULL){
            temp = temp -> next;
        }
        temp -> next = new_node(e);
        */
      
        //SOLUZIONE CORRETTA
        //Alloco lo spazio per il nuovo nodo assegnandolo al next della coda e poi collego la coda vecchia con la coda nuova.
        l -> tail -> next = new_node(e);
        l -> tail = l -> tail -> next;
    }
}
