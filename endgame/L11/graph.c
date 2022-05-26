#include <stdlib.h>
#include <stdio.h>

//nodo di lista di adiacenza
struct node{
    int v;
    struct node *next;
};

//struttura per ogni grafo
//Considerato come grafo non orientato
struct graph{
    int V;            //numero di nodi
    int E;            //numero di archi
    struct node **A;  //array di liste di adiacenza
};

struct graph *createGraph(int nv,int ne);
void readGraph(struct graph *g, FILE *fp);
struct node *vertexinsert(struct node *p, int k);
void depthFirstSearch1(struct graph *g, int i, int *aux);
void depthFirstSearch(struct graph *g);
void breadthFirstSearch1(struct graph *g,int i, int *aux);
void breadthFirstSearch(struct graph *g);

//main per testare la struttura del grafo
int main(void){
    //WIP
    return 0;
}

//Creazione di un grafo con numero di archi e vertici. Alloca spazio per ogni lista di adiacenza
struct graph *createGraph(int nv, int ne){
    struct graph *g = malloc(sizeof(struct graph));
    if(!g){
        fprintf(stderr,"Errore allocazione\n");
        exit(-1);
    }
    g -> E = ne;
    g -> V = nv;
    if(!(g -> A = calloc(nv,sizeof(struct node *)))) {
        fprintf(stderr,"Errore allocazione\n");
        exit(-2);
    }
    return g;
}

//lettura da file per gli archi per popolare le liste di adiacenza. Nel file ogni riga contiene un unico arco rappresentato come coppia di vertici.
void readGraph(struct graph *g, FILE *fp){
    int i,v1,v2;
    for(i = 0;i< g -> E;i++){
        fscanf(fp,"%d %d",&v1,&v2);
        g->A[v1-1] = vertexinsert(g->A[v1-1],v2);
        g->A[v2-1] = vertexinsert(g->A[v2-1],v1);
    }
}

//Funzione per inserimento vertici nelle liste di adiacenza(Prettamente identico a un inserimento in testa per una lista concatenata).
//per ogni arco è necessario chiamare questa funzione due volte su due liste diverse dell'array g -> A.
//Se ho un arco (v,w) in E, inserisco v nella lista associata a w e inserisco w nella lista associata a v.
struct node *vertexinsert(struct node *p, int k){
    struct node *q = malloc(sizeof(struct node));
    if(!q){
        fprintf(stderr,"Errore allocazione\n");
        exit(-3);
    }
    q -> v = k;
    q -> next = p;
    return q;
}

//funzione ricorsiva che usa una array di appoggio per memorizzare quando un vertice è già stato incontrato.
//questa funzione richiama ricorsivamente se stessa finché non ha visitato tutta la componente.
void depthFirstSearch1(struct graph *g, int i, int *aux){
    struct node *t;
    aux[i] = 1;
    for(t= g -> A[i];t;t = t -> next){
        if(!aux[t -> v-1]){
            printf("%d,",t -> v);
            depthFirstSearch1(g,t -> v-1,aux);
        }
    }
}

//quando viene richiamata la funzione ricorsiva qua sopra si entra in una nuova componente connessa.
//ogni arco viene esaminato due volte mentre la lista di adiacenza di ogni vertice è scandita una volta sola(per via del ciclo sulla lista di adiacenza della funzione ricorsiva).
//questa visita ha un costo di O(|V| + |E|)
//questo tipo di visita permette di individuare le componenti connesse di un grafo
void depthFirstSearch(struct graph *g) {
    int i, *aux = calloc(g -> V,sizeof(int));
    if(!aux) {
        fprintf(stderr,"Errore allocazione\n");
        exit(-4);
    }
    for(i = 0;i < g -> V;i++){
        if(!aux[i]){
            printf("\n%d,",i+1);
            depthFirstSearch1(g,i,aux);
        }
    }
    free(aux);
}


//implementazione visita in ampiezza facendo uso di una coda
void breadthFirstSearch1(struct graph *g,int i, int *aux){
    struct node *t;
    intqueue *q = createqueue();
    enqueue(q,i);
    while(!emptyq(q)){
        i = dequeue(q);
        aux[i] = 1;
        for(t = g -> A[i]; t;t = t -> next){
            if(!aux[t -> v -1]){
                enqueue(q,t -> v-1);
                printf("%d,",t -> v);
                aux[t -> v -1] = 1;
            }
        }
    }
    destroyqueue(q);
}

//quando viene chiamata la funzione qui di sopra si entra in una nuova componente connessa.
//la funzione chiamata usa una coda per memorizzare da quali vertici riprendere la visita quando la lista di adiacenza del vertice corrente è stata tutte esplorata.
//ogni lista è visitata una volta, ogni arco due volte.
//Questo tipo di visita così implementata con liste di adiacenza ha un costo di O(|V| + |E)
void breadthFirstSearch(struct graph *g){
    int i, *aux = calloc(g -> V, sizeof(int));
    if(!aux){
        fprintf(stderr,"Errore allocazione\n");
        exit(-4);
    }
    for(i = 0; i < g -> V; i ++){
        if(!aux[i]) {
            printf("\n%d,",i+1);
            breadthFirstSearch1(g,i,aux);
        }
    }
    free(aux);
}

