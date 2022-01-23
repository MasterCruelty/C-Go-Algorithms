#include <stdio.h>
#define N 20
typedef enum{False,True} bool;

bool palindromo(char parola[]);
int len(const char *str);

int main(int argc,char *argv[]){
    int i;
    for(i = 1; i < argc; i++){
        if(palindromo(argv[i]))
	    printf("la parola %s è palindroma.\n",argv[i]);
	else
	    printf("la parola %s non è palindroma.\n",argv[i]);
    }
    return 0;
}
//funzione ausiliaria per calcolare lunghezza stringa, non serve alla fine in questo esercizio ma lascio per utilità
int len(const char *str){
    int n = 0;
    while(*str++)
        n++;
    return n;
}

bool palindromo(char parola[]){
    char *i;
    char *f;
    for(f = parola; f < parola + N;f++){
	    if(*f == '\0')
	        break;
    }
    i = parola - 1;
    while(--f >= ++i){
        if(*f == *i)
	    continue;
	else
	    return False;
    }
    return True;
}
