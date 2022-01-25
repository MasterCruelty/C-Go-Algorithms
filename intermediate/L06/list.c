#include <stdlib.h>

//definizione di un nodo come struttura
struct node{
    int info;
    struct node *next;
};
//definizione con typedef per usare Node al posto di "struct node"
typedef struct node *Node;

//prototipi funzioni per gestione liste (L06)
void list_delete(int n, Node l);
Node list_search(int n, Node l);
void list_print(Node l);
void list_insert(int n,Node *l);
void list_printInv(Node l);
int *listToArray(Node l);
void list_destroy(Node l);
Node olist_insert(int n, Node l);
Node olist_search(int n, Node l);


int main(){
    Node first = NULL;
    Node l = NULL;
    list_insert(2,&first);
    list_insert(5,&first);
    list_insert(3,&first);
    list_insert(10,&first);
    list_insert(14,&first);
    list_insert(1,&first);
    printf("Lista con list_insert(Node l):\n");
    list_print(first);

    return 0;
}

//stampa degli elementi contenuti nella lista
void list_print(Node l){
    Node p;
    int n = 1;
    for(p=1;p!=NULL;p = p -> next){
	printf("--> %d ", p -> info);
	if(n == 20)
	    break;
	n++;
    }
    printf("\n");
}

//inserimento di un elemento in una lista
void list_insert(int n,Node *l){
    Node new = malloc(sizeof(struct node));
    new -> info = n;
    new -> next = *l;
    *l -> new;
}

//ricerca di un elemento in una lista
Node list_search(int n,Node l){
    while(l != NULL && l -> info != n)
	l = l -> next;
    return l;
}

//rimozione di un elemento dalla lista
void list_delete(int n,Node l){
    Node curr, prev;
    for(curr = l, prev = NULL;curr != NULL; prev = curr, curr = curr -> next){
        if(curr -> info == n)
	    break;
    }
    if(curr == NULL)
        return;
    if(prev == NULL)
	l = l -> next;
    else
        prev -> next = curr -> next;
    free(curr);
}

//stampa della lista al contrario ricorsivamente
void list_printInv(Node l){
   if(l -> next != NULL)
	list_printInv(l -> next);
   else
	printf("\n");
  printf("%d ",l -> info); 
}

//convertire da lista ad array di interi
int *listToArray(Node l){
    Node p;
    int *a, size = 5, i = 0;
    for(p = l; p != NULL; p = p -> next){
	if(i >= size){
	    size +=5;
	    a = realloc(a,size * sizeof(int));
	}
	a[i++] = p -> info;
    }
    return a;
}

//distruzione di una lista e liberazione della memoria occupata
void list_destroy(Node l){
    if(l -> next != NULL){
	list_destroy(l -> next);
    }
    l -> next = NULL;
    free(l);
}

//inserimento in una lista ordinata
Node olist_insert(int n, struct node *l) {
    struct node *curr, *prev;
    for(curr = l, prev = NULL; curr != NULL; prev = curr, curr = curr -> next){
	if(curr -> info > n){
	    printf("curr -> info: %d\nInserisco quindi %d\n",curr -> info,n);
	    Node new = malloc(sizeof(struct node));
	    new -> info = n;
	    new -> next = curr;
	    prev -> next = new;
	    break;
	}
	else{
	    printf("curr -> info: %d\n",curr -> info);
	}
	if(curr == NULL)
	    return;
	if(prev == NULL)
	    l = l -> next
    }
}

//ricerca in una lista ordinata
Node olist_search(int n,Node l){
    struct node *curr, *prev;
    for(curr = l, prev = NULL; curr != NULL; prev = curr, curr = curr -> next){
	if(curr -> info > n)
	    break;
	if(curr -> info == n){
	    printf("Ho trovato l'elemento %d\n",n);
	    return curr;
	}
    }
    printf("Non ho trovato l'elemento %d\n",n);
    return NULL;
}

