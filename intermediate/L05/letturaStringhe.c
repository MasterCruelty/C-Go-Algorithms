#include <stdio.h>
#include <stdlib.h>
#include <ctype.h>

//legge una riga terminata dal carattere c
char *read_line(char c);
//legge una parola con caratteri alfanumerici e la lettura si deve interrompere al primo carattere non alfanumerico
//se già il primo non lo è restituisce la stringa vuota
char *read_word(void);


int main(void){
    char *c,*ch;
    c = read_line('b');
    printf("risultato read_line: %c\n",*c);
    ch = read_word();
    printf("risultato read_word: %c\n",*ch);
    return 0;
}

char *read_line(char c){
    char *a,ch;
    int size = 5, i = 0;
    //allocazione di spazio per size numero di caratteri
    a = malloc(size *sizeof(char));
    while(1){
	scanf(" %c",&ch);
	if(ch == 'c'){
	   a[i] = '\0';
	   break;
	}
	if(i >= size){
	    size *=2;
	    //re-allocazione di spazio per size raddoppiata numero di caratteri
	    a = realloc(a,size * sizeof(char));
	}
	//assegno alla posizione i-esima del puntatore a il caratere appena inserito
	a[i++] = ch;
    }
    return a;
}

char *read_word(void){
    char *p,c;
    int n = 0;
    int size = 5;
    p = malloc(size * sizeof(char));
    while(1){
	scanf(" %c",&c);
	if(n >= size){
	    size *=2;
	    p = realloc(p,size * sizeof(char));
	}
	if(!isalpha(c)){
	    p[n] = '\0';
	    break;
	}
	p[n++] = c;
    }
    return p;
}
