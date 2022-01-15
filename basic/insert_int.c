#include <stdio.h>

int main(void){
    int n, x = 0;
    printf("Inserisci una serie di numeri: ");
    do{
	scanf("%d",&n);
	x = x + n;
    }while( n != 0);

    printf("La somma dei numeri inseriti è: %d\n",x);
    return 0;
}
