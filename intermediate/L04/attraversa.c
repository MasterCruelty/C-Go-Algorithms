#include <stdio.h>

#define N 100
//inserimento di valori in un array tramite puntatore e uso del puntatore nei cicli per riempire l'array e per scorrerlo al contrario
int main(void){
    int a[N];
    int *p;
    for(p = a;p < a + N;p++){
        scanf("%d",p);
	if(*p == 0)
	    break;
    }
    while(--p >= a)
        printf("%d ", *p);
    printf("\n");
    return 0;	     
}
