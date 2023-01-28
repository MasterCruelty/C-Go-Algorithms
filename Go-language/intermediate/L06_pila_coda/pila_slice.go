package main

import "fmt"

/*
 * In una pila(Stack), gli elementi vengono inseriti e rimossi dalla stessa estremità(chiamata cima della pila) in un ordine LIFO.
 * Quindi l'ultimo elemento inserito è il primo ad essere rimosso.
 *
 * In una coda(Queue) gli elementi vengono inseriti dalla coda e rimossi dalla testa in un ordine FIFO, quindi il primo inserito è il primo ad essere rimosso.
 * Dunque la funzione Push e Enqueue di pila e coda hanno un comportamento simile ma l'effetto sulla struttura dati è differente.
*/

//implementazione della pila mediante slice
type Stack struct{
	items []int
}

//metodo push per inserire un elemento nella pila
func (s *Stack) Push(item int) {
	s.items = append(s.items,item)
}

//metodo pop per prelevare un elemento dalla pila eliminandolo da essa
func (s *Stack) Pop() int {
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

//metodo peek che restituisce l'ultimo elemento della pila
func (s *Stack) Peek() int {
	return s.items[len(s.items)-1]
}

//controlla se la pila è vuota
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

//main di test per la pila
func main(){
	s := Stack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println(s.Peek())
	fmt.Println(s.Pop())
	fmt.Println(s.Peek())
}
