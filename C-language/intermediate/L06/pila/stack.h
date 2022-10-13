#ifndef STACK_H
#define STACK_H


//creo lo stack
void create_stack(void);

//svuota la pila
void make_empty(void);

//restituisce 1 se la pila è vuota, 0 altrimenti
int is_empty(void);

//se la pila non è vuota, restituisce il top della pila; Altrimenti esce con messaggio di errore
int top(void);

//se la pila non è vuota, estrae il top della pila; altrimenti esce con messaggio di errore
int pop(void);

//se c'è spazio, aggiunge n alla pila; altrimenti esce con messaggio di errore
void push(int n);

//stampa il contenuto della pila, partendo dal top
void print_stack(void);

#endif
