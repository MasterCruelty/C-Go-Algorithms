#include <stdio.h>
#define N 100
//definisco la struttura di tipo data
typedef struct {
    int giorno;
    int mese;
    int anno;
}Data;

int main(void){
    int gg,mm,yy;
    int i = 0;
    Data date[N];
    printf("Inserisci la sequenza di date: \n");
    scanf("%d / %d / %d",&gg,&mm,&yy);
    date[i].giorno = gg;
    date[i].mese = mm;
    date[i].anno = yy;	
    i++;
    while(gg != 0 && mm != 0 && yy != 0){
        scanf("%d / %d / %d",&gg,&mm,&yy);
        if(gg != 0 && mm != 0 && yy != 0){
	    date[i].giorno = gg;
            date[i].mese = mm;
            date[i].anno = yy;	
    	    i++;
        }
    }
    for(int j = 0;j<i;j++){
        printf("%02d/%02d/%d\n",date[j].giorno,date[j].mese,date[j].anno);
    }
    return 0;
}
