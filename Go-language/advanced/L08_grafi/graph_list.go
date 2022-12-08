package main

import "fmt"


//riciclo l'implementazione delle liste concatenate

type listNode struct {
	item int
	next *listNode
}

type linkedList struct {
	head *listNode
}

func newNode(val int) *listNode{
	return &listNode{val,nil}
}

func addNewNode(l *linkedList,val int) {
	node := newNode(val)
	node.next = l.head
	l.head = node
}

func printList(l *linkedList) {
	p := l.head
	for p != nil {
		fmt.Print(p.item)
		if p.next != nil {
			fmt.Print(" ")
		}
		p = p.next
	}
	fmt.Println()
}

/*******************************************************/
//definizione del tipo grafo
type grafo struct {
	n int //numero di vertici
	adiacenti []linkedList //vettore di liste di adiacenza
}


//Restituisce l'indirizzo di un nuovo grafo con n nodi
func nuovoGrafo(n int) *grafo {
	return &grafo{n,make([]linkedList,n)}
}

//Lettura grafo da standard input
//prima un intero n per il numero di vertici e poi a coppie gli archi con due vertici per riga
//restituisce false se ci sono errori
//si considerano grafi non orientati
func leggiGrafo() *grafo{
	fmt.Println("Inserisci il numero di vertici")
	var n int
	fmt.Scan(&n)
	g := nuovoGrafo(n)

	fmt.Println("Inserisci gli archi del grafo, uno per riga, nel formato 'v1 v2'")
	var v1,v2 int
	_, err := fmt.Scan(&v1,&v2)

	for err == nil{
		//leggo i vertici
		addNewNode(&g.adiacenti[v1],v2)
		addNewNode(&g.adiacenti[v2],v1)
		stampaGrafo(g)
		_,err = fmt.Scan(&v1,&v2)
	}
	return g
}

//stampa del grafo
func stampaGrafo(g *grafo) {
	fmt.Printf("Il grafo ha %d nodi\n", g.n)
	for i,lista :=range g.adiacenti{
		fmt.Print(i,": ")
		printList(&lista)
	}
}

//main semplice per testare grafo implementato come array di liste di adiacenza
func main() {
	g := leggiGrafo()
	stampaGrafo(g)
}
