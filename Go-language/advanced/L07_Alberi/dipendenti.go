package main

import (
	"fmt"
)


//Rappresentazione in memoria di un albero binario
//una struct per definire la radice e un'altra per il nodo di un albero
type treeNode struct {
	left  *treeNode
	center  *treeNode
	right *treeNode
	item string
}

type tree struct {
	root *treeNode
}

/*
  Uso della struttura dati ad albero binario per rappresentare un'azienda con dipendenti e i loro supervisori/subordinati
  
  todo:
  dato un certo dipendente, stampare il suo supervisore
  dato un certo dipendente, stampare la lista dei dipendenti che si trovano sopra di lui fino alla radice
  stampa dell'elenco di tutti i dipendenti(non importa l'ordine) indicando per ciascuno chi è il supervisore.
*/
func main(){

	fmt.Println("Dipendenti: Anna =>{Bruno,Carlo,Daniela} Bruno =>{Enrico,Francesco} Francesco => Irene\n\n")
	//creo l'albero e popolo i nodi secondi i subordinati e i supervisori
	t := &tree{nil}
	t.root = &treeNode{nil,nil,nil,"Anna"}
	t.root.left = newNode("Bruno")
	t.root.center = newNode("Carlo")
	t.root.right = newNode("Daniela")
	t.root.left.left = newNode("Enrico")
	t.root.left.right = newNode("Francesco")
	t.root.left.right.right = newNode("Irene")
	fmt.Println("Stampa albero appena costruito")
	stampaAlberoASommario(t.root,0)
	fmt.Print("\n\nSubordinati di Bruno: ")
	stampaSubordinati(t.root.left,true)
	fmt.Println()
	fmt.Println("Numero di dipendenti senza subordinati: ",countDipendentiSenzaSubordinati(t.root))

}


//funzione che stampa i subordinati di un dipendente(ovvero i sottoalberi sx e dx)
func stampaSubordinati(node *treeNode,check_node bool) {
	if node == nil{
		return
	}
	stampaSubordinati(node.left,false)
	stampaSubordinati(node.center,false)
	if !check_node{
		fmt.Print(node.item + "; ")
	}
	stampaSubordinati(node.right,false)
}

func countDipendentiSenzaSubordinati(node *treeNode) int{
	if node == nil{
		return 0
	}
	if node.left == nil && node.right == nil {
		return 1 + countDipendentiSenzaSubordinati(node.left) + countDipendentiSenzaSubordinati(node.center) + countDipendentiSenzaSubordinati(node.right)
	}else{
		return 0 + countDipendentiSenzaSubordinati(node.left) + countDipendentiSenzaSubordinati(node.center) + countDipendentiSenzaSubordinati(node.right)
	}
}
//Crea un nuovo nodo con il valore inserito in input e assegna null ai figli sx e dx
func newNode(val string) *treeNode{
	return &treeNode{nil,nil,nil,val}
}


//stampa dell'albero con radice e figli evidenziati
func stampaAlberoASommario(node *treeNode,spaces int) {
	for i := 0; i < spaces;i++{
		fmt.Print("  ")
	}
	fmt.Print("*")
	if node != nil{
		fmt.Println(node.item)
		if node.left != nil || node.right != nil || node.center != nil{
			stampaAlberoASommario(node.left,spaces+1)
			stampaAlberoASommario(node.center,spaces+1)
			stampaAlberoASommario(node.right,spaces+1)
		}
	}else{
		fmt.Println()
	}

}
