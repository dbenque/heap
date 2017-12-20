package main

import "fmt"

//MinHeap properties
//1 - A complete binary tree is a tree in which each level has all of its nodes
//2 - In a heap, for every node x with parent p, the key in pp is smaller than or equal to the key in x
type MinHeap struct {
	size  uint
	array []byte
}

func (h *MinHeap) Add(b byte) error {
	if h.size+1 >= uint(len(h.array)) {
		return fmt.Errorf("Capacity exceeded")
	}
	h.size++
	h.array[h.size] = b
	h.up(h.size)
	return nil
}

func (h *MinHeap) up(i uint) {
	if i == 1 {
		return
	}
	p := i / 2
	if h.array[p] > h.array[i] {
		h.array[p], h.array[i] = h.array[i], h.array[p]
		h.up(p)
	}
	return
}

func (h *MinHeap) Pop() (byte, error) {
	if h.size == 0 {
		return 0, fmt.Errorf("Empty Heap")
	}

	ret := h.array[1]
	h.array[1] = h.array[h.size]
	h.array[h.size] = 0
	h.size--
	h.down(1)
	return ret, nil
}

func (h *MinHeap) down(i uint) {

	if 2*i > h.size { // i is already a leaf
		return
	}
	if 2*i == h.size { // i has only one child
		if h.array[i] > h.array[2*i] {
			h.array[i], h.array[2*i] = h.array[2*i], h.array[i]
		}
		return
	}

	minChildIndex := 2 * i
	if h.array[minChildIndex+1] < h.array[minChildIndex] {
		minChildIndex++
	}

	if h.array[i] > h.array[minChildIndex] {
		h.array[i], h.array[minChildIndex] = h.array[minChildIndex], h.array[i]
		h.down(minChildIndex)
	}
	return
}

func NewHeap(capacity uint) *MinHeap {
	return &MinHeap{size: 0, array: make([]byte, capacity+1)}
}