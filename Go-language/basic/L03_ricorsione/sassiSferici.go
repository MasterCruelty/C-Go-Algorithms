package main

import "fmt"


func main(){
	var n int
	fmt.Print("Inserire l'altezza della piramide di sassi: ")
	fmt.Scan(&n)
	fmt.Println("Numero di sassi che compongono la piramide: ",sassi(n))
}


/*
  Dei sassi sono ammucchiati a formare una piramide con un sasso in cima.
  Questo sasso è posto al centro di un quadarto formato da 4 sassi(2 per lato)
  posti a loro volta sopra un quadrato formato da 9 sassi(3 per lato) e cosi via.
  f(1) = 1, f(2) = 5.
  scrivere una funzione ricorsiva che data l'altezza, restituisce il numero di sassi.

  hint: è informalmente la sommatoria di 'height' al quadrato.

*/
func sassi(height int) int {
	if(height > 0){
		return sassi(height-1) + height *height
	}else{
		return 0
	}
}
