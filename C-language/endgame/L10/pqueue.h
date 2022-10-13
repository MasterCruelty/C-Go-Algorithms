#ifndef PQUEUE_H
#define PQUEUE_H

#define N 15

typedef struct pqueue * Pqueue;
typedef int Item;
typedef Item *Heap;

Pqueue pqueue_new(int n);
void pqueue_destroy(Pqueue pq);
int pqueue_length(Pqueue pq);
void pqueue_insert(Pqueue pq,Item k);
Item pqueue_extractmin(Pqueue pq);
Item pqueue_min(Pqueue pq);

void heapify_up(Heap h,int i);
void heapify_down(Heap h,int i,int n);

int father(int i);
void swap(Heap h,int i,int j);
void heap_printassummary(Pqueue pq,int i);
void heap_printnode(Heap h, int j);
void heap_printarray(Pqueue pq, int i);

#endif
