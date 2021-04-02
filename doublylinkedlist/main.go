package main

import "fmt"

type node struct {
	value    int
	next     *node
	previous *node
}

type doublylinkedlist struct {
	head   *node
	tail   *node
	length int
}

func (d *doublylinkedlist) push(v int) {
	newNode := &node{value: v}
	if d.length == 0 {
		d.head = newNode
		d.tail = newNode
		d.length++
		return
	}
	currentNode := d.tail
	currentNode.next = newNode
	newNode.previous = currentNode
	d.tail = newNode
	d.length++
}

func (d *doublylinkedlist) pop() (int, error) {
	if d.length == 0 {
		return 0, fmt.Errorf("empty linkedlist")
	}
	currentNode := d.tail
	if d.length == 1 {
		d.head = nil
		d.tail = nil
	} else {
		previousNode := currentNode.previous
		previousNode.next = nil
		d.tail = previousNode
	}
	d.length--
	return currentNode.value, nil
}

func (d *doublylinkedlist) list() {
	if d.length == 0 {
		fmt.Println("empty linkedlist")
		return
	}
	currentNode := d.head
	for i := 0; i < d.length; i++ {
		fmt.Printf("%d ", currentNode.value)
		currentNode = currentNode.next
	}
	fmt.Println()
	return
}

func (d *doublylinkedlist) shift() (int, error) {
	if d.length == 0 {
		return 0, fmt.Errorf("empty linkedlist")
	}
	currentNode := d.head
	var nextNode *node
	if d.length == 1 {
		d.head = nil
		d.tail = nil
	} else {
		nextNode = currentNode.next
		nextNode.previous = nil
		d.head = nextNode
	}
	d.length--
	return currentNode.value, nil
}

func (d *doublylinkedlist) unshift(v int) {
	newNode := &node{value: v}
	if d.length == 0 {
		d.head = newNode
		d.tail = newNode
		d.length++
		return
	}
	newNode.next = d.head
	d.head.previous = newNode
	d.head = newNode
	d.length++
	return
}

func (d *doublylinkedlist) get(i int) (int, error) {
	if i < 0 || i >= d.length || d.length == 0 {
		return 0, fmt.Errorf("index not found")
	} else if i == 0 {
		return d.head.value, nil
	} else if i == d.length-1 {
		return d.tail.value, nil
	} else if i < d.length/2 {
		currentNode := d.head
		for j := 0; j < i; j++ {
			currentNode = currentNode.next
		}
		return currentNode.value, nil
	} else {
		currentNode := d.tail
		for j := d.length - 1; j > i; j-- {
			currentNode = currentNode.previous
		}
		return currentNode.value, nil
	}
}

func (d *doublylinkedlist) set(i, v int) error {
	if i < 0 || i > d.length || d.length == 0 {
		return fmt.Errorf("index not found")
	} else if i == 0 {
		d.head.value = v
		return nil
	} else if i == d.length-1 {
		d.tail.value = v
		return nil
	} else if i < d.length/2 {
		currentNode := d.head
		for j := 0; j < i; j++ {
			currentNode = currentNode.next
		}
		currentNode.value = v
		return nil
	} else {
		currentNode := d.tail
		for j := d.length - 1; j > i; j-- {
			currentNode = currentNode.previous
		}
		currentNode.value = v
		return nil
	}
}

func (d *doublylinkedlist) insert(i, v int) error {
	if i < 0 || i > d.length {
		return fmt.Errorf("cannot insert at index")
	} else if i == 0 {
		d.unshift(v)
		return nil
	} else if i == d.length {
		d.push(v)
		return nil
	}
	var currentNode *node
	if i < d.length/2 {
		currentNode = d.head
		for j := 0; j < i; j++ {
			currentNode = currentNode.next
		}
	} else {
		currentNode = d.tail
		for j := d.length - 1; j > i; j-- {
			currentNode = currentNode.previous
		}
	}
	newNode := &node{value: v}
	previousNode := currentNode.previous
	previousNode.next = newNode
	newNode.previous = previousNode
	newNode.next = currentNode
	currentNode.previous = newNode
	d.length++
	return nil
}

func (d *doublylinkedlist) remove(i int) error {
	if i < 0 || i >= d.length || d.length == 0 {
		return fmt.Errorf("cannot find index")
	} else if i == 0 {
		d.shift()
		return nil
	} else if i == d.length-1 {
		d.pop()
		return nil
	}
	var currentNode *node
	if i < d.length/2 {
		currentNode = d.head
		for j := 0; j < i; j++ {
			currentNode = currentNode.next
		}
	} else {
		currentNode = d.tail
		for j := d.length - 1; j > i; j-- {
			currentNode = currentNode.previous
		}
	}
	previousNode := currentNode.previous
	nextNode := currentNode.next
	previousNode.next = nextNode
	nextNode.previous = previousNode
	d.length--
	return nil
}

func main() {
	var d doublylinkedlist
	d.push(1)
	d.push(2)
	d.push(3)
	d.list()
	fmt.Println(d.pop())
	d.list()
	fmt.Println(d.pop())
	d.list()
	fmt.Println(d.pop())
	d.list()
	fmt.Printf("length: %d\n", d.length)
	d.push(1)
	d.push(2)
	d.push(3)
	d.list()
	fmt.Println(d.shift())
	d.list()
	fmt.Println(d.shift())
	d.list()
	fmt.Println(d.shift())
	d.list()
	fmt.Printf("length: %d\n", d.length)
	d.unshift(1)
	d.list()
	d.unshift(2)
	d.list()
	d.unshift(3)
	d.list()
	d.unshift(4)
	d.list()
	d.unshift(5)
	d.list()
	fmt.Printf("length: %d\n", d.length)
	fmt.Println(d.get(0))
	fmt.Println(d.get(1))
	fmt.Println(d.get(2))
	fmt.Println(d.get(3))
	fmt.Println(d.get(4))
	d.set(0, 1)
	d.list()
	d.set(1, 2)
	d.list()
	d.set(2, 3)
	d.list()
	d.set(3, 4)
	d.list()
	d.set(4, 5)
	d.list()
	fmt.Printf("length: %d\n", d.length)
	d.insert(4, 9)
	d.list()
	d.insert(2, 9)
	d.list()
	fmt.Printf("length: %d\n", d.length)
	d.remove(2)
	d.list()
	d.remove(4)
	d.list()
	fmt.Printf("length: %d\n", d.length)
}
