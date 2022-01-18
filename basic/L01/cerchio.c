#include <stdio.h>
#include <math.h>
int main(void){
	float  r, area;
	float pi = M_PI;
	
	printf("Inserisci il raggio del cerchio: ");
	scanf("%f",&r);
	area = pi * r * r;
	printf("\nl'area del cerchio è: %0.4f\n",area);
	return 0;
}
