package main

type listNode struct {
	item int
	next *listNode
}

type linkedListWithTail struct {
	head,tail *listNode
}

func newNode(val int) *listNode {
	return &listNode{val,nil}
}

func addNewNodeAtEnd( l *linkedListWithTail,val int) {
	if l.tail == nil {
		l.tail = newNode(val)
		l.head = l.tail
	} else{
		l.tail.next = newNode(val)
		l.tail = l.tail.next
	}
}

func main(){
	return
}
