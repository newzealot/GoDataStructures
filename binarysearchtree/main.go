package main

import "fmt"

type node struct {
	value int
	left  *node
	right *node
	count int
}

type binarySearchTree struct {
	root *node
}

type queue []node

func (q *queue) enqueue(n node) {
	*q = append(*q, n)
}

func (q *queue) dequeue() (node, error) {
	if len(*q) == 0 {
		return node{}, fmt.Errorf("empty queue")
	}
	a, b := (*q)[1:], (*q)[0]
	*q = a
	return b, nil
}

func (b *binarySearchTree) insert(v int) {
	newNode := &node{value: v}
	if b.root == nil {
		b.root = newNode
		return
	}
	currentNode := b.root
	for {
		if v == currentNode.value {
			currentNode.count++
			return
		} else if v < currentNode.value {
			if currentNode.left == nil {
				currentNode.left = newNode
				newNode.count++
				return
			}
			currentNode = currentNode.left
		} else {
			if currentNode.right == nil {
				currentNode.right = newNode
				newNode.count++
				return
			}
			currentNode = currentNode.right
		}
	}
}

func (b *binarySearchTree) contains(v int) bool {
	if b.root == nil {
		return false
	}
	currentNode := b.root
	for currentNode != nil {
		if v == currentNode.value {
			return true
		} else if v < currentNode.value {
			currentNode = currentNode.left
		} else {
			currentNode = currentNode.right
		}
	}
	return false
}

func (b *binarySearchTree) bfs() error {
	var q, visited queue
	var output []int
	q = make([]node, 0)
	visited = make([]node, 0)
	if b.root == nil {
		return fmt.Errorf("empty tree")
	}
	q = append(q, *b.root)
	for len(q) != 0 {
		v, err := q.dequeue()
		if err != nil {
			return err
		}
		visited = append(visited, v)
		if v.left != nil {
			q.enqueue(*v.left)
		}
		if v.right != nil {
			q.enqueue(*v.right)
		}
	}
	for _, v := range visited {
		output = append(output, v.value)
	}
	fmt.Println(output)
	return nil
}

func traverse(visited *queue, n *node, s string) {
	if s == "preorder" {
		visited.enqueue(*n)
	}
	if n.left != nil {
		traverse(visited, n.left, s)
	}
	if s == "inorder" {
		visited.enqueue(*n)
	}
	if n.right != nil {
		traverse(visited, n.right, s)
	}
	if s == "postorder" {
		visited.enqueue(*n)
	}
}

func (b *binarySearchTree) dfs(s string) error {
	var q, visited queue
	var output []int
	if b.root == nil {
		return fmt.Errorf("empty tree")
	}
	q = append(q, *b.root)
	v, err := q.dequeue()
	if err != nil {
		return err
	}
	traverse(&visited, &v, s)
	for _, v := range visited {
		output = append(output, v.value)
	}
	fmt.Println(output)
	return nil
}

func main() {
	var b binarySearchTree
	b.insert(10)
	b.insert(6)
	b.insert(15)
	b.insert(3)
	b.insert(8)
	b.insert(20)
	fmt.Println(b)
	fmt.Println(b.contains(1))
	fmt.Println(b.contains(5))
	fmt.Println(b.contains(11))
	fmt.Println(b.contains(18))
	b.bfs()
	b.dfs("preorder")
	b.dfs("postorder")
	b.dfs("inorder")
}
