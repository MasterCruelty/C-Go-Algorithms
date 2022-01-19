#include <stdio.h>
#define N 10
int main(void){
    int cifre[N];
    int n,k;
    int check = 0;
    printf("Inserisci un numero: ");
    scanf("%d",&n);
    //inizializzo l'arrray con zeri
    for(int i = 0;i < N;i++){
        cifre[i] = 0;
    }
    //divido il mio numero per 10 n volte e popolo l'array contando le occorrenze delle cifre.
    while(n>0){
        k = n % 10;
	cifre[k]++;
	n /= 10;
    }
    //stampo quali cifre sono riptetute e quante volte
    for(int i = 0;i<N;i++){
        if(cifre[i] >1){
	    check = 1;
	    printf("%d --> è ripetuta %d volte\n",i,cifre[i]);
	}
    }
    if(check == 0)
        printf("Non ci sono cifre ripeture\n");
    return 0;
}
