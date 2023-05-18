package main

import(
	"fmt"
	"bufio"
	"strconv"
	"os"
)

/*
 * struttura dati: 
 * lista doppiamente concatenata circolare.
 * puntatore alla testa e alla coda, ogni nodo ha un puntatore al successivo e al precedente.
*/
type Node struct{
	data int
	next *Node
	prev *Node
}

type circNode struct {
	head *Node
	tail *Node
}

/*
 * Aggiunge un nodo alla lista.
 * Valutazione complessità: O(1)
*/
func insert(dl *circNode, val int) *circNode{
	newNode := &Node{val,nil,nil}
	if dl.head == nil{
		dl.head = newNode
	}else{
		dl.tail.next = newNode
		newNode.prev = dl.tail
		newNode.next = dl.head
		dl.head.prev = newNode
	}
	dl.tail = newNode
	return dl
}

/*
 * Stampa tutta la lista a partire dal nodo di contenuto 0.
 * Valutazione complessità: O(n) + O(n) = O(n)
 * complessità lineare, nel primo for cerco il nodo di valore zero, 
 * nel secondo for stampo tutta la lista scorrendola.
*/
func stampaDaZero(dl *circNode) {
	p := dl.head
	for {
		if p.data == 0 {
			break
		}
		p = p.next
	}
	for {
		fmt.Print(p.data," ")
		p = p.next
		if p.data == 0{
			break
		}
	}
}

func stampaAvanti(c *circNode){
	current := c.head
	current = current.next
	fmt.Println("Stampa andando dalla testa alla coda")
	fmt.Printf("%d ",c.head.data)
	for current != c.head{
		fmt.Printf("%d ",current.data)
		current = current.next
	}
	fmt.Println()
}

func stampaIndietro(c *circNode) {
	current := c.head
	current = current.prev
	fmt.Println("Stampa andando dalla coda alla testa")
	for current != c.head{
		fmt.Printf("%d ",current.data)
		current = current.prev
	}
	fmt.Printf("%d ",c.head.data)
	fmt.Println()
}

/*
 * funzione che sposta il nodo puntato dato in ingresso dl
 * all'interno della lista circolare in base al suo valore.
 * Se è positivo viene spostato in avanti del suo valore.
 * Se è negativo viene spostato indietro del suo valore.
 * Esempio con la lista 3 -2 0 1 7:
 * Se sposto il nodo 3 la lista diventa: -2 0 1 3 7 
 * Se sposto il nodo -2 la lista diventa: 3 0 1 -2 7
 * Valutazione complessità: O(n) lineare
*/
func (dl *circNode) sposta(p *Node) {
	data := p.data
	if data == 0{
		return
	}
	//Tolgo il nodo da spostare dalla sua posizione attuale
	prev := p.prev
	next := p.next
	prev.next = next
	next.prev = prev
	//mi salvo il puntatore al nodo da spostare
	move := p
	if data > 0{
		for i:= 0; i <= data;i++ {
			move = move.next
		}
		//inserisco il nodo nella posizione nuova
		movePrev := move.prev
		movePrev.next = p
		move.prev = p
		p.prev = movePrev
		p.next = move
	}else{
		for i := 0;i > data ;i--{
			move = move.prev
		}
		//inserisco il nodo nella posizione nuova
		movePrev := move.prev
		movePrev.next = p
		move.prev = p
		p.prev = movePrev
		p.next = move
	}
	//se necessario aggiorno testa e/o coda
	if dl.head == move{
		dl.head = p
	}
	if dl.tail == move{
		dl.tail = p
	}
}


//main di test per le funzioni
func main(){
	c := &circNode{head:nil,tail:nil}
	c = insert(c,3)
	c = insert(c,-2)
	c = insert(c,0)
	c = insert(c,1)
	c = insert(c,7)
	//stampa della lista andando avanti
	stampaAvanti(c)
	//stampa della lista andando indietro
	stampaIndietro(c)
	//test funzione stampa da zero
	fmt.Println("Stampa partendo dal nodo di valore zero")
	stampaDaZero(c)
	fmt.Println()
	//inserimento da stdin e test funzioni
	c2 := &circNode{head:nil,tail:nil}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "exit"{
			break
		}
		val,_ := strconv.Atoi(scanner.Text())
		//aggiungo il valore letto alla lista circolare
		c2 = insert(c2,val)
	}
	//test sposta con nodo 3
	fmt.Println("Test funzione sposta con nodo di valore 3")
	c2.sposta(c2.head)
	stampaAvanti(c2)
	//test sposta con nodo -2
	fmt.Println("Test funzione sposta con nodo di valore -2")
	c.sposta(c.head.next)
	stampaAvanti(c)
}

