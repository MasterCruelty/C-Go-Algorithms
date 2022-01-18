#include <stdio.h>

int main(void){
    int contAperta, contChiusa = 0;    
    char c;
    while(( c = getchar()) != '.'){
        if(c == '(')
            contAperta++;
	if(c == ')')
	    contChiusa++;
    }
    printf("\n");
    if(contAperta > contChiusa)
        printf("L'espressione non è ben parentesizzata:\nmancano parentesi chiuse!\n");
    if(contChiusa > contAperta)
        printf("L'espressione non è ben parentesizzata:\nTroppe parentesi chiuse!\n");
    if(contAperta == contChiusa)
        printf("L'espressione è ben parentesizzata\n");
    return 0;
}
