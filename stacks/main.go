package main

import (
	"fmt"
)

type stackSlice []int

type node struct {
	value int
	next  *node
}

type stackLinkedList struct {
	head   *node
	tail   *node
	length int
}

func (s *stackSlice) add(v int) {
	*s = append(*s, v)
}

func (s *stackSlice) remove() (int, error) {
	l := len(*s)
	if len(*s) > 0 {
		a, b := (*s)[:l-1], (*s)[l-1]
		*s = a
		return b, nil
	}
	return 0, fmt.Errorf("empty stackSlice")
}

func (l *stackLinkedList) add(v int) {
	newNode := &node{value: v}
	if l.length == 0 {
		l.head = newNode
		l.tail = newNode
	} else {
		newNode.next = l.head
		l.head = newNode
	}
	l.length++
}

func (l *stackLinkedList) remove() (int, error) {
	if l.length == 0 {
		return 0, fmt.Errorf("empty stackLinkedList")
	}
	v := l.head.value
	l.head = l.head.next
	l.length--
	if l.head == nil {
		l.tail = nil
	}
	return v, nil
}

func main() {
	var s stackSlice
	s = make([]int, 0)
	s.add(1)
	fmt.Println(s)
	s.add(2)
	fmt.Println(s)
	s.add(3)
	fmt.Println(s)
	fmt.Println(s.remove())
	fmt.Println(s.remove())
	fmt.Println(s.remove())
	fmt.Println(s.remove())
	fmt.Println(s)
	var l stackLinkedList
	l.add(1)
	l.add(2)
	l.add(3)
	fmt.Println(l)
	fmt.Println(l.remove())
	fmt.Println(l.remove())
	fmt.Println(l.remove())
	fmt.Println(l.remove())
	fmt.Println(l)
}
