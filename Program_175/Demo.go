package main

import "fmt"

type Node struct { 
	prev *Node
	next *Node
	key interface{}
}

type List struct {
	head *Node
	tail *Node
}

func(L *List) Insert(key interface{}) {
	list := &Node{
		next: L.head,
		key: key,
	}

	if L.head != nil {
		L.head.prev = list
	}
	L.head = list

	l := L.head

	for l.next != nil {
		l = l.next
	}
	L.tail = l
}

func (l *List) Display() {
	list := l.head
	for list != nil {
		fmt.Printf("%v ->",list.key)
		list = list.next
	}
	fmt.Println()
}

func Display(list *Node) {
	for list != nil {
		fmt.Printf("%v->",list.key)
		list = list.next 
	}
	fmt.Println()
}

func (l *List) Len() int {
	iCnt := 0
	list := l.head
	for list != nil {
		list = list.next
		iCnt++
	}

	return iCnt
}
func ShowBackwards(list *Node) {
	for list != nil {
		fmt.Printf("%v <-",list.key)
		list = list.prev
	}
	fmt.Println()
}
func (l *List) Remove() {
	temp := l.head

	if temp != nil {
		temp = temp.next
	}
	temp = nil
	fmt.Println(temp)
}

func (l *List) Reverse() {
	curr := l.head
	var prev *Node
	l.tail = l.head

	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	l.head = prev
	Display(l.head)
}

func main() { 
	ret := 0
	link := List{}
	link.Insert(5)
	link.Insert(9)
	link.Insert(13)
	link.Insert(22)
	link.Insert(28)
	link.Insert(36)

	fmt.Println("Jay Ganesh........")
	fmt.Printf("Head: %v\n",link.head.key)
	fmt.Printf("Tail: %v\n",link.tail.key)
	link.Display()
	fmt.Printf("Head: %v\n",link.head.key)
	fmt.Printf("Tail: %v\n",link.tail.key)
	link.Reverse()
	ret = link.Len()
	link.Remove()
	fmt.Println("Number of Elements in Linked List are :",ret)
}