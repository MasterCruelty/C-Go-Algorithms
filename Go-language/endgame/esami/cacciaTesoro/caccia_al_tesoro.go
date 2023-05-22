package main

/*
 * Il problema è modellabile tramite un albero binario come struttura dati.
 * La radice è il primo oggetto da cui iniziare a cercare.
 * Se il nodo corrisponde a un'operazione, esso avrà due figli, altrimenti è una foglia.
 * riformulazione compiti nei termini della struttura dati:
 * 
 * a) dato il nome associato a un nodo, se è un numero restituiamo quel numero, altrimenti se è un simbolo di operazione 
 *    del tipo a + b con a figlio sx e b figlio dx, visito a e b e se sono numeri li associo all'operazione e restituisco il risultato,
 *    altrimenti continuo la visita dell'albero per poi ritornare il risultato finale.
 * 
 * b) todo
 *
 * c) Inizio a fare una visita in profondità dell'albero, per esempio in ordine simmetrico, calcolando per ogni nodo la funzione associata al compito a.
 *    Tutti questi numeri li metto dentro un array e successivamente lo ordino in ordine crescente.
 *
 * d) Dato un nodo a, se a è il nodo radice, non faccio niente.
 *    Altrimenti, a avrà un padre e lo aggiungo alla lista dei nodi a cui serve a, iterando fino alla radice.
 *
 * È opportuno implementare la struttura dati scelta con una struct che definisce il nodo con una serie di campi che indica nome, valore, operazione e i figli.
 * Inoltre ho un'altra struct contenente una mappa che funge da tabella dei padri per segnare il padre di ogni nodo, poi un'ulteriore mappa che contiene tutti i
 * nodi, con chiave il nome del nodo e il valore la struct del nodo corrispondente.
 * 
 *
*/
import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"fmt"
)

type treeNode struct{
	nome string
	value int
	op string
	sx string
	dx string
}

type  tree struct {
	padri map[string]string
	nodi map[string]*treeNode
}

/*
 * funzione che svolge il compito assegnato al punto a
 * Valutazione complessità:
 * La complessità dipende dall'altezza dell'albero, 
 * quindi questa funzione ha costo O(logn)
*/
func calcolaNumero(t tree,ogg *treeNode) int {
	result := 0
	if ogg.value != 0{
		return ogg.value
	}else{
		sx := calcolaNumero(t,t.nodi[ogg.sx])
		dx := calcolaNumero(t,t.nodi[ogg.dx])
		switch ogg.op {
			case "+":
				result = sx + dx
				break
			case "-":
				result = sx - dx
				break
			case "*":
				result = sx * dx
				break
			case "/":
				result = sx / dx
				break
		}
	}
	return result
}



func main(){
	tr := &tree{make(map[string]string),make(map[string]*treeNode)}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		line := scanner.Text()
		s := strings.Split(line,": ")
		name := s[0]
		val,err := strconv.Atoi(s[1])
		if err != nil{
			op := strings.Fields(s[1])
			tr.nodi[name] = &treeNode{name,0,op[1],op[0],op[2]}
			tr.padri[op[0]] = name
			tr.padri[op[2]] = name
		}else{
			tr.nodi[name] = &treeNode{name,val,"","",""}
		}
	}
	fmt.Println()
	fmt.Println("test calcolaNumero:")
	a := calcolaNumero(*tr,tr.nodi["microonde"])
	fmt.Println("valore microonde: ",a)
	a = calcolaNumero(*tr,tr.nodi["mensola"])
	fmt.Println("valore mensole: ",a)
	a = calcolaNumero(*tr,tr.nodi["frigorifero"])
	fmt.Println("valore frigorifero: ",a)
	a = calcolaNumero(*tr,tr.nodi["sedia"])
	fmt.Println("valore sedia: ",a)
	a = calcolaNumero(*tr,tr.nodi["bambola"])
	fmt.Println("valore bambola: ",a)
	a = calcolaNumero(*tr,tr.nodi["letto"])
	fmt.Println("valore letto: ",a)
}
