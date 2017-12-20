package main

import (
	"math/rand"
	"reflect"
	"sync"
	"testing"
	"time"
)

func TestFifoEmpty(t *testing.T) {
	f, _ := NewFifo(2)
	if !f.Empty() {
		t.Fatalf("Should be empty as it is new")
	}

	f.Push(1)
	f.Push(2)
	f.Pop()
	f.Push(5)
	f.Pop()
	f.Pop()

	if !f.Empty() {
		t.Fatalf("Should be empty push-pop=0")
	}
}

func TestFifo(t *testing.T) {
	c := 5
	s := 100
	f, _ := NewFifo(c)
	r := rand.New(rand.NewSource(time.Now().Unix()))
	in := make([]byte, s)
	out := make([]byte, s)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		index := 0
		for {
			b := f.Pop()
			out[index] = b
			index++
			if index == s {
				break
			}
		}
	}()

	for i := 0; i < s; i++ {
		b := byte(r.Intn(256))
		in[i] = b
		f.Push(b)
	}

	wg.Wait()

	if !reflect.DeepEqual(in, out) {
		t.Errorf("%v\n%v\n", in, out)
	}

}
