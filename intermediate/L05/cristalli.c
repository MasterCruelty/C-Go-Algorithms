#include <stdlib.h>
#include <stdio.h>


//prototipi funzioni per gestire il cristallo
int latoCristallo(int t);
char** creaMatrice(int n);
void stampaMatrice(char** m, int n);
void cristallo(char** m, int l);
void crist(char** m, int r0, int c0, int l);

int main() {
    char** matrix ;
    int t , lato;
    scanf ( "%d" , &t );                // legge il tempio t
    
    if (t >= 0) {
        lato = latoCristallo(t);        // il lato definisce la dimensione della matrice
        matrix = creaMatrice(lato);     // crea matrice per rappresentare il cristallo
        cristallo(matrix, lato);        // costruisce il cristallo sulla matrice con il lato assegnato
        stampaMatrice(matrix, lato);    // stampa la matrice risultante
    }
    //il seguente ciclo non fa altro che liberare lo spazio di memoria occupato dalla matrice costruita e stampata nelle linee precedenti.
    for (int i = 0; i < lato; i++) {
        free(matrix[i]);
    }
    free(matrix);
}

//definizione ricorsiva per assegnare il lato del cristallo.
//se t = 0 l vale 1; se t = 1 l vale 3; ==> se t è >=1 allora l vale 2 * latoCristallo(t-1) +1
int latoCristallo(int t) {
    if (t == 0)
        return 1;
    else 
        return 2 * latoCristallo(t-1) + 1;
}

//creazione della matrice, allocazione della memoria necessaria e viene popolata da punti.
char** creaMatrice(int n) {
    char** mat = (char**) malloc(n * sizeof(char*));
    for (int i = 0; i < n; i++) {
        mat[i] = (char*) malloc(n * sizeof(char));
        for (int j = 0; j < n; j++)
            mat[i][j] = '.';
    }

    return mat;
}

//semplicemente la stampa della matrice riga per colonna
void stampaMatrice(char** m, int n) {
    for (int i = 0; i < n; i++) {
        for (int j = 0; j < n; j++)
            printf("%c", m[i][j]);
        printf("\n");
    }
}

//inizializzazione della creazione del cristallo in riga zero e colonna zero
void cristallo(char** m, int l) {
    crist(m, 0, 0, l);
}


//creazione del cristallo ricorsivamente
void crist(char** m, int r0, int c0, int l) {
    if (l == 1) {
        m[r0][c0] = '*';
    } else {
        // Si assume che la matrice contenga di base solo '.'.
        m[r0+l/2][c0+l/2] = '*';
        crist(m, r0, c0, l/2);
        crist(m, r0, c0 + l/2 + 1, l/2);
        crist(m, r0 + l/2 + 1, c0, l/2);
        crist(m, r0 + l/2 + 1, c0 + l/2 + 1, l/2);
    }
}
