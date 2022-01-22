#include <stdio.h>

void hanoi(int n, char from, char to, char temp);
//il numero minimo di mosse per risolvere il gioco delle torri di hanoi è di 2^n -1 dove n è il numero di dischi da spostare.
int main(){
    int num;
    printf("Inserire l'altezza della pila: ");
    scanf("%d",&num);
    hanoi(num,'0','2','1');
    return 0;
}

void hanoi(int n,char from, char to, char temp){
    if(n == 1){
        printf("Sposto 1 da paletto %c a paletto %c\n",from,to);
	return;
    }
    hanoi(n-1,from,temp,to);
    printf("Sposto %d da paletto %c a paletto %c\n",n,from,to);
    hanoi(n-1,temp,to,from);
}
