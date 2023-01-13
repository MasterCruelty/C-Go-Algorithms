package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/* 
 * Modellazione: nodi sono svincoli, archi le gallerie
 * Se esiste una galleria, vi è un arco tra due svincoli.
 * Il peso su ogni arco è determinato dalla luminosità della galleria.
 * È un problema di esistenza di cammino, scegliendo sempre e solo l'arco di peso minore(strategia di harmony)
 * Algoritmo: c'è una lista di nodi visitati, un dato che indica il numero di archi attraversati e un dato che indica in che nodo siamo se sono in S finito.
 * Se non sono in S: scelgo l'arco di peso inore e mi muovo nel nodo collegato dall'arco. Si aggiunge 1 alla variabile degli archi visitati
 * Se il nodo è nella lista dei visitati: non esiste cammino che rispetti la strategia e restituisco -1
 * Se il nodo non è nella lista si inserisce il nodo nella lista e si ripete l'algoritmo. 
*/

type arcoMigliore []int

//pos 0 = vicino migliore
//pos 1 = luminosità

type Grafo = []arcoMigliore //indice = nodo di partenza

func main(){
	var numGallerie int
	var nodiVisitati []int
	var posizione int

	g,h,s := CreaGrafo()

	posizione = h

	for {
		if posizione == s{
			fmt.Println(numGallerie)
			return
		}
		posizione = g[posizione-1][0]
		numGallerie++;
		if !(CercaInSlice(nodiVisitati,posizione)){
			nodiVisitati = append(nodiVisitati,posizione)
		}else{
			fmt.Println(-1)
			return
		}
	}
}

func CreaGrafo() (g Grafo, h int, s int) { //restituisce il grafo con solo le gallerie sceglibili da harmony, il nodo di partenza dell'algoritmo e il nodo d'arrivo
	var num []int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		campi := strings.Split(scanner.Text(),"")
		for _,cont := range campi{
			n,_ := strconv.Atoi(cont)
			num = append(num,n)
		}
		if len(num) == 4{ // solo prima riga input
			g = make(Grafo, num[0])
			h = num[2]
			s = num[3]
			num = nil
			continue
		}

		if(g[num[0]-1] == nil) || (g[num[0]-1][1] > num[2]) {
			g[num[0]-1] = arcoMigliore{num[1],num[2]}
		}
		if (g[num[1]-1] == nil) || (g[num[1]-1][1] > num[2]){
			g[num[1]-1] = arcoMigliore{num[0],num[2]}
		}
		num = nil
	}
	//fmt.Println(g)
	return g,h,s
}

func CercaInSlice(s []int,n int) bool{
	for _,num:= range s{
		if num == n{
			return true
		}
	}
	return false
}
