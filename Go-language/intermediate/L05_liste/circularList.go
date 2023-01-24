package main

import "fmt"

//Definizione struttura dati di tipo lista circolare
type circNode struct {
	value int
	next *circNode
}

//Crea un nuovo nodo per la lista
func newNode(val int) *circNode{
	return &circNode{value:val,next:nil}
}

//Setto la testa della lista circolare facendo pntare il successivo a sè stesso
func setHead(val int) *circNode{
	head := newNode(val)
	head.next = head
	return head
}

//Restituisce l'ultimo nodo della lista, funzione ausiliaria per AddNewNode
func getLastNode(head *circNode) *circNode{
	lastNode := head
	for lastNode.next != head {
		lastNode = lastNode.next
	}
	return lastNode
}

//Aggiunge un nuovo nodo alla lista circolare.
//Creo un nuovo nodo, recupero l'ultimo della lista a partire dal nodo di testa passato.
//Faccio puntare il successore dell'ultimo al nuovo nodo e il successore del nuovo nodo alla testa.
//In questo modo ho ancora una lista circolare
func AddNewNode(val int, head *circNode){
	newNode := newNode(val)
	lastNode := getLastNode(head)
	lastNode.next = newNode
	newNode.next = head
}

//main di test dove creo e faccio scorrere completamente una lista circolare partendo dalla testa e ritornandoci una volta percorsa tutta.
func main(){
	head := setHead(5)
	AddNewNode(10,head)
	AddNewNode(2,head)
	AddNewNode(-5,head)

	//scorro la lista circolare dalla testa fino a ritornare alla testa
	current := head
	for {
		fmt.Println(current.value)
		current = current.next
		if current == head{
			break
		}
	}
	fmt.Println(current.value)
}
