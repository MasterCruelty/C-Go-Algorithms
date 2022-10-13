#include <stdio.h>
#include <stdlib.h>
#include <time.h>

#define N 15

typedef int Item;
typedef struct bit_node *Bit_node;

//definizione di struttura di un nodo di un albero
struct bit_node{
    Item item;
    struct bit_node *l,*r;
};

//prototipi funzioni per la gestione di un albero
void bit_inorder(Bit_node p,int j);
void bit_postorder(Bit_node p,int j);
void bit_preorder(Bit_node p,int j);
void bit_printassummary(Bit_node p,int j);
void bit_printnode(Bit_node p,int j);
void bit_destroy(Bit_node p);
Bit_node bit_new(Item item);
Bit_node bit_arr2tree(Item a[], int size,int j);
Bit_node bit_left(Bit_node p);
Bit_node bit_right(Bit_node p);
Item bit_item(Bit_node p);

//Item bist_search_it(Bit_node p,int k);

int main(void){
    int tree[N];
    int i;
    //genero seme 
    srand(time(0));
    //genero 15 numeri casuali tra 1 e 100
    for(i = 0;i<N;i++)
	tree[i] = 1 + rand()%100;

    //costruisco l'albero a partire dall'array
    Bit_node radice = bit_arr2tree(tree,N,0);
    printf("Stampa dell'albero completo\n\n");
    bit_printassummary(radice,0);
    printf("VISITA INORDER\n\n");
    bit_inorder(radice,0);
    printf("VISITA POSTORDER\n\n");
    bit_postorder(radice,0);
    return 0;
}


//conversione da array ad albero
//la radice è contenuta in a[0]
//il figlio sinistro è a[2*i+1]
//il figlio destro è a[2*i+2]
Bit_node bit_arr2tree(Item a[], int size,int j){
    if(j >= size|| a[j] == -1){
	return NULL;
    }
    Bit_node root = bit_new(a[j]);
    root -> l = bit_arr2tree(a,size,(j+j+1));
    root -> r = bit_arr2tree(a,size,(j+j+2));
    return root;
}

//creazione nuovo nodo
Bit_node bit_new(Item item){
    Bit_node root = malloc(sizeof(struct bit_node));
    root -> item = item;
    root -> l = malloc(sizeof(struct bit_node));
    root -> r = malloc(sizeof(struct bit_node));
    return root;
}

//Elimina il nodo p dall'albero e di conseguenza anche i suoi figli(?)
void bit_destroy(Bit_node p){
	free(p -> l);
	free(p -> r);
	free(p);
}

//Restituisce il valore associato al nodo p
Item bit_item(Bit_node p){
    return p -> item;
}
//Restituisce il figlio sinistro
Bit_node bit_left(Bit_node p){
    return p -> l;	
}

//Restituisce il figlio destro
Bit_node bit_right(Bit_node p){
    return p -> r;	
}

//visita dell'albero in inorder(prima il sottoalbero di sx, poi la radice, infine il sottoalbero di dx)
void bit_inorder(Bit_node p,int j){
    if(p){
	if(p -> l || p -> r)
	    bit_inorder(p -> l,j+1);
	bit_printnode(p,j);
	if(p -> l || p -> r)
	    bit_inorder(p -> r,j+1);
    }
    else{
	for(int i = 0;i<j;i++)
	    printf("  ");
	printf("*X\n");
    }
}

//visita dell'albero in ordine anticipato(prima la radice, poi sottoalbero sx e poi sottoalbero dx)
void bit_preorder(Bit_node p,int j){
    if(p){
	bit_printnode(p,j);
	if(p -> l || p -> r){
	    bit_preorder(p -> l,j+1);
	    bit_preorder(p -> r,j+1);
	}
    }
    else{
	for(int i = 0;i<j;i++)
	    printf("  ");
	printf("*X\n");
    }
	
}

//visita dell'albero in postorder(prima sotto albero sx, poi dx e infine la radice)
void bit_postorder(Bit_node p,int j){
    if(p){
	if(p -> l || p -> r){
	    bit_postorder(p -> l,j+1);
	    bit_postorder(p -> r,j+1);
	}
	bit_printnode(p,j);
    }
    else{
	for(int i = 0;i<j;i++)
	    printf("  ");
	printf("*X\n");
    }
    
}

//stampa a schermo di tutto l'albero
void bit_printassummary(Bit_node p,int j){
    if(p){
	bit_printnode(p,j);
	if(p -> l||p-> r){
	    bit_printassummary(p->l,j+1);
	    bit_printassummary(p->r,j+1);
	}
    }
    else{
	for(int i = 0;i<j;i++)
	    printf("  ");
	printf("*X\n");
    }
}

//stampa a schermo del nodo richiesto
void bit_printnode(Bit_node p, int j){
    for(int i = 0;i<j;i++)
	printf("  ");
    printf("*%d \n",p -> item);
}

/*
Item bist_search_it(Bit_node p,int k){
    int res;
    while(p &&(res= cmp(k,key(p->item))) != 0)
	p = res < 0 ? p -> : p -> r;
    if(p == NULL)
	return NULLitem;
    else
	return p -> item;
}*/
