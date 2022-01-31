#include <stdio.h>
#include <stdlib.h>
#include "pqueue.h"


struct pqueue{
    Heap h;
    int size,cont;
};


int main(void){
    return 0;
}

//creazione coda di n item
Pqueue pqueue_new(int n){
    Pqueue pq = malloc(sizeof(Pqueue));
    pq -> h = malloc(sizeof(Heap));
    pq -> cont = 0;
    pq -> size = n;
    return pq;
}

//distruzione coda, liberazione memoria
void pqueue_destroy(Pqueue pq){
    free(pq -> h);
    free(pq);
}

//lunghezza coda
int pqueue_length(Pqueue pq){
    return pq -> cont;
}

//Inserisce k nella coda
void pqueue_insert(Pqueue pq,Item k){
    if(pq -> cont < pq -> size -1){
	pq -> cont +=1;
	pq -> h[pq -> cont] = k;
	heapify_up(pq -> h,pq -> cont);
    }
}

//Rimuove k dalla coda
void pqueue_remove(Pqueue pq, Item k){
    swap(pq-> h,1,pq->cont);
    pq -> cont -=1;
    heapify_down(pq -> h,1,pq -> cont);
}

//estrae dalla coda item con chiave minima
Item pqueue_extractmin(Pqueue pq){
    Item min = pqueue_min(pq);
    pqueue_remove(pq,1);
    return min;
}

//item con chiave minima nella coda
Item pqueue_min(Pqueue pq){
    return pq -> h[1];
}

void heapify_up(Heap h,int i){
    if(i > 1){
	int j = father(i);
	if(h[i] < h[j]){
	    swap(h,i,j);
	    heapify_up(h,j);
	}
    }
}

void heapify_down(Heap h,int i, int n){
    if(2* i <= n){
	int j;
	if(2 * i == n){
	    j = 2* i;
	}
	else{
	    j = h[2*i] < h[2*i+1] ? 2*i : 2*i+1;
	    if(h[2*i]<h[2*i+1]){
		swap(h,i,j);
		heapify_down(h,j,n);
	    }
	}
    }
}

int father(int i){
    return i/2;
}

void swap(Heap h,int i,int j){
    Item temp;
    temp = h[i];
    h[i] = h[j];
    h[j] = temp;
}
