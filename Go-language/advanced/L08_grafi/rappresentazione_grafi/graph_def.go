package main


//varie rappresentazioni su macchina della struttura dati dei grafi


/*
  *Implementazione con struttura vertice
  *Utile se ai vertici del grafo vi sono informazioni articolate
  *Simile all'implementazione degli alberi binari con figlio sx e dx, in questo caso il puntatore è ai vertici vicini
  *Per rappresentare il grafo è sufficiente una mappa che associa, ad ogni chiave, il suo vertice.
  *Nell'albero binario bastava avere la radice, qui ci serve tutto l'insieme dei vertici.
*/
type vertice struct{
	item string
	key int
	adj []*vertice
}

type graph map[int]*vertice

/*********************

/*
  *Riciclo dell'implementazione delle liste concatenate
  *La slice del grado conterrà liste e non puntatori a nodi.
*/

type listNode struct {
	item int
	next *listNode
}

type linkedList struct {
	head *listNode
}

//func addNewNode(l *linkedList, n int)

type graph2 struct {
	n int
	adij []linkedList
}

/***********************

/*
  *Implementazione alla lettera
*/

type graph3 struct{
	n int
	adj []*listNode2
}

type listNode2 struct {
	item int
	next *listNode2
}
