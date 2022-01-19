#include <stdio.h>
#include <stdlib.h>
//dato un n da input, inserire in ordine crescente n numeri tranne uno, l'algoritmo dovrà trovare il valore mancante
int main(void){
    int n;
    scanf("%d",&n);
    int a[n];
    int mancante = 0;
    int k = 0;
    int z = 0;
    //popolo l'array da standard input
    for(int i = 0;i < n; i++){
        scanf("%d",&a[i]);
    }
    int sx = 0;
    int dx = n -1;
    //caso in cui il numero n è dispari
    if(n % 2 != 0){
	while(abs(a[sx] - a[dx]) % 2 != 0){
            sx++;
            dx--;
	}
        mancante = a[dx] + 1;
    }
    //TODO questo caso
    if(n % 2 == 0){
	while(k <= 1 || z <= 1){
	    if(abs(a[sx] - a[dx]) % 2 == 0){
		k++;
		dx--;
		continue;
            }
	    if(abs(a[sx] - a[dx]) % 2 != 0){
		z++;
		dx--;
		continue;
            }
	    k = 0;
	    z = 0;
            //sx++;
            dx--;
	}
	mancante = a[dx] + 3;
    }
    printf("%d\n",mancante);
    return 0;
}
