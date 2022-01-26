#include <stdio.h>
#include <stdlib.h>
#include "stack.h"

#define N 100



int cont = 0;
int pila[N];

void create_stack(void){
    printf("stack creato\n");
}

void make_empty(void){
    cont = 0;
}

int is_empty(void){
    if(cont == 0)
	return 1;
    else
	return 0;
}


int top(void){
    if(!is_empty())
	return pila[cont-1];
    else
	exit(EXIT_FAILURE);
}

int pop(void){
    if(!is_empty()){
	int n = pila[cont-1];
	cont--;
	print_stack();
	return n;
    }
    else
	exit(EXIT_FAILURE);
}

void push(int n){
    pila[cont++] = n;
    print_stack();
}

void print_stack(void){
    for(int i = 0;i< cont;i++)
	printf("%d ",pila[i]);
    printf("\n");
}
