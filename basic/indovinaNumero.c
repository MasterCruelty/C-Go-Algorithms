#include <stdio.h>
#define MAX 1000

//il programma fa uso della ricerca dicotomica per indovinare il numero sempre in 10 domande.
int main(void){
    int min = 0, max = MAX, n;
    char risposta;

    for( ; ;){
        n = min + (max - min) / 2;
	printf("il numero è %d? ", n);
	scanf(" %c",&risposta);
	switch(risposta){
	    case '<': max = n - 1; break;
	    case '>': min = n + 1; break;
	    case '=': return 0;
        }
    }
    return 0;
}
