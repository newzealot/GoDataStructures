package main

import (
	"fmt"
)

type node struct {
	value int
	next *node
}

type linkedlist struct {
	head *node
	tail *node
	length int
}
func (l *linkedlist) list() {
	currentNode := l.head
	if currentNode == nil {
		// 0 item in linkedlist
		fmt.Printf("Empty linkedlist\n")
		return
	}
	for currentNode.next != nil {
		fmt.Printf("%d, ",currentNode.value)
		currentNode = currentNode.next
	}
	fmt.Printf("%d\n",currentNode.value)
}

func (l *linkedlist) push(v int)  {
	newNode := &node{value: v}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		l.length++
		return
	}
	l.tail.next = newNode
	l.tail = newNode
	l.length++
	return
}

func (l *linkedlist) pop() (int, error) {
	var previousNode *node
	currentNode := l.head
	if currentNode == nil {
		// 0 item in linkedlist
		return 0, fmt.Errorf("cannot pop from empty linkedlist")
	} else if currentNode.next == nil {
		// Only 1 item in linkedlist
		l.head = nil
		l.tail = nil
		l.length = 0
		return currentNode.value, nil
	}
	for currentNode.next != nil {
		previousNode = currentNode
		currentNode = currentNode.next
	}
	previousNode.next = nil
	l.tail = previousNode
	l.length--
	return currentNode.value, nil
}

func (l *linkedlist) shift() (int,error) {
	currentNode := l.head
	if currentNode == nil {
		// 0 item in linkedlist
		return 0, fmt.Errorf("cannot shift from empty linkedlist")
	} else if currentNode.next == nil {
		l.tail = nil
	}
	l.head = currentNode.next
	l.length--
	return currentNode.value, nil
}

func (l *linkedlist) unshift(v int) {
	newNode := &node{value: v}
	if l.head == nil {
		//0 items in linkedlist
		l.head = newNode
		l.tail = newNode
		l.length = 1
		return
	}
	newNode.next = l.head
	l.head = newNode
	l.length++
	return
}

func (l *linkedlist) get(i int) (int, error) {
	if  l.length == 0 || i >= l.length || i < 0 {
		return 0, fmt.Errorf("index not found")
	}
	currentNode := l.head
	for j:=0;j<i;j++ {
		currentNode = currentNode.next
	}
	return currentNode.value, nil
}

func (l *linkedlist) set(i, v int) error {
	if l.length == 0 || i >= l.length || i < 0 {
		return fmt.Errorf("index not found")
	}
	currentNode := l.head
	for j:=0;j<i;j++ {
		currentNode = currentNode.next
	}
	currentNode.value = v
	return nil
}

func (l *linkedlist) insert(i,v int) error {
	if i > l.length || i < 0 {
		return fmt.Errorf("index not found")
	} else if i == 0 {
		l.unshift(v)
		return nil
	} else if i == l.length {
		l.push(v)
		return nil
	}
	newNode := &node{value: v}
	var previousNode *node
	currentNode := l.head
	for j:=0;j<i;j++{
		previousNode = currentNode
		currentNode = currentNode.next
	}
	previousNode.next = newNode
	newNode.next = currentNode
	l.length++
	return nil
}

func (l *linkedlist) remove(i int) error {
	if i >= l.length {
		return fmt.Errorf("index not found")
	} else if i == 0 {
		l.shift()
		return nil
	}
	var previousNode *node
	currentNode := l.head
	for j:=0;j<i;j++ {
		previousNode = currentNode
		currentNode = currentNode.next
	}
	if currentNode.next == nil {
		l.pop()
		return nil
	}
	previousNode.next = currentNode.next
	l.length--
	return nil
}

func (l *linkedlist) reverse() {
	if l.length == 0 || l.length == 1 {
		return
	}
	currentNode := l.head
	var previousNode *node
	var nextNode *node
	for i:=0;i<l.length;i++ {
		nextNode = currentNode.next
		currentNode.next = previousNode
		previousNode = currentNode
		currentNode = nextNode
	}
	l.head, l.tail = l.tail, l.head
	return

}

func main() {
	var l linkedlist
	l.push(5)
	l.push(9)
	l.push(3)
	l.list()
	fmt.Printf("length: %d\n", l.length)
	fmt.Println(l.pop())
	l.list()
	fmt.Println(l.pop())
	l.list()
	fmt.Println(l.pop())
	l.list()
	fmt.Println(l.pop())
	l.list()
	fmt.Printf("length: %d\n", l.length)
	l.push(5)
	l.push(9)
	l.push(3)
	l.list()
	fmt.Printf("length: %d\n", l.length)
	fmt.Println(l.shift())
	l.list()
	fmt.Println(l.shift())
	l.list()
	fmt.Println(l.shift())
	l.list()
	fmt.Println(l.shift())
	fmt.Printf("length: %d\n", l.length)
	l.unshift(5)
	l.unshift(9)
	l.unshift(3)
	l.list()
	fmt.Printf("length: %d\n", l.length)
	fmt.Println(l.get(0))
	fmt.Println(l.get(1))
	fmt.Println(l.get(2))
	fmt.Println(l.get(3))
	fmt.Println(l.set(0,8))
	l.list()
	fmt.Println(l.set(1,8))
	l.list()
	fmt.Println(l.set(2,8))
	l.list()
	fmt.Println(l.set(3,8))
	l.list()
	fmt.Println(l.insert(3,9))
	l.list()
	fmt.Println(l.insert(2,9))
	l.list()
	fmt.Println(l.insert(1,9))
	l.list()
	fmt.Println(l.insert(0,9))
	l.list()
	fmt.Printf("length: %d\n", l.length)
	fmt.Println(l.remove(5))
	l.list()
	fmt.Println(l.remove(3))
	l.list()
	fmt.Println(l.remove(1))
	l.list()
	fmt.Println(l.remove(3))
	l.list()
	fmt.Println(l.remove(0))
	l.list()
	fmt.Println(l.remove(0))
	l.list()
	fmt.Println(l.remove(0))
	l.list()
	fmt.Printf("length: %d\n", l.length)
	l.push(1)
	l.push(2)
	l.push(3)
	l.push(4)
	l.push(5)
	l.list()
	fmt.Printf("length: %d\n", l.length)
	l.reverse()
	l.list()
	fmt.Println(l.head)
	fmt.Println(l.tail)
}
