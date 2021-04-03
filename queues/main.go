package main

import "fmt"

type queueSlice []int

type node struct {
	value int
	next  *node
}

type queueLinkedList struct {
	head   *node
	tail   *node
	length int
}

func (q *queueLinkedList) enqueue(v int) {
	newNode := &node{value: v}
	if q.length == 0 {
		q.head = newNode
		q.tail = newNode
	} else {
		q.tail.next = newNode
		q.tail = newNode
	}
	q.length++
	return
}

func (q *queueLinkedList) dequeue() (int, error) {
	if q.length == 0 {
		return 0, fmt.Errorf("empty queue")
	}
	v := q.head.value
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}
	q.length--
	return v, nil
}

func (q *queueSlice) enqueue(v int) {
	*q = append(*q, v)
}

func (q *queueSlice) dequeue() (int, error) {
	if len(*q) == 0 {
		return 0, fmt.Errorf("empty queue")
	}
	a, b := (*q)[1:], (*q)[0]
	*q = a
	return b, nil
}
func main() {
	var q queueSlice
	q = make([]int, 0)
	q.enqueue(1)
	q.enqueue(2)
	q.enqueue(3)
	fmt.Println(q)
	fmt.Println(q.dequeue())
	fmt.Println(q)
	fmt.Println(q.dequeue())
	fmt.Println(q)
	fmt.Println(q.dequeue())
	fmt.Println(q)
	fmt.Println(q.dequeue())
	fmt.Println(q)
	var l queueLinkedList
	l.enqueue(1)
	l.enqueue(2)
	l.enqueue(3)
	fmt.Println(l)
	fmt.Println(l.dequeue())
	fmt.Println(l)
	fmt.Println(l.dequeue())
	fmt.Println(l)
	fmt.Println(l.dequeue())
	fmt.Println(l)
	fmt.Println(l.dequeue())
	fmt.Println(l)
}
