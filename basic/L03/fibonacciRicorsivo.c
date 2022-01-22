#include <stdio.h>
unsigned long f_rec(int);

int main(void){
    int num;
    unsigned long result;
    printf("Inserire un numero: ");
    scanf("%d",&num);
    printf("La sequenza di Fibonacci calcolata in %d è:\n",num);
    result = f_rec(num);
    printf("%lu\n",result);
    return 0;
}


unsigned long f_rec(int n){
    if(n == 1 || n == 2)
        return 1;
    else
	return f_rec(n -1) + f_rec(n-2);
}

