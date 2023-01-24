package main

import(
	"bufio"
	"fmt"
	"os"
	"strings"
)

//definizione tipo pila tramite lista

type listNode struct {
	next *listNode
	item string
}

type stack struct{
	head *listNode
}

func newNode(item string) *listNode{
	return &listNode{nil,item}
}

func (list *stack) push(item string) {
	newNode := newNode(item)
	newNode.next = list.head
	list.head = newNode
}

func (list *stack) pop() string{
	node := list.head
	list.head = node.next
	return node.item
}

func (list stack) isEmpty() bool {
	if list.head == nil {
		return true
	}
	return false
}

func (list stack) print(){
	fmt.Print("[")
	var node *listNode = list.head
	for node != nil {
		fmt.Print(node.item," " )
		node = node.next
	}
	fmt.Println("]")
}

/*
  * Pila definita come struttura dati astratta(tipo 'stack') 
  * metodi generali pop, push, etc.etc.

  * Pila usata come struttura dati astratta
  * nel main dichiaro una variabile stack di tipo stack
  * la variabile stack viene modificata solo mediante l'uso di push e pop


  * Pila implementata mediante lista concatenata
  * il metodo push effettua un inserimento in testa
  * il metodo pop effettua una rimozione in testa


  todo: 
  se riceve una pila vuota c'è errore sul metodo pop in checkTags
  
  soluzione: 
  la funzione chiamante deve fare un controllo sulla pila vuota 
  oppure si aggiunge un controllo usando isEmpty

  miglioramenti:
  completare strttura dati astratta con funzione top(tipica per le pile)
*/



func main(){
	fmt.Print("Inserisci documento html(una riga): ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	doc := scanner.Text()
	benFormato, errPos, tag_rimasti := checkTags(doc)
	if benFormato {
		fmt.Println("Il documento è ben formato")
	} else{
		fmt.Println("Il documento NON è ben formato")
		if len(tag_rimasti) > 0 {
			fmt.Println("\ttag non chiuso:",tag_rimasti)
		} else{
			fmt.Println("\tPosizione errore:",errPos)
		}
	}
}

func checkTags(doc string) (bool,int,[]string) {
	var stack stack = stack{nil}
	tags := strings.Split(doc," ")
	for i, t := range tags {
		if !strings.Contains(t,"/") {
			stack.push(t)
		}else if strings.Contains(t,"/") {
			openingTag := stack.pop()
			if openingTag[1:] != t[2:] {
				return false, i + 1, nil
			}
		}
	}
	if stack.isEmpty() {
		return true,-1,nil
	}else {
		var tag_rimasti = make([]string,0)
		for !stack.isEmpty() {
			tag_rimasti = append(tag_rimasti,stack.pop())
		}
		for i, j := 0, len(tag_rimasti)-1;i<j;i,j=i+1,j-1 {
			tag_rimasti[i],tag_rimasti[j] = tag_rimasti[j],tag_rimasti[i]
		}
		return false,-1,tag_rimasti
	}
}
