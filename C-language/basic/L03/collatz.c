#include <stdio.h>

void Collatz(int n);

int main(void){
    int num;
    printf("Inserisci un numero per calcolarne la sequenza di Collatz: ");
    scanf("%d",&num);
    printf("\n");
    Collatz(num);
    return 0;
}

void Collatz(int n){
    int i = 1;
    printf("Numero: %d\n",n);
    printf("%d ",n);
    while(n != 1){
        if(n % 2 == 0){
            n /= 2;
	    printf("%d ",n);
	}
	else{
            n *= 3;
	    n += 1;
	    printf("%d ",n);
	}
	i++;
    }
    printf("\nLunghezza: %d\n",i);
}
