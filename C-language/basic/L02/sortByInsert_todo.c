#include <stdio.h>
#define N 100
int main(void){
    int n = 1;
    int numeri[N];
    int i = 0;
    int salva = 0;
    for(int i = 0;i <N;i++){
        numeri[i] = 0;
    }
    printf("Inserisci la sequenza di numeri: \n");
    //ogni volta che viene inserito un intero, il vettore viene scorso fino a trovare l'esatta collocazione del numero e vengono spostati in avanti i numeri già memorizzati. TODO
    while(n != 0){
        scanf("%d",&n);
	numeri[i] = n;
        for(int i = 0;i<N;i++){
            if(numeri[i] < salva){
                numeri[i] = salva;
	        numeri[i-1] = n;
	    }
	}
        salva = n;
	i++;
    }
    for(int j = 0; j <i; j++){
        printf("%d\n",numeri[j]);
    }
    return 0;
}
