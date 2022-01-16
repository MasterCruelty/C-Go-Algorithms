#include <stdio.h>

int main(void){
	int n = 1;
	int i = 0;
	int trovato = 0;
	printf("Ricerco i primi 10 quadrati perfetti...\n");
	while(i <= 10){
		for(int j = 0; j <= n;j++){

			if(	n == j * j){
				printf("%d\n",n);
				j++;
				trovato = 1;
			}
		}
		if(trovato == 1){
			i++;
			trovato = 0;
		}
		n++;
	}
	return 0;
}
