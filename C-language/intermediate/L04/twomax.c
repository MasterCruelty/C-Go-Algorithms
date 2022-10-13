#include <stdio.h>

void max2(int a[],int n, int *max1,int *max2);

int main(void){
    int a[] = {1,6,4,5,7,11};
    int n = 6;
    int m,m2;
    max2(a,n,&m,&m2);
    printf("max1: %d\nmax2: %d\n",m,m2);
    return 0;
}

void max2(int a[],int n,int *max1,int *max2){
    int *p;
    //inizializzo entrambi i puntatori con il primo valore del vettore.
    *max1 = *a;
    *max2 = *a;
    for(p = a; p< a +n; p++){
	if(*p > *max2){
	    if(*p > *max1){
	        *max2 = *max1;
		*max1 = *p;
	    }
	    else
	        *max2 = *p;
	}	
    }
}
