#include <stdio.h>
#include <stdlib.h>
#include <ctype.h>
#include <string.h>

#define N 100
//confrontare con treVolteBUGGATA.c per vedere cosa è stato corretto
void treVolte(char *a[], int n){
    char **p,*q;
    int conta = 0;
    q = *a; //il puntatore *q non era inizializzato il che poteva portare a errori potendoci essere qualsiasi cosa dentro *q

    for(p = a;p<a+n;p++){
        int contaE = 0;
	char *c;
	c = *p;

	while(*c){
	    if(*c == 'e')
	        contaE++;
	    if(contaE == 3){
		conta++;
		break; //chiudo il ciclo forzatamente nel momento in cui ho la certezza che sia una parola con tre caratteri 'e'
	    }
	    c++;
	}
	if(conta == 3){
	    //tolto tutte le strcpy e al posto di p ho messo *p per prendere il contenuto di **p
	    q = *p; 
	    *p = a[0];
	    a[0] = q;
	    return;
	}
    }
}

int main(int argc,char **argv){
    treVolte(argv + 1,argc);
    //aggiunto +1 a argv per non fare stampare anche il nome del programma ma solo le parole
    for(char** p= argv+1; p < argv + argc; p++)
	printf("%s ",*p);
    printf("\n");
    return 0;
}
