#include <stdio.h>


typedef int Item;
typedef struct bit_node *Bit_node;

//definizione di struttura di un nodo di un albero
struct bit_node{
    Item item;
    struct bit_node *l,*r;
};

int cmp(int x, int y);
Item bist_search(Bit_node p,Key k);
Item bist_search_it(Bit_node p,Key k);


int cmp(int x,int y){
    if(x == y)
        return 0;
    if(x < y)
        return -1;
    if(x > y)
        return 1;
}

//versione ricorsiva per la ricerca di un nodo in un albero binario di ricerca
Item bist_search(Bit_node p,Key k){
    if(!p)
        return NULLitem;
    int res = cmp(k,key(p -> item));
    printf("ho confrontato k %d con p -> item %d e ho trovato %d\n", k, key(p -> item);
    if(res == 0)
        return p -> item;
    if(res < 0)
        return bist_search(p -> l,k);
    return bist_search(p -> r,k);
}
 
           
//versione iterativa per la ricerca di un nodo in un albero binario di ricerca
Item bist_search_it(Bit_node p,Key k){
    int res;
    while(p && ( res = cmp(k,key(p -> item))) != 0){
        p = res < 0 ? p -> l : p -> r; 
    }
    if(p == NULL)
        return NULLitem;
    else
        return p -> item;
}
