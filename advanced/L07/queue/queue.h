#ifndef QUEUE_H
#define QUEUE_H

typedef struct queue *Queue;
typedef int Item;

void new_queue(Queue *q,int n);
void destroy(Queue q);
//aggiunge un elemento in coda
void enqueue(Queue q, Item i);

//elimina un elemento dalla testa
Item dequeue(Queue q);

//restituisce l'elemento in testa
Item frontValue(Queue q);

//controllo se la coda è vuota
int is_empty_queue(Queue q);


//stampa il contenuto della coda
void print_queue(Queue q);

#endif
