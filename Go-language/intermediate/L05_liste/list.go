package main

import (
	"fmt"
	"strings"
	"strconv"
	"bufio"
	"os"
)

/*
  struttura per definire la lista concatenata
  un oggetto item intero e un puntatore al prossimo nodo della lista
  tante strutture di tipo listNode comporranno una lista concatenata
*/
type listNode struct {
	item int
	next *listNode
}

//struttura che punta al nodo di testa della lista concatenata
type linkedList struct {
	head *listNode
}

/*
   +: aggiunge il valore alla lista(se non c'è)
   -: elimina il valore dalla lista(se c'è)
   ?: controlla se il valore è nella lista
   p: stampa la lista
   o: stampa la lista al contrario
   d: elimina tutti gli elementi dalla lista
   c: restituisce il numero di elementi totale nella lista
*/
func main() {
	var list linkedList
	var expr string
	scanner := bufio.NewScanner(os.Stdin)
	for  {
		scanner.Scan()
		expr = scanner.Text()
		split := strings.Split(expr," ")
		if split[0] == "f" {
			break
		}
		switch(split[0]) {
			case "+":
				n,_ := strconv.Atoi(split[1])
				_,p := searchList(list,n)
				if p == nil {
					list = addNewNode(list,n)
				}
				break
			case "-":
				n,_ := strconv.Atoi(split[1])
				_,p := searchList(list,n)
				if p != nil {
					removeItem(&list,n)
				}
				break
			case "?":
				n,_ := strconv.Atoi(split[1])
				_,p := searchList(list,n)
				if p != nil{
					fmt.Println("Il valore appartiene all'insieme")
				}else{
					fmt.Println("Il valore non appartiene all'insieme")
				}
				break
			case "c":
				fmt.Println("Numero elementi: ",countItems(list))
				break
			case "p":
				printList(list)
				break
			case "o":
				printReverseList(list)
				break
			case "d":
				p := list.head
				for p != nil {
					removeItem(&list,p.item)
					p = p.next
				}
				break
		}
	}
	fmt.Println("Terminato")
}

//Numero di elementi nella lista
func countItems(l linkedList) int {
	p := l.head
	var count int
	for p != nil{
		count++
		p = p.next
	}
	return count
}
/*
  Con una singola linea alloco memoria per il nuovo nodo
  memorizzo il dato nel nuovo nodo
  inserisco il nodo nella lista
*/
func newNode(val int) *listNode {
	return &listNode{val,nil}
}

/*
  Inserimento di un nuovo nodo in testa
  restituendo la linkedlist modificata, dovrò fare un assegnamento
  list = addNewNode(list,42) per esempio
*/
func addNewNode(l linkedList, val int) linkedList {
	node := newNode(val)
	node.next = l.head
	l.head = node
	return l
}

/*
  stampa della lista
  scorro la lista a partire dalla testa tramite un puntatore 
  ad ogni passo il puntatore punta al nodo successivo e lo stampa
*/
func printList(l linkedList) {
	p := l.head
	for p != nil {
		fmt.Print(p.item, " ")
		p = p.next
	}
	fmt.Println()
}

/*
  stampa della lista in modo inverso
  scorro la lista a partire dalla testa tramite un puntatore 
  ad ogni passo il puntatore punta al nodo successivo 
  metto ogni valore in una lista e la stampo al contrario
*/
func printReverseList(l linkedList) {
	p := l.head
	list := make([]int,countItems(l))
	var k int
	for p != nil {
		list[k] = p.item
		p = p.next
		k++
	}
	for i:= countItems(l)-1;i >= 0;i--{
		fmt.Print(" ",list[i])
	}
	fmt.Println()
}

/*
   ricerca di un nodo nella lista
   Scorro la lista a partire dalla testa cercando il nodo desiderato.
   Uso un puntatore e ad ogni passo punta al nodo successivo che si sta visitando.
   Quando viene trovato, viene restituito.
*/
func searchList(l linkedList, val int) (bool, *listNode) {
	p := l.head
	for p != nil {
		if p.item == val {
			return true, p
		}
		p = p.next
	}	
	return false, nil
}

/*
  Cancellazione di un nodo nella lista
  Prima cerco il nodo desiderato da cancellare
  poi verifico se il nodo non si trovi in testa,
  in quel caso punto la testa al nodo successivo.
  Non mi basta un solo puntatore, uso curr per puntare al nodo corrente
  prev per puntare al nodo precedente. Quando visito la testa prev è nil.
*/
func removeItem(l *linkedList, val int) bool {
	var curr, prev *listNode = l.head,nil
	for curr != nil {
		if curr.item == val {
			if prev == nil {
				l.head = l.head.next
			} else {
				prev.next = curr.next
			}
			return true
		}
		prev = curr
		curr = curr.next
	}
	return false
}

