#include <stdio.h>
#include <math.h>

int main(void){
	int a,b,c;

	printf("Inserisci i tre coefficienti: ");
	scanf("%d %d %d",&a,&b,&c);

	float delta = sqrt((b*b) - 4 * a * c);
	float result1 = (-b + delta) / (2 * a);
	float result2 = (-b - delta) / (2 * a);
	printf("\nDelta trovato: %f\n",delta);
	printf("\nLe soluzioni sono: \n %f \n %f\n",result1,result2);
	return 0;
}
