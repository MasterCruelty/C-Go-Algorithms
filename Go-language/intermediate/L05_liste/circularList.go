package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

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

//Stampa della lista circolare a partire dal nodo p contenente il valore 0
func stampaDaZero(p *circNode){
	current := p
	for {
		if current.value == 0{
			break
		}
		current = current.next
	}
	for {
		fmt.Println(current.value)
		current = current.next
		if current.value == 0{
			break
		}
	}
}

func Sposta(p *circNode){
	var lunghezza int
	temp := p
	for temp.next != p{
		temp = temp.next
		lunghezza++
	}
	current := p
	if p.value > 0 {
		for i:= 0;i<=p.value;i++{
			current = current.next
		}
		save := current.value
		current.value = p.value
		p.value = save
	}else{
		for i:= 0;i<=lunghezza - p.value;i++{
			current = current.next
		}
		save := current.value
		current.value = p.value
		p.value = save
	}
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


	//da qui inizio a testare la funzione stampaDaZero
	head2 := setHead(3)
	//leggo gli interi da standard input
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
			if scanner.Text() == "exit"{
				break
			}
			val,_ := strconv.Atoi(scanner.Text())
			//aggiungo il valore letto alla lista circolare
			AddNewNode(val,head2)
	}
	//invoco la funzione stampaDaZero
	stampaDaZero(head2)
	fmt.Println("Sposto il nodo:",head2.value)
	//ora testo la funzione sposta
	Sposta(head2)
	stampaDaZero(head2)
}
