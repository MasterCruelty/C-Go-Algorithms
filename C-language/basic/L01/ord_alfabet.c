#include <stdio.h>
#include <math.h>
int main(void){
	char x1,x2;
	printf("Inserisci i due caratteri: ");
	scanf("%c %c",&x1,&x2); //mettere &x1 significa assegnare il valore inserito da tastiera all'indirizzo in cui è memorizzata x1
	int result = abs(x1-x2) + 1; //calcolo il valore assoluto della differenza tra i due caratteri(ogni char rappresenta un intero in ASCII)
	printf("\nLa distanza tra i due caratteri è: %d\n",result);
	return 0;
}
