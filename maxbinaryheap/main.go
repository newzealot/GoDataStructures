package main

import (
	"fmt"
)

type maxBinaryHeap []int

func (m *maxBinaryHeap) insert(v int) {
	*m = append(*m, v)
	insertedIndex := len(*m) - 1
	parentIndex := (insertedIndex - 1) / 2
	for parentIndex >= 0 && (*m)[insertedIndex] > (*m)[parentIndex] {
		(*m)[insertedIndex], (*m)[parentIndex] = (*m)[parentIndex], (*m)[insertedIndex]
		insertedIndex = parentIndex
		parentIndex = (insertedIndex - 1) / 2
	}
	fmt.Println(m)
}

func (m *maxBinaryHeap) bubbleDown(i int) {
	leftIndex := 2*i + 1
	rightIndex := 2*i + 2
	var LargestValueIndex int
	if leftIndex <= len(*m)-1 && (*m)[leftIndex] > (*m)[i] {
		// leftIndex still in array and LeftValue bigger than current value
		LargestValueIndex = leftIndex
	} else {
		// current value is still the largest
		LargestValueIndex = i
	}
	if rightIndex <= len(*m)-1 && (*m)[rightIndex] > (*m)[LargestValueIndex] {
		// rightIndex still in array and RightValue bigger than LargestValue
		LargestValueIndex = rightIndex
	}
	if LargestValueIndex != i {
		(*m)[LargestValueIndex], (*m)[i] = (*m)[i], (*m)[LargestValueIndex]
		m.bubbleDown(LargestValueIndex)
	}
	return
}

func (m *maxBinaryHeap) maxExtract() (int, error) {
	if len(*m) == 0 {
		return 0, fmt.Errorf("empty heap")
	}
	// remember returnedValue
	returnedValue := (*m)[0]
	// make the root contain the last position
	(*m)[0] = (*m)[len(*m)-1]
	// cut off the last position
	*m = (*m)[:len(*m)-1]
	// bubble down
	m.bubbleDown(0)
	fmt.Println(returnedValue, m)
	return returnedValue, nil
}

func main() {
	m := maxBinaryHeap{41, 39, 33, 18, 27, 12}
	m.insert(55)
	m.maxExtract()
	m.maxExtract()
	m.maxExtract()
	m.maxExtract()
	m.maxExtract()
	m.maxExtract()
	m.maxExtract()

}
