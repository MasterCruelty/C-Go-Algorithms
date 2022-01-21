#include <stdio.h>
#include <ctype.h>


//dopo aver inserito una matrice r * c composta da lettere e asterischi, deve venire stampata la stessa matrice con tutte le lettere in basso alla matrice e 
//gli asterischi in alto
int main(void){
    int r,c, save_col, save_row;
    printf("Inserire numero righe: ");
    scanf("%d",&r);
    printf("\nInserire numero colonne: ");
    scanf("%d",&c);
    char matrix[r][c];
    char letter;
    printf("\nInserisci i valori della matrice: \n");
    //popolo la matrice inserendo i valori
    for(int i = 0; i < r;i++){
        for(int j = 0;j < c;j++){
	    scanf(" %c",&matrix[i][j]);
	}
    }
    //parto dal fondo della matrice
    //per ogni riga controllo su ogni colonna se è un carattere alfabetico
    //se non lo è, allora salvo l'indice della colonna e sostituisco il primo carattere alfabetico scorrendo solo quella colonna dall'alto
    for(int i = r - 1; i >= 0;i--){
        for(int j = c - 1;j >= 0;j--){
	    letter = '*';
	    if(isalpha(matrix[i][j])){
                continue;
	    }
	    else{
	        save_col = j;
		save_row = i - 1;
		for(int k = 0;k <= save_row ;k++){
		    if(isalpha(matrix[k][save_col])){
			letter = matrix[k][save_col];
			matrix[k][save_col] = '*';
			break;
		    }
		}
		matrix[i][j] = letter;
	    }
	}
    }
    printf("\n\n");
    for(int i = 0; i < r;i++){
        for(int j = 0;j < c;j++){
	    printf("%c ",matrix[i][j]);
	}
	printf("\n");
    }
    return 0;
}
