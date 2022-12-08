package main

import (
	"fmt"
)

/*
  *Modellazione:
  *I dipendenti si possono rappresentare come nodi dell'albero.
  *Le relazioni tra dipendenti(supervirore e subordinato) si possono rappresentare facilmente tramite la gerarchia padre/figlio
  *La radice dell'albero è il CEO dell'azienda e quindi supervisore di tutti ricorsivamente.
  *Scendendo nei rami dell'albero è possibile ricavare chi supervisione e chi è subordinato.
  
  *Implementazione struttura dati:
  *Albero implementato definendo una struttura tree contenente un riferimento alla radice definita come nodo.
  *La struttura nodo è definita con 3 riferimenti ai nodi figli sx,centrale e dx. Inoltre anche un campo dati di tipo stringa, conterrà il nome dipendente.
  *Infine la struttura nodo conterrà un riferimento al nodo padre per sapere il supervisore.
  
*/

//Rappresentazione in memoria di un albero binario
//una struct per definire la radice e un'altra per il nodo di un albero
type treeNode struct {
	left    *treeNode
	center  *treeNode
	right   *treeNode
	father  *treeNode
	item string
}

type tree struct {
	root *treeNode
}

/*
  Uso della struttura dati ad albero binario per rappresentare un'azienda con dipendenti e i loro supervisori/subordinati
*/
func main(){

	fmt.Println("Dipendenti: Anna =>{Bruno,Carlo,Daniela} Bruno =>{Enrico,Francesco} Francesco => Irene\n\n")
	//creo l'albero e popolo i nodi secondi i subordinati e i supervisori
	t := &tree{nil}
	t.root = &treeNode{nil,nil,nil,nil,"Anna"}
	t.root.left = newNode("Bruno",t.root)
	t.root.center = newNode("Carlo",t.root)
	t.root.right = newNode("Daniela",t.root)
	t.root.left.left = newNode("Enrico",t.root.left)
	t.root.left.right = newNode("Francesco",t.root.left)
	t.root.left.right.right = newNode("Irene",t.root.left.right)
	fmt.Println("Stampa albero appena costruito")
	stampaAlberoASommario(t.root,0)
	fmt.Print("\n\nSubordinati di Bruno: ")
	stampaSubordinati(t.root.left,true)
	fmt.Println()
	fmt.Println("Numero di dipendenti senza subordinati: ",quantiSenzaSubordinati(t.root))
	fmt.Println()
	fmt.Println("Il supervisore di Francesco è: ",supervisore(t.root.left.right).item)
	fmt.Println()
	fmt.Println("Stampa degli impiegati sopra Irene: ")
	stampaImpiegatiSopra(t.root.left.right.right)
	fmt.Println()
	fmt.Println("Stampa degli impiegati sopra Bruno: ")
	stampaImpiegatiSopra(t.root.left)
	fmt.Println()
	fmt.Println("Stampa degli impiegati sopra Enrico: ")
	stampaImpiegatiSopra(t.root.left.left)
	fmt.Println()
	fmt.Println("Elenco dei supervisori per ogni dipendente")
	stampaImpiegatiConSupervisore(t.root)

}


//funzione che stampa i subordinati di un dipendente
//Visito ricorsivamente prima i sottoalberi sinistri e centrale, poi ignoro la radice e stampo il sottoalbero destro
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

//Visita i nodi in ordine simmetrico e restituisce 1 quando il nodo non ha figli
//altrimenti 0 e continua finché non raggiunge la fine dell'albero, facendo prima il sottoalbero sinistro, centrale e poi destro.
func quantiSenzaSubordinati(node *treeNode) int{
	if node == nil{
		return 0
	}
	if node.left == nil && node.right == nil {
		return 1 + quantiSenzaSubordinati(node.left) + quantiSenzaSubordinati(node.center) + quantiSenzaSubordinati(node.right)
	}else{
		return 0 + quantiSenzaSubordinati(node.left) + quantiSenzaSubordinati(node.center) + quantiSenzaSubordinati(node.right)
	}
}

//Restituisco il riferimento al padre del nodo
func supervisore(node *treeNode) *treeNode{
	return node.father
}

//Controllo se il riferimento alpadre è nullo e quindi concludo
//altrimenti stampo il padre e ricorsivamente al padre dopo
func stampaImpiegatiSopra(node *treeNode) {
	if node.father == nil{
		return
	}
	fmt.Println(supervisore(node).item)
	stampaImpiegatiSopra(node.father)
}

//visita in ordine simmetrico tralasciando la radice e stampo i supervisori per ogni nodo
func stampaImpiegatiConSupervisore(node *treeNode) {
	if node == nil{
		return
	}
	stampaImpiegatiConSupervisore(node.left)
	stampaImpiegatiConSupervisore(node.center)
	if node.father != nil{
		fmt.Println(node.item + " ha come supervisore " + node.father.item)
	}
	stampaImpiegatiConSupervisore(node.right)
}

//Crea un nuovo nodo con il valore inserito in input e assegna null ai figli sx e dx
func newNode(val string,father *treeNode) *treeNode{
	return &treeNode{nil,nil,nil,father,val}
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
