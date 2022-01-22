#include <stdio.h>

typedef enum{False,True} bool;

bool isPrimo(int num);


int main(void){
    int n;
    printf("Inserisci un numero: ");
    scanf("%d",&n);
    if(isPrimo(n))
        printf("il numero è primo\n");
    else
	printf("Il numero non è primo\n");
    return 0;
}

bool isPrimo(int num){
    int check = 0;
    int i;
    for(i = 1;i <= num;i++){
        if(num % i == 0){
	    check++;
	}
	if(check > 2)
	    return False;
    }
    return True;
}
