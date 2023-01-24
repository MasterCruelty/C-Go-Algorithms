package main

import "fmt"

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
