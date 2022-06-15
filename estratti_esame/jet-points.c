#include <stdio.h>
#include <stdlib.h>



/*
 * La colonna di destra sono punti jet mentre la colonna di sinistra sono valori in dollari nella matrice qui di seguito.
 * Passata questa matrice a una funzione insieme a un numero K di jet-points, si vuole trovare il massimo valore di dollari 
 * acquistabili con jet points. Ad esempio nella tabella 20000 punti equivalgono a 100 dollari, quindi se ne possiedo 40000,
 * sceglierei due volte la prima riga avendo così accesso a 200 dollari(che sono di più rispetto ai 195 se avessi scelto l'ultima riga).
 *
 *
 * Il problema verrà risolto tramite programmazione dinamica riempiendo un vettore V[i] con i sotto valori massimi per ogni riga della tabella.
 * Infine verrà stampato il valore massimo contenuto nel vettore, ovvero il nostro risultato ottimo.
 *
 * */

int maxValue(int matrice[][2],int k);

int main(void){
	int matrix[11][2] = {{100,20000},
						 {107,22000},
						 {124,24000},
						 {133,26000},
						 {139,28000},
						 {155,30000},
						 {172,32000},
						 {178,34000},
						 {184,36000},
						 {190,38000},
						 {195,40000}};
	int n = maxValue(matrix,140000);
	printf("valore massimo: %d \n",n);
	return 0;
}

//funzione per trovare il valore masimo
int maxValue(int matrice[][2], int k){
	int V[11] = {0,0,0,0,0,0,0,0,0,0,0};
	int save = k;
	int max = 0;
	//riempimento array V[i] 
	for(int i = 10; i >= 0;i--){
		k = save;
		for(int j = i;j >= 0;j--){
			if(k >= matrice[j][1]){
				V[i] += matrice[j][0];
				k = k - matrice[j][1];
			}
		}
	}
	//trovo il max all'interno di V[i] e lo restituisco
	for(int i = 0;i < 11;i++){
		if(V[i] > max)
			max = V[i];
	}
	return max;
}
