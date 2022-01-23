#include <stdio.h>

//funzione che dati due puntatori a variabili intere, scambia il valore tra i due puntatori.
void scambia(int *p,int *q);

int main(void){
    int x,y;
    int *p;
    int *q;
    p = &x;
    q = &y;
    printf("Inserisci due interi: ");
    scanf("%d %d",&x,&y);
    scambia(p,q);
    printf("Interi scambiati: %d %d\n",*p,*q);
    return 0;
}

void scambia(int *p, int *q){
    int n = *p;
    int m = *q;
    *q = n;
    *p = m;
}
