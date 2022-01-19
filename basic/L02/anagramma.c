#include <stdio.h>
#include <ctype.h>
#define N 26
int main(void){
    char c;
    int cont1[N];
    int cont2[N];
    int anagramma = 0;
    //riempio di zeri l'array
    for(int i = 0; i < 26; i++){
        cont1[i] = 0;
        cont2[i] = 0;
    }
    printf("Inserisci la prima parola: ");
    //mappo il contatore per incrementare sempre nella posizione della lettera che incontra
    while((c = getchar()) != '.'){
        c = tolower(c);
	cont1[(int)c-97]++;
    }
    printf("Inserisci la seconda parola: ");
    while((c = getchar()) != '.'){
        c = tolower(c);
	cont2[(int)c-97]++;
    }
    //stampo le lettere e le loro occorrenze
    for(int i = 0;i<N;i++){
        if(cont1[i] == cont2[i]){
            anagramma = 1;
	}
	else{
	    anagramma = 0;
	    break;
	}
    }
    if(anagramma == 1)
        printf("le due parole sono anagrammi\n");
    else
        printf("le due parole non sono anagrammi\n");
    return 0;
}
