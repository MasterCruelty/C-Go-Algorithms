#include <stdio.h>
#include <stdlib.h>
#include "stack.h"


int main(){
    create_stack(); 
    printf("Lista vuota: %d\n",is_empty());
    push(7);
    push(42);
    printf("Top: %d\n",top());
    print_stack();
    pop();
    print_stack();
    pop();
    return 0;
}
