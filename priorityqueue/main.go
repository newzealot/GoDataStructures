package main

import (
	"fmt"
	"math"
)

type node struct {
	value    string
	priority int
}

// based on min binary heap
type priorityQueue []node

func (p *priorityQueue) bubbleUp(currentNodeIndex int) {
	parentNodeIndex := int(math.Floor(float64(currentNodeIndex-1) / float64(2)))
	// while parentNode still within array
	for parentNodeIndex >= 0 && (*p)[currentNodeIndex].priority < (*p)[parentNodeIndex].priority {
		(*p)[currentNodeIndex], (*p)[parentNodeIndex] = (*p)[parentNodeIndex], (*p)[currentNodeIndex]
		currentNodeIndex = parentNodeIndex
		parentNodeIndex = int(math.Floor(float64(currentNodeIndex-1) / float64(2)))
	}
	return
}

func (p *priorityQueue) bubbleDown(currentNodeIndex int) {
	leftNodeIndex := (2 * currentNodeIndex) + 1
	rightNodeIndex := (2 * currentNodeIndex) + 2
	var smallestNodeIndex int
	// if leftNode still inside slice bounds
	// compare leftNode with currentNode
	if leftNodeIndex < len(*p) && (*p)[leftNodeIndex].priority < (*p)[currentNodeIndex].priority {
		smallestNodeIndex = leftNodeIndex
	} else {
		smallestNodeIndex = currentNodeIndex
	}
	// if rightNode still inside slice bounds
	// compare rightNode with leftNode
	if rightNodeIndex < len(*p) && (*p)[rightNodeIndex].priority < (*p)[leftNodeIndex].priority {
		smallestNodeIndex = rightNodeIndex
	}
	if smallestNodeIndex != currentNodeIndex {
		(*p)[currentNodeIndex], (*p)[smallestNodeIndex] = (*p)[smallestNodeIndex], (*p)[currentNodeIndex]
		p.bubbleDown(smallestNodeIndex)
	}
	return

}

func (p *priorityQueue) dequeue() (node, error) {
	if len(*p) == 0 {
		return node{}, fmt.Errorf("empty queue")
	}
	// remember the rootNode so you can return without it being changed
	rootNode := (*p)[0]
	// change firstNode value to lastNode value
	(*p)[0] = (*p)[len(*p)-1]
	// remove the last node
	*p = (*p)[:len(*p)-1]
	// bubbleDown the node such that it reaches its correct place
	p.bubbleDown(0)
	return rootNode, nil
}

func (p *priorityQueue) enqueue(n node) {
	*p = append(*p, n)
	p.bubbleUp(len(*p) - 1)
	return
}

func main() {
	p := make(priorityQueue, 0)
	p.enqueue(node{"common cold", 5})
	p.enqueue(node{"gunshot wound", 1})
	p.enqueue(node{"high fever", 4})
	p.enqueue(node{"broken arm", 2})
	p.enqueue(node{"glass in foot", 3})
	fmt.Println(p.dequeue())
	fmt.Println(p.dequeue())
	fmt.Println(p.dequeue())
	fmt.Println(p.dequeue())
	fmt.Println(p.dequeue())
	fmt.Println(p.dequeue())
}
