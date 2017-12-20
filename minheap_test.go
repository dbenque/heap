package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestHeapAdd(t *testing.T) {
	h := NewHeap(5)
	h.Add(5)
	h.Add(8)
	h.Add(1)
	h.Add(3)
	h.Add(10)

	str := fmt.Sprintf("%v", h.array)
	if str != "[0 1 3 5 8 10]" {
		t.Fatalf("Not expecting: %s\n", str)
	}
}

func TestHeapPop(t *testing.T) {
	h := NewHeap(5)
	h.Add(5)
	h.Add(8)
	h.Add(1)
	h.Add(3)
	h.Add(10)

	if v, _ := h.Pop(); v != 1 {
		t.Fatalf("Not expecting: %v but 1\n", v)
	}
	if v, _ := h.Pop(); v != 3 {
		t.Fatalf("Not expecting: %v but 3\n", v)
	}
	if v, _ := h.Pop(); v != 5 {
		t.Fatalf("Not expecting: %v but 5\n", v)
	}
	if v, _ := h.Pop(); v != 8 {
		t.Fatalf("Not expecting: %v but 8\n", v)
	}
	if v, _ := h.Pop(); v != 10 {
		t.Fatalf("Not expecting: %v but 10\n", v)
	}
	if _, e := h.Pop(); e == nil {
		t.Fatalf("Error was expected\n")
	}
}

func TestHeapRandom(t *testing.T) {
	c := uint(5000)
	h := NewHeap(c)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < int(c); i++ {
		h.Add(byte(r.Intn(256)))
	}
	last := byte(0)
	for i := 0; i < int(c); i++ {
		v, _ := h.Pop()
		if v < last {
			t.Fatalf("Super min %d, %d<%d\n", i, v, last)
		}
		last = v
	}
}
