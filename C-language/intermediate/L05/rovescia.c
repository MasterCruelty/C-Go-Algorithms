#include <stdlib.h>
#include <stdio.h>

#define L 2
#define L0 3

int *read_lin(int *num);

int main(void){
    int i;
    printf("Inserisci n interi terminati da uno zero per concludere: ");
    int *a = read_lin(&i);
    //read_lin ha modificato i assegnandogli il valore dell'ultima posizione quindi posso scorrerlo al contrario
    while(i-->0)
	printf("a[%d] = %d\n",i,a[i]);
    return 0;
}

int *read_lin(int *num){
    int *a, size = L0, i = 0, n;
    //alloco inizialmente 3 spazi di memoria per 3 int e associo tale allocazione al puntatore intero a
    a = malloc(size * sizeof(int));
    while(1){
	scanf("%d",&n);
	if(n==0)
	    break;
	if(i >= size){
	    //se supero la memoria allocata, allora aumento lo spazio allocato reallocando la stessa size + 2 spazi aggiuntivi
	    size += L;
	    a = realloc(a,size * sizeof(int));
	}
	//assegno alla posizione attuale l'n appena inserito
	a[i++] = n;
    }
    //aggiorno il puntatore passato come argomento con il valore ultimo assunto da i;
    *num = i;
    return a;
}
