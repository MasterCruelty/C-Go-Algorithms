#include <stdio.h>

void hanoi(int n, char from, char to, char temp);
unsigned long contaMosse(int n);
//il numero minimo di mosse per risolvere il gioco delle torri di hanoi è di 2^n -1 dove n è il numero di dischi da spostare.
int main(){
    int num;
    unsigned long  mosse;
    char risp = '.';
    printf("Inserire il numero di dischi: ");
    scanf("%d",&num);
    printf("\n");
    printf("Vuoi vedere solo il numero di mosse necessarie? (s/n)\n");
    scanf(" %c",&risp);
    if(risp == 'n')
        hanoi(num,'0','2','1');
    mosse = contaMosse(num);
    printf("Numero mosse: %lu\n",mosse);
    return 0;
}

unsigned long contaMosse(int n){
    if(n == 0)
        return 0;
    return 2* contaMosse(n - 1) + 1;
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
