#include <stdio.h>
#include <stdlib.h>
#include "stack.h"

struct node{
    int val;
    struct node *next;
};

typedef struct node *Node;

typedef struct{
    Node head;
    int cont;
}Set;

Set *s;

void create_stack(void){
    s = malloc(sizeof(Set));
    s -> head = NULL;
    s -> cont = 0;
}

void make_empty(void){
    if(s -> head == NULL)
	return;
    else{
	s -> head = s -> head -> next;
	make_empty();
	s -> cont--;
	free(s -> head);
    }
}

int is_empty(void){
    if(s -> head == NULL)
	return 1;
    else
	return 0;
}

int pop(void){
    if(!is_empty()){
	Node p = s -> head;
	int n = s -> head -> val;
	s -> cont--;
	s -> head = s -> head -> next;
	free(p);
	print_stack();
	return n;
    }
    else
	exit(EXIT_FAILURE);
}

int top(void){
    if(!is_empty())
	return s -> head -> val;
    else
	exit(EXIT_FAILURE);
}


void push(int n){
    Node new = malloc(sizeof(struct node));
    new -> val = n;
    new -> next = s -> head;
    s -> head = new;
    s -> cont +=1;
    print_stack();
}

void print_stack(){
    Node p;
    int n = 1;
    for(p = s -> head; p != NULL; p = p -> next){
	printf("%d ",p -> val);
	n++;
    }
    printf("\n");
}
