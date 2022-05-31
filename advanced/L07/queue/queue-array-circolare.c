#include <stdio.h>
#include <stdlib.h>
#include "queue.h"

struct queue{
	int front;
	int rear;
	int n_element;
	int n_max;
	Item *array;
};

//crea una coda nuova
void new_queue(Queue *q,int n){
	Queue new = malloc(sizeof(struct queue));
	if(new == NULL){
		printf("Allocazione con errore");
		exit(EXIT_FAILURE);
	}
	new -> front = 0;
	new -> rear = -1;
	new -> n_element = 0;
	new -> n_max = n;
	new -> array = malloc(n*sizeof(int));

	*q = new;
}

void destroy(Queue q){
	free(q -> array);
	free(q);
}

//aggiunge un elemento in coda
void enqueue(Queue q, Item e){
	if(q -> n_element == q -> n_max){
		printf("coda piena.\n");
	}
	q -> rear++;
	if(q -> rear > (q -> n_max)-1){
		q -> rear = 0;
	}
	q -> array[q -> rear] = e; //inserito l'item passato come argomento
	q -> n_element++; //aumento il numero di elementi avendone inserito uno
	printf("elementi inseriti: %d", q -> n_element);
	printf("\tRear %d in posizione %d", q -> array[q -> rear], q -> rear);
	printf("\tFront % d in posizione %d\n", q -> array[q -> front], q -> front);
}


//rimuove e restituisce un elemento dal front
Item dequeue(Queue q){
	Item first_element;
	first_element = q -> array[q -> front];
	if(q -> front == (q -> n_max) -1)
		q -> front = 0;
	else
		q -> front++;
	q -> n_element--; //diminuisco il numero di elementi avendone rimosso uno
	return first_element;
}

//Restituisce l'elemento front
Item frontValue(Queue q){
	return q -> array[q -> front];
}

//controllo se la coda è vuota
int is_empty_queue(Queue q){
	if(q -> n_element == 0)
		return 1;
	else
		return 0;
}

//stampa il contenuto della coda
void print_queue(Queue q){
	if(!(is_empty_queue(q))){
		printf("\nLista di %d valori contenuti nella coda:\t", q -> n_element);
		for(int i = 0; i <=(q -> n_element -1); i++)
			printf("|%d|", q -> array[i]);
	}
	printf("\n");
}
