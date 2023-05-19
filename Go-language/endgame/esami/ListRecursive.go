package main

import "fmt"

//struttura dati lista concatenata
type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

//aggiunge un nuovo nodo alla lista
//inserimento in testa alla lista costo O(1)
func insert(l LinkedList,val int) LinkedList{
	node := &Node{val,nil}
	node.next = l.head
	l.head = node
	return l
}

/*
 * La funzione f dato in ingresso il puntatore alla testa di una lista
 * e un intero k, stampa il contenuto del nodo in cui il risultato della chiamata ricorsiva sommata a 1 è uguale a k
 * e restituisce il numero di nodi che ha attraversato. 
*/
func f(p *Node,k int) int{
	var a int
	if p == nil{
		return 0
	}
	a = 1 + f(p.next,k)
	if a == k{
		fmt.Println(p.data)
	}
	return a
}

//stampa il contenuto della lista concatenata
//costo O(n) lineare
func printList(l LinkedList){
	h := l.head
	for h != nil{
		fmt.Print(h.data, " ")
		h = h.next
	}
	fmt.Println()
}

//main di test
func main(){
	l := LinkedList{nil}
	l = insert(l,7)
	l = insert(l,1)
	l = insert(l,0)
	l = insert(l,-2)
	l = insert(l,3)
	printList(l)
	fmt.Println("Chiamo f con k = 1")
	a := f(l.head,1)
	fmt.Println(a) //5
	fmt.Println("Chiamo f con k = 3")
	a = f(l.head,3)
	fmt.Println(a) //5
	fmt.Println("Chiamo f con k = 5")
	a = f(l.head,5)
	fmt.Println(a) //5
	fmt.Println("Chiamo f con k = 10")
	a = f(l.head,10)
	fmt.Println(a) //5
}
