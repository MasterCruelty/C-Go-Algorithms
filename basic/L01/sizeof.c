#include <stdio.h>
#include <limits.h>
#include <float.h>

int main(void){
    int n;
    float f;
    char c;
    printf("int: %d\nfloat: %d\nchar: %d\n",sizeof(n),sizeof(f),sizeof(c));
    printf("minimo intero: %d\nmassimo intero: %d\n\nminimo float: %f\nmassimo float: %f\n\nminimo char: %c\nmassimo char: %c\n",INT_MIN,INT_MAX,FLT_MIN,FLT_MAX,SCHAR_MIN,SCHAR_MAX);
    return 0;	
}
