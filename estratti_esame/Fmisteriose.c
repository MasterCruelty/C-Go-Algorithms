#include <stdio.h>
#include <stdlib.h>

struct node{
    struct node *next;
    int val;
};

typedef struct node *Node;

Node newNode(int val,Node p);
int f1 (Node p);
Node f2 (Node p);
Node f3(Node p, Node q);

//1)Quale valore restituisce la funzione f1 se invocata sul nodo x con val 1?
//restituisce 1 perchè il suo next è NULL
//
//
//2)Se invochiamo f1 sul nodo head cosa succede?
//La funzione viene chiamata su entrambi i nodi e restituisce 1
//3)Cosa restituisce f2(head)?
//viene restituito il nodo head
//4) Se head è il puntatore di una lista qualsiasi formata da almeno 2 nodi, allora la chiamata f2(head) restituisce un puntatore a...
//>se stesso
int main() {
    int n;
    Node x = newNode(1,NULL);
    Node head = newNode(2,x);
    printf("Calcolo f1 su nodo x con valore 1 e puntando a next NULL\n");
    n = f1(x);
    printf("Risultato: %d\n",n);
    printf("Calcolo f1 su nodo head con valore 2 e puntando a next Nodo x\n");
    n = f1(head);
    printf("Risultato: %d\n",n);
    printf("Calcolo f2 su nodo head con valore 2 e puntando a next Nodo x\n");
    Node s = f2(head);
    return 0;
}

Node newNode(int val, Node p){
    Node new = malloc(sizeof(struct node));
    new -> val = val;
    new -> next = p;
    return new;
}

int f1(Node p){
    getchar();
    if(p -> next == NULL)
	return 1;
    else if(p-> val < p -> next -> val)
	return 0;
    else
	return f1(p-> next);
}

Node f2(Node p){
    getchar();
    if(p -> next == NULL)
	return p;
    else
	return f3(p -> next,p);
}

Node f3(Node p,Node q){
    Node temp;
    if(p -> val < q -> val)
	temp = p;
    else
	temp = q;
    if(p -> next == NULL)
	return temp;
    else
	return f3(p -> next,temp);
}
