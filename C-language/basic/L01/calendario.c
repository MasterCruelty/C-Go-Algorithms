#include <stdio.h>
#include <stdlib.h>

int main(void){
    int mese;
    int giorno;
    printf("Numero del mese (1= gennaio, ..., 12= dicembre): ");
    scanf("%d",&mese);
    printf("\nNumero del giorno di inizio mese (1= lunedi, ..., 7= domenica): ");
    scanf("%d",&giorno);
    printf("\nL  M  M  G  V  S  D\n");
    int cont = 1;
    int sett = 0;
    int end = 0;
    if(giorno != 1){
        for(int j = 1;j < giorno;j++){
	    printf("   ");
        }
    }	
    for(int i = 0; i <= abs(giorno-7); i++){
       printf("%d  ",i+1); 
       cont++;
    }
    printf("\n");
    if(mese == 2){
        end = 28;
    }
    else if(mese == 4 || mese == 6 || mese == 9 || mese == 11)
        end = 30;
    else
        end = 31;
    for(int i = cont; i <= end; i++){
        if(sett == 7){
            printf("\n");
	    sett = 0;
        }
	if(i >= 10){
	    printf("%d ",i);
	}
	else
	    printf("%d  ",i);
	sett++;
    } 
    printf("\n");
    return 0;
}
