#include <stdio.h>

int potenza(int num,int pow);

int main(void){
    int n,power;
    printf("Inserisci un numero e la potenza da calcolare: ");
    scanf("%d %d",&n,&power);
    printf("il risultato della potenza è: %d\n",potenza(n,power));
    return 0;
}

int potenza(int num,int pow){
    if(pow == 0)
        return 1;
    else
	return num * potenza(num,pow -1);
}
