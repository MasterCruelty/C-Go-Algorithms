#include <stdio.h>
#define N 100
int main(void){
    int n;
    int i = 0;
    int j = 0;
    int key = 0;
    int numeri[7] ={7,6,5,4,3,2,1};
    for(int i = 0;i<7;i++){
	printf("%d",numeri[i]);
    }
    printf("\n");
    //ogni volta che viene inserito un intero, il vettore viene scorso fino a trovare l'esatta collocazione del numero e vengono spostati in avanti i numeri già memorizzati. 
    for(int i = 0;i < sizeof(numeri);i++){
	key = numeri[i];
	j = i - 1;
	while( j >= 0 && numeri[j] > key){
	    numeri[j+1] = numeri[j];
	    j--;
	}
	numeri[j+1] = key;
    }
    for(int j = 0; j <i; j++){
        printf("%d\n",numeri[j]);
    }
    for(int i = 0;i<7;i++){
	printf("%d",numeri[i]);
    }
    printf("\n");
    return 0;
}
