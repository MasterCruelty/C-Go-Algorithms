#include <stdio.h>
#include <stdlib.h>

#define N 15

char **creaMatrice(int n);

//Allocazione dinamica in memoria di una matrice n * n
int main() {
  **matrix = creaMatrice(N);
  for(int r = 0; r <= N; r++){
    printf("\n");
    for(int c = 0; c <= N; c++){
      printf("%c",matrix[r][c]);
    }
  }
  return 0;
}

char **creaMatrice ( int n ){
  char ** m;
  int r , c;
  m = malloc ( n * sizeof ( char * ) );
  for ( r = 0; r < n; r ++ ) {
    *( m+r) = malloc ( n * sizeof ( char ) );
  }
  for ( r = 0; r < n; r ++ )
    for ( c = 0; c < n; c ++ )
      m[r ][ c] = ’.’;
  return m;
}
