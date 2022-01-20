#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#define N 10

//passeggiata aleatoria(random walk)
typedef enum{False,True} bool;

int main(void){
    char matrix[N][N];
    int i,j;
    //popolo la matrice con soli punti
    for(i = 0;i <N;i++){
        for(j = 0;j<N;j++){
	    matrix[i][j] = '.';
        }
    }
    char c = 'A';
    i = 0;
    j = 0;
    matrix[i][j] = c;
    //inizializzo un seme casuale
    srand(time(NULL));
    //finchè ci sono lettere disponibili e spazio per posizionarle
    while(c < 'Z'){
        bool check_UP = False;
	bool check_RIGHT = False;
	bool check_DOWN = False;
	bool check_LEFT = False;
	bool try_UP = False;
	bool try_RIGHT = False;
	bool try_DOWN = False;
	bool try_LEFT = False;
	bool placed = False;
        // se il test è negativo si continua altrimenti si stoppa perchè ha fallito
	while(!(try_UP && try_RIGHT && try_DOWN && try_LEFT)){
            switch(rand() % 4){
		case 0:
		    if(i != 0 && matrix[i-1][j] == '.'){
		        c++;
			matrix[i-1][j] = c;
			i--;
			check_UP = True;
		    }
		    try_UP = True;
		    break;
		case 1:
		    if(j != (N - 1) && matrix[i][j+1] == '.'){
		        c++;
			matrix[i][j+1] = c;
			j++;
			check_RIGHT = True;
		    }
		    try_RIGHT = True;
		    break;
		case 2:
		    if(i != (N - 1) && matrix[i+1][j] == '.'){
		        c++;
			matrix[i+1][j] = c;
			i++;
			check_DOWN = True;
		    }
		    try_DOWN = True;
		    break;
		case 3:
		    if(j != 0 && matrix[i][j-1] == '.'){
		        c++;
			matrix[i][j-1] = c;
			j--;
			check_LEFT = True;
		    }
		    try_LEFT = True;
		    break;
		default: 
		    break;
	    }
	    //se viene superato almeno un check, una lettera viene posizionata e non c'è bisogno di riprovare.
	    if(check_UP || check_RIGHT || check_DOWN || check_LEFT){
	        placed = True;
		break;
	    }
       }
       //se sono state provate tutte le direzioni e nessun check è passato allora non c'è più spazio dove mettere le lettere, quindi mi fermo
       if(try_UP && try_RIGHT && try_DOWN && try_LEFT && !placed){
           break;
       }
    }
    for(i = 0;i<N;i++){
        for(j=0;j<N;j++){
	    printf("%c ",matrix[i][j]);
        }
	printf("\n");
    }
    return 0;
}
