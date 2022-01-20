#include <stdio.h>

//dato un numero n, inserisco da zero a n una sequenza ordinata di interi tranne uno. Il programma deve stampare il valore mancante.
int main(void){
    int n;
    int temp;
    int cont = 0;
    int prev = -1;
    int mancante = -1;
    printf("inserire la quantità n di numeri da inserire nell'insieme ordinato {0,1,2,...,n}: ");
    scanf("%d",&n);
    //finchè non inserico n numeri continuo a leggerli
    while(cont < n && scanf("%d",&temp)){
	//a ogni inserimento controllo se la differenza con il precedente è diversa da uno, in quel caso in mezzo c'è il valore mancante.
        if((temp - prev) != 1){
            mancante = temp -1;
	    break;
	}
	prev = temp;
	cont++;
    }
    printf("Il valore mancante è: %d\n",mancante);
    return 0;
}
