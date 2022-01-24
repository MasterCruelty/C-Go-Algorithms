#include <stdio.h>
#include <string.h>
#define N 10
//definizione del nuovo tipo String
typedef char *String;
//versione funzione che restituisce l'indice come intero
int smallest_word_index(String s[], int n);
//versione funzione che restituisce proprio l'indirizzo come puntatore
String *smallest_word_address(String s[], int n);

int main(void){
    String dict[] = {"ciao", "aaa","mondo","funzione","come","funziona"};
    int lun = 6;
    printf("smallest con index: %s\n", dict[smallest_word_index(dict,lun)]);
    printf("smallest con puntatore: %s\n", *smallest_word_address(dict,lun));
    return 0;
}

String *smallest_word_address(String s[],int n){
    String *min = s;
    //uso il puntatore per scorrere il vettore di stringhe
    for(String *p = s + 1; p < s + n;p++){
        if(strcmp(*p,*min) <0)
	    min = p;
    }
    //restituisco il puntatore all'indirizzo di memoria dove si trova la parola minima nel vettore
    return min;
}

int smallest_word_index(String s[], int n){
    int min = 0;
    //uso un indice intero per scorrere il vettore di stringhe
    for(int i = 1;i < n; i++){
        if(strcmp(s[i],s[min]) < 0)
	    min = i;
    }
    //restituisco la posizione del vettore in cui vi è la parola minima
    return min;
}
