package main


import "fmt"


/*
 * Data una matrice di n righe e n colonne, in alcune delle celle vi è una moneta con valore positivo.
 * Quando il giocatore raggiunge un quadrato contenente una moneta, se ne impossessa.
 * L'obiettivo è massimizzare il valore complessivo delle monete delle quali entra in possesso.
 * Per farlo deve muoversi lungo un cammino che inizia da un qualunque quadrato della colonna più a sinistra,
 * e terminare in un quadrato della colonna più a destra.
 * 
 * Sono stabilite inoltre queste regole:
 * da una cella il giocatore si può muovere nel quadrato immediatamente a destra sulla stessa riga, se quest'ultimo è vuoto.
 * da una cella il giocatore si può muovere nel quadrato immediatamente a destra sulla riga immediatamente superiore(se esiste), se quest'ultimo contiene moneta.
 * da una cella il giocatore si può muovere nel quadrato immediatamente a destra sulla riga immediatamente inferiore(se esiste), se quest'ultimo contiene moneta.
 * Data la matrice M in input, ricavare la matrice C contenente c_ij il valore massimo di monete.
 *********************************
 * Definisco formule di programmazione dinamica:
 * c[i,j] = m[i,j] + max{c[i-1,j-1],c[i+1,j-1]} se esiste moneta
 * c[i,j] = m[i,j] altrimenti
*/
func maxCoin(m [][]int) int {
	rows := len(m)
	cols := len(m[0])
	c := make([][]int,rows)
	for i:= range c{
		c[i] = make([]int,cols)
	}
	//riempio la prima colonna della matrice c
	for i:=0;i<=rows-1;i++{
		c[i][0] = m[i][0]
	}
	//Riempio la matrice c per colonne e non per righe
	max := 0
	for j := 1; j <= cols-1;j++{
		for i:= 0; i <= rows-1;i++{
			max = c[i][j-1]
			if (m[i][j] != 0) {
				if i > 0 && (c[i-1][j-1] > max) {
					max = c[i-1][j-1]
				}
				if i < rows-1 && (c[i+1][j-1] > max) {
					max = c[i+1][j-1]
				}
				c[i][j] = m[i][j] + max
			}else{
				c[i][j] = m[i][j]
			}
		}
	}
	//stampa della matrice di partenza e della matrice finale
	fmt.Println("Matrice M =====> Matrice C")
	for i:=0;i<=rows-1;i++{
		fmt.Println(m[i],"	",c[i])
	}
	//ricerca del valore massimo e lo restituisco
	max = c[0][cols-1]
	for i:=1; i <= rows-1;i++{
		if c[i][cols-1] > max{
			max = c[i][cols-1]
		}
	}
	return max


}

func main(){
	matrix := [][]int{
		{7,0,10,0,0},
		{0,4,0,12,0},
		{0,0,0,0,5},
		{0,8,0,50,0},
		{0,0,0,16,0},
	}

	fmt.Println("valore massimo di monete: ",maxCoin(matrix))
}
