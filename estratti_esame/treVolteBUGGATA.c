#include <stdio.h>
#include <stdlib.h>
#include <ctype.h>
#include <string.h>

#define N 100
//SORGENTE BACATO
//la funzione, date n stringhe da linea di comando, appena trova la terza stringa contenente almeno tre 'e' la sostituisce con la prima
//in treVolteFIXATA si trovano le correzioni commentate
void treVolte(char *a[], int n){
    char **p,*q;
    int conta = 0;

    for(p = a;p<a+n;p++){
        int contaE = 0;
	char *c;
	c = *p;

	while(*c){
	    if(*c == 'e')
	        contaE++;
	    if(contaE == 3){
		conta++;
	    }
	    c++;
	}
	if(conta == 3){
	    strcpy(q,p);
	    strcpy(p,a[0]);
	    strcpy(a[0],q);
	return;
	}
    }
}

int main(int argc,char **argv){
    treVolte(argv + 1,argc -1);

    for(char** p= argv; p < argv + argc; p++)
	printf("%s ",*p);
    printf("\n");
    return 0;
}
