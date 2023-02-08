package main

import (
	"fmt"
)


type oggetto struct{
	name string
	val int
	op rune
	sx string
	dx string
	tipo string
}

/*
 * Foresta di alberi implementata come tabella dei padri.
 * Il tipo oggetto contiene al suo interno i figli sx e dx, quindi si può risalire con la mappa dei padri e scendere con l'oggetto salvato
 * se il padre è vuoto, allora è la radice dell'albero
*/
type foresta struct {
	oggetti map[string]*oggetto
	padri map[string]string
}

/*
 * Inizializzo una mappa per la tabella dei padri
 * Successivamente scorro la mappa degli oggetti
 * Se l'oggetto è un'operazione, allora ha due figli
 * E quindi popolo la tabella dei padri inserendo l'oggetto corrente come padre dei suoi figli.
*/
func costruisciForesta(oggetti map[string]*oggetto) foresta{
	padri := make(map[string]string)

	for k:= range oggetti{
		padri[k]  = ""
	}

	for key,value := range oggetti{
		if value.tipo == "op"{
			padri[value.sx] = key
			padri[value.dx] = key
		}
	}
	return foresta{oggetti,padri}
}

/*
 * Ispeziono la foresta usando la stringa n come chiave nella mappa degli oggetti
 * Se il valore non è la stringa nulla, allora restituisco sx e true
 * Altrimenti restituisco stringa vuota e false
*/
func sx(f foresta,n string) (string,bool) {
	sx := f.oggetti[n].sx
	if sx == ""{
		return "",false
	}
	return sx,true
}

/*
 * Ispeziono la foresta usando la stringa n come chiave per la tabella dei padri
 * Se trovo un valore diverso dalla stringa vuota, restituisco padre e true
 * Altrimenti restituisco stringa vuota e false
*/
func up(f foresta,n string) (string,bool){
	padre := f.padri[n]
	if padre == ""{
		return "",false
	}
	return padre,true
}

/*
 * Uguale alla funzione sx ma si ispeziona sull'attributo dx
*/
func dx(f foresta,n string) (string,bool) {
	dx := f.oggetti[n].dx
	if dx == ""{
		return "",false
	}
	return dx,true
}

/*
 * Viene eseguita una visita in profondità in ordine simmetrico(inorder)
 * Quindi prima figlio sx, poi radice, poi figlio destro sfruttando la ricorsione
*/
func stampaAlbero(f foresta,n string){
	Sx,foundsx := sx(f,n)
	if foundsx{
		stampaAlbero(f,Sx)
	}

	if f.oggetti[n].tipo == "val" {
		fmt.Print(n, " (val = ",f.oggetti[n].val,")\n")
	}else{
		fmt.Println(n)
	}

	Dx,founddx := dx(f,n)
	if founddx {
		stampaAlbero(f,Dx)
	}
}


/*
 * Se sono su una foglia, stampo direttamente il valore dell'oggetto.
 * Altrimenti ricorsivamente eseguo la funzione sui figli sx e dx.
 * In base al simbolo di operazione la eseguo e restituisco il valore finale.
*/
func calcolaPrezzo(f foresta, n string)  int{
	if f.oggetti[n].tipo == "val"{
		return f.oggetti[n].val
	}

	sx := calcolaPrezzo(f,f.oggetti[n].sx)
	dx := calcolaPrezzo(f,f.oggetti[n].dx)
	var result int

	switch f.oggetti[n].op {
	case '+':
		result = sx+dx
		break
	case '-':
		result = sx-dx
		break
	case '*':
		result = sx*dx
		break
	case '/':
		result =sx/dx
		break
	}
	return result
}

func main(){
	fmt.Println("todo")
}
