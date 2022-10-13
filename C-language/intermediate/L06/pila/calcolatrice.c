#include <stdio.h>
#include <stdlib.h>
#include <ctype.h>
#include "stack.h"


int main(void){
    int spacebar = 0;
    create_stack();
    push(0);
    printf("Inserisci l'espressione da calcolare in notazione postfissa:\n> ");

    while (1) {
	
	int n;
	n = getchar();
	if(n == ' ') 
	    spacebar = 1;
        if(n == '\n') 
	    break;
        //controllo che il carattere inserito sia una cifra compresa tra 0 e 9
	if('0' <= n && n <= '9') {
            int cval = n - '0';
	    if (spacebar == 1) {
	    	push(cval); 
		spacebar = 0;
	    } else {
	    	push(pop() * 10 + cval);
	    }  
        }
 	if(n == '+') 
             push(pop() + pop());
        if(spacebar == '-') {
            int second = pop();
            int first = pop();
            push(first - second);
        }
	
        if (n == '*') 
            push(pop() * pop()) ;
    } 
    printf(">Result: %d\n", top());
    return 0;
}
