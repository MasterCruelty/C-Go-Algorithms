#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <ctype.h>


//struct di tipo Registro con posti e nomi
typedef struct{
   int posti;
   char **nomi;
}Registro;

//inizializzo il puntatore a registro NULL
Registro *registro = NULL;

//prototipi funzioni principali per le features
Registro* newBook(int);
void cancelBook();
void book(int, char*);
void cancel(int);
void printBook();
void move(int,int);

//prototipi funzioni ausiliarie
char *read_word(void);
void *my_malloc(size_t);
void *my_realloc(void*, size_t);

int main(){
   char input, *nome;
   int n, to;
   while ((input = getchar()) != 'f') {
      switch(input) {
         case 'b': // b con intero n => newBook(n)
            scanf("%d", &n);
            if (n <= 0 )
               break;
            if (registro != NULL)
               cancelBook();
            registro = newBook(n);
            break;
         case '+': // simbolo + con intero k e stringa name => book(k, name)
            if (registro != NULL) {
               scanf("%d ", &n);
               nome = read_word();
               if (n < 0 || n >= (registro->posti) || strlen(nome)==0){
		  printf("ERRORE! POSSIBILI CAUSE:\n1)STRINGA VUOTA: NESSUN INSERIMENTO\n2)INSERITO N TROPPO GRANDE O NEGATIVO\n");
                  break;
	       }
               book(n, nome);
            } else {
               printf("ERRORE: REGISTRO NON ESISTENTE\n");
            }
            break;
         case '-': // simbolo - con intero k => cancel(k)
            if (registro != NULL) {
               scanf("%d", &n);
               if (n < 0 || n >= (registro->posti))
                  break;
               cancel(n);
            } else {
               printf("ERRORE: REGISTRO NON ESISTENTE\n");
            }  
            break;
         case 'm': // m con intero from e intero to => sposta nome da from a to
            scanf("%d ", &n);
            scanf("%d", &to);
            if (n < 0 || n >= (registro->posti) || to < 0 || to >= (registro->posti))
               break;
            if (registro != NULL) {
               move(n, to);
            } else {
               printf("ERRORE: REGISTRO NON ESISTENTE\n");
            }  
            break;
         case 'p': // p =>  printBook()
            if (registro != NULL) {
               printBook();
            } else {
               printf("ERRORE: REGISTRO NON ESISTENTE\n");
            }  
            break;
      }
   } 
}
// Crea un nuovo registro che permetta la prenotazione di n posti, da 0 a n 1
// Se esiste già un registro di prenotazione, quest'ultimo deve essere cancellato
Registro *newBook(int nPosti) {
   Registro *registro = calloc(1, sizeof(Registro));
   (registro -> posti) = nPosti;
   (registro -> nomi) = calloc(nPosti, sizeof(char*));
   return registro;
}
// deAlloca lo spazio di memoria occupato dal registro
void cancelBook() {
   for (char **nome = (registro->nomi); nome < (registro->nomi)+(registro->posti); nome++)
      free(*nome);
   free(registro);
}
// Se il posto k è libero, prenota il posto k per il cliente identificato da name. Altrimenti, stampa un messaggio di errore.
void book(int k, char *name){
   if ((registro -> nomi)[k] != NULL){
      printf("POSTO GIA' OCCUPATO\n");
      return;
   }
   (registro -> nomi)[k] = name;
}
// Se il posto k è occupato, cancella la prenotazione di k. Altrimenti, stampa un messaggio di errore.
void cancel(int k) {
   if ((registro -> nomi)[k] == NULL){
      printf("POSTO NON OCCUPATO\n");
      return;
   }
   free((registro -> nomi)[k]);
   (registro -> nomi)[k] = NULL;
   printf("Eseguito.\n");
}
// Sposta il cliente dal posto from al posto to se ciò è possibile (ossia, from è occupato e to libero). 
// Altrimentistampa un messaggio di errore.
void move(int from, int to) {
   if ((registro->nomi)[to] == NULL && (registro->nomi)[from] != NULL){
      (registro->nomi)[to] = (registro->nomi)[from];
      print("Eseguito.\n");
      (registro->nomi)[from] = NULL;
   } else {
      printf("IMPOSSIBILE SPOSTARE IL POSTO\n");
   }
}
// Stampa il contenuto del registro (posti prenotati)
void printBook() {
   printf("REGISTER[0..%d]:\n", registro -> posti-1);
   for (char **parola = registro->nomi; parola < registro->nomi+registro->posti; parola++){
      if (*parola != NULL)
         printf("%d --> %s\n", parola-(registro->nomi), *parola);
   }
}

//funzione ausiliaria per leggere i nomi da registrare
char *read_word(void) {
   char *nome,c;
   int n = 0;
   int size = 5;
   nome = my_malloc(size * sizeof(char));
   while(1){
      scanf(" %c",&c);
      if(n >= size){
         size *=2;
	 nome = my_realloc(nome,size *sizeof(char));
      }
      if(c == '.'){
	 printf("%s registrato.\n",nome);
	 nome[n] = '\0';
	 break;
      }
      nome[n++] = c;
   }
   return nome;
}

//ri-definizione delle funzioni malloc e realloc con opportuni controlli se il puntatore è nullo
void *my_malloc(size_t n) {
  void *p;
  p = malloc(n);
  if (p == NULL) {
    printf( "errore allocazione malloc\n" );
    exit(EXIT_FAILURE);
  } 
  return p;
}
void *my_realloc(void *p, size_t n) {
   p = realloc(p, n);
   if (p == NULL) {
      printf("errore realloc\n");
      exit(EXIT_FAILURE);
   }
   return p;
}
