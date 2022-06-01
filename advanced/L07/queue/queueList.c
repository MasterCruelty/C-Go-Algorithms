#include <stdio.h>
#include <stdlib.h>
#include "queue.h"


//definizione delle struct necessarie per la lista concatenata
struct node{
	Item info;
	struct node *next;
};

typedef struct node *Node;

struct queue{
	Node front;
	Node rear;
	int n_element;
};

//creazione di una nuova coda allocando lo spazio necessario
void new_queue(Queue *q,int n){
	Queue new = malloc(sizeof(struct queue));
	if(new == NULL){
		printf("Errore allocazione memoria\n");
		exit(EXIT_FAILURE);
	}
	new -> front = NULL;
	new -> rear = NULL;
	new -> n_element = 0;
	*q = new;
}

//distruzione della coda mediante la distruzione di ogni nodo della lista concatenata
void destroy(Queue q){
	for(Node i = q -> front; i !=NULL;i = i -> next){
		printf("Distruggo il nodo %d\n",q -> front -> info);
		q -> front = q -> front -> next;
		free(q -> front);
	}
	free(q);
}

//funzione ausiliaria per l'allocazione di un nuovo nodo della lista concatenata
Node new_node(Item info){
	Node new = malloc(sizeof(struct node));
	if(new == NULL){
		printf("Errore allocazione memoria\n");
		exit(EXIT_FAILURE);
	}
	new -> info = info;
	new -> next = NULL;
	return new;
}

//nuovo elemento aggiunto alla coda
void enqueue(Queue q, Item info){
	Node new = new_node(info);
	if(is_empty_queue(q) != 0){
		q -> front = new;
		q -> rear = new;
	}
	else {
		q -> rear -> next = new;
		q -> rear = new;
	}
	q -> n_element++;
}

//rimozione di un elemento dalla coda
Item dequeue(Queue q){
	Node n = q -> front;
	Item val;
	if(is_empty_queue(q)!=0){
		printf("La coda risulta vuota.\n");
		exit(EXIT_FAILURE);
	}
	else {
		val = q -> front -> info;
		q -> front = q -> front -> next;
		q -> n_element--;
		if(q -> front == NULL)
			q -> rear = NULL;
		free(n);
		return val;
	}
}

//restituisce la testa della coda
Item frontValue(Queue q){
	return q -> front -> info;
}

//controllo se è vuota la coda
int is_empty_queue(Queue q){
	if(q -> front == NULL)
		return 1;
	else
		return 0;
}

//stampa del contenuto della coda
void print_queue(Queue q){
	if(is_empty_queue(q) != 0)
		printf("La coda risulta vuota\n");
	else {
		printf("Stampo il contenuto della lista:\n");
		for(Node i = q -> front; i != NULL; i = i -> next)
			printf("|%d|", i -> info);
	}
}
