#include <stdio.h>

int main(void){
    float a,b,c;
    printf("Inserisci tre numeri separati da spazi: ");
    scanf("%f %f %f",&a,&b,&c);

    if(a == b && a == c)
    	printf("Il triangolo è equilatero\n");
	else if(a == b || a == c || b == c)
		printf("Il triangolo è isoscele\n");
	else
		printf("Il triangolo è scaleno\n");
	return 0;
}
