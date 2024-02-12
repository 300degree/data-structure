package main

import "log"

type Stack struct {
	stack []int
}

func (s *Stack) Push(n int) {
	if len(s.stack) < 5 {
		s.stack = append(s.stack, n)
	} else {
		log.Fatal("Stack Overflow")
	}
}

func (s *Stack) Pop() int {
	n := 0
	if len(s.stack) > 0 {
		n = s.stack[len(s.stack)-1]
		s.stack = s.stack[:len(s.stack)-1]
	} else {
		log.Fatal("Stack Underflow")
	}
	return n
}

func (s *Stack) Peek() int {
	return s.stack[len(s.stack)-1]
}

func (s *Stack) Size() int {
	return len(s.stack)
}

func (s *Stack) Empty() bool {
	return len(s.stack) == 0
}
