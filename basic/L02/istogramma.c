#include <stdio.h>
#include <ctype.h>
#define N 26
int main(void){
    char c;
    int cont[N];
    //riempio di zeri l'array
    for(int i = 0; i < 26; i++){
        cont[i] = 0;
    }
    printf("Inserisci la frase: ");
    //mappo il contatore per incrementare sempre nella posizione della lettera che incontra
    while((c = getchar()) != '.'){
        c = tolower(c);
	cont[(int)c-97]++;
    }
    //stampo le lettere e le loro occorrenze
    for(int i = 0;i<26;i++){
        if(cont[i] >0){
	    printf("%c: ",i+65);
	    for(int j = 0;j < cont[i];j++){
                printf("*");
	    }
	    printf("\n");
	}
    }
    return 0;
}
