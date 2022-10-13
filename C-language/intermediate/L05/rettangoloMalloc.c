#include <stdio.h>
#include <math.h>
#include <stdlib.h>

//definisco la struct di tipo rettangolo
typedef struct{
    float x,y;
}Punto;

typedef struct{
    Punto p1;
    Punto p2;
} Rettangolo;

//prototipo di funzione che restituisce un puntatore a variabile di tipo Rettangolo
Rettangolo *creaRettangolo();

int main(void){
    float b,h,area,duep;
    Rettangolo *r;
    printf("Rettangolo:\n");
    r = creaRettangolo();
    //essendo r un puntatore a struttura non uso il punto ma l'operatore -> per indirizzarmi ai suoi attributi
    b = r -> p2.x - r -> p1.x;
    h = r -> p2.y - r -> p1.y;
    area = b*h;
    duep = 2 *(b+h);
    printf("area: %f\nperimetro: %f\n",area,duep);
    return 0;
}


Rettangolo *creaRettangolo(){
    Rettangolo *r;
    //alloco dinamicamente lo spazio necessario per una variabile di tipo rettangolo
    r = malloc(sizeof(Rettangolo));
    printf("Inserisci coordinate del punto in basso a sx\n");
    //nella scanf passo come argomento l'indirizzo di memoria dove si trovano gli attributi del rettangolo
    scanf("%f %f",&r -> p1.x,&r -> p1.y);
    printf("Inserisci coordinate del punto in alto a dx\n");
    scanf("%f %f",&r -> p2.x,&r -> p2.y);
    return r;
}
