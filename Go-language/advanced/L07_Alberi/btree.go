package main

import (
	"fmt"
	"math/rand"
)


//Rappresentazione in memoria di un albero binario
//una struct per definire la radice e un'altra per il nodo di un albero
//il tipo di "item" è arbitrario in funzione dell'informazione contenuta nei nodi
type treeNode struct {
	left  *treeNode
	right *treeNode
	item int
}

type tree struct {
	root *treeNode
}

//main d'esempio per testare le funzioni e i tipi struct per l'albero binario
func main(){

	//esempio di creazione albero e test funzioni preliminari
	fmt.Println("Albero binario per testare le funzioni:")
	t := &tree{nil} //albero vuoto
	t.root = &treeNode{nil,nil,78}
	//figli sx e dx della radice
	t.root.left = newNode(54)
	t.root.right = newNode(21)
	//discendente destro del figlio sinistro della radice + i suoi figli sx e dx
	t.root.left.right = newNode(90)
	t.root.left.right.left = newNode(19)
	t.root.left.right.right = newNode(95)
	//discendente sinistro del figlio destro della radice + il suo figlio sx
	t.root.right.left = newNode(16)
	t.root.right.left.left = newNode(5)
	//discendente destro del figlio destro della radice + i suoi figli sx e dx
	t.root.right.right = newNode(19)
	t.root.right.right.left = newNode(56)
	t.root.right.right.right = newNode(43)
	stampaAlberoASommario(t.root,0)
	//fine test funzioni preliminari alberi binari
	fmt.Println("Array con int casuali e albero binario costruito a partire da esso:")
	//genero interi casuali e memorizzo in una slice
	numeri := make([]int,11)
	for i:=0;i < len(numeri);i++{
		n := rand.Intn(100)
		numeri[i] = n
	}

	//costruisco un albero binario a partire dalla slice di interi
	fmt.Println(numeri)
	alb := arr2tree(numeri,0)
	stampaAlberoASommario(alb,0)

	//test ricerca inorder
	fmt.Println("cerco il nodo 47")
	if inOrderSearch(alb,47) != nil{
		fmt.Println("Trovato il nodo 47")
	}else{
		fmt.Println("non è presente il nodo 47")
	}

}

//Crea un nuovo nodo con il valore inserito in input e assegna null ai figli sx e dx
func newNode(val int) *treeNode{
	return &treeNode{nil,nil,val}
}

//A partire da un vettore costruisce un albero binario
func arr2tree(a []int,i int) (root *treeNode) {
	if i >= len(a) {
		return nil
	}
	tr := &tree{nil}
	tr.root = &treeNode{nil,nil,a[i]}
	tr.root.left = arr2tree(a,2*i+1)
	tr.root.right = arr2tree(a,2*i+2)
	return tr.root
}

//stampa dell'albero con radice e figli evidenziati
func stampaAlberoASommario(node *treeNode,spaces int) {
	for i := 0; i < spaces;i++{
		fmt.Print("  ")
	}
	fmt.Print("*")
	if node != nil{
		fmt.Println(node.item)
		if node.left != nil || node.right != nil {
			stampaAlberoASommario(node.left,spaces+1)
			stampaAlberoASommario(node.right,spaces+1)
		}
	}else{
		fmt.Println()
	}

}

//ricerca di un nodo usando la visita inorder
func inOrderSearch(node *treeNode,value int) *treeNode{
	if node == nil {
		return nil
	}
	result := inOrderSearch(node.left,value)
	if result != nil{
		return result
	}
	if node.item == value {
		return node
	}
	return inOrderSearch(node.right,value)
}

//visita albero in ordine simmetrico (prima sottoalbero sx, poi la radice, poi sottoalbero dx)
func inorder(node *treeNode){
	if node == nil {
		return
	}
	inorder(node.left)
	fmt.Print(node.item)
	inorder(node.right)
}

//visita albero in ordine anticipato(prima la radice, poi sotto albero sx e dx)
func preorder(node *treeNode) {
	if node == nil{
		return
	}
	fmt.Print(node.item)
	preorder(node.left)
	preorder(node.right)
}

//visita albero in ordine posticipato(prima sottoalbero sx, poi sottoalbero dx e infine la radice)
func postorder(node *treeNode){
	if node == nil{
		return
	}
	postorder(node.left)
	postorder(node.right)
	fmt.Print(node.item)
}

