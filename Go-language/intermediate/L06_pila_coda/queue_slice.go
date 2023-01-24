package main

import "fmt"

//definizione di una coda tramite slice di interi
type Queue struct {
	items []int
}

//Inserisce un elemento in fondo alla coda
func (q *Queue) Enqueue(item int) {
	q.items = append(q.items,item)
}

//Preleva il primo elemento della coda eliminandolo da essa
func (q *Queue) Dequeue() int{
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

//Restituisce il primo elemento della coda
func (q *Queue) Peek() int{
	return q.items[0]
}


func main(){
	q:=Queue{}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	fmt.Println(q.Peek()) //stampa 1
	fmt.Println(q.Dequeue()) //stampa 1
	fmt.Println(q.Peek()) //stampa 2
}
