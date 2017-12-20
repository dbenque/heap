package main

import (
	"fmt"
	"sync"
	"time"
)

//Fifo with block push and pop. Not an interesting struct, this is what channel do (and they do it better)
type Fifo struct {
	sync.Mutex
	index, count int
	buf          []byte
}

func NewFifo(capacity int) (*Fifo, error) {
	if capacity == 0 {
		return nil, fmt.Errorf("No capacity allocated")
	}
	return &Fifo{index: 0, count: 0, buf: make([]byte, capacity)}, nil
}

func (s *Fifo) Empty() bool {
	return s.count == 0
}

func (s *Fifo) Pop() byte {
	s.Lock()
	if s.count == 0 {
		s.Unlock()
		for {
			time.Sleep(time.Millisecond)
			s.Lock()
			if s.count > 0 {
				break
			}
			s.Unlock()
		}
	}
	defer s.Unlock()
	b := s.buf[s.index]
	s.index = (s.index + 1) % len(s.buf)
	s.count--
	return b
}

func (s *Fifo) Push(v byte) {
	s.Lock()
	if s.count+1 > len(s.buf) {
		s.Unlock()
		for {
			time.Sleep(time.Millisecond)
			s.Lock()
			if s.count+1 <= len(s.buf) {
				break
			}
			s.Unlock()
		}
	}
	defer s.Unlock()
	s.buf[(s.index+s.count)%len(s.buf)] = v
	s.count++
}
