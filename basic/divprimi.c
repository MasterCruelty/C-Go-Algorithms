#include <stdio.h>

int main(void) {

    int n = 0;
    int primo = 0;
    printf("Inserisci il numero: ");
    scanf("%d",&n);
    for(int i = 1; i <= n;i++){
        if(n % i == 0){
	    printf("%d\n",i);
	    primo++;
	}
    }
    if(primo == 2)
	printf("il numero %d risulta primo\n",n);	
    else
	printf("il numero %d non è primo\n",n);
    return 0;
}
