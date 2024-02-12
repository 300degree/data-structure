package main

import "testing"

func Test_Push(t *testing.T) {
	s := Stack{}

	for i := 0; i < 5; i++ {
		s.Push(i)
	}

	if size := s.Size(); size != 5 {
		t.Errorf("Expected stack size to be 5, but got %d", size)
	}
}

func Test_Pop(t *testing.T) {
	s := Stack{}

	s.Push(1)
	s.Push(2)

	popped := s.Pop()
	if popped != 2 {
		t.Errorf("Expected popped value to be 2, but got %d", popped)
	}

	if size := s.Size(); size != 1 {
		t.Errorf("Expected stack size to be 1 after pop, but got %d", size)
	}
}

func Test_Peek(t *testing.T) {
	s := Stack{}

	s.Push(10)

	if peeked := s.Peek(); peeked != 10 {
		t.Errorf("Expected peeked value to be 10, but got %d", peeked)
	}

	if size := s.Size(); size != 1 {
		t.Errorf("Expected stack size to remain 1 after peek, but got %d", size)
	}
}

func Test_Size(t *testing.T) {
}

func Test_Empty(t *testing.T) {
	s := Stack{}

	if empty := s.Empty(); !empty {
		t.Error("Expected stack to be empty")
	}

	s.Push(5)

	if empty := s.Empty(); empty {
		t.Error("Expected stack to not be empty after pushing element")
	}
}
