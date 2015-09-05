package stack_test

import (
	"bfi/stack"
	"testing"
)

func TestStackPush(t *testing.T) {
	stack := stack.NewStack()
	stack.Push(42)
	if stack.Count() != 1 {
		t.Error("expected stack count to be: 1")
	}
	stack.Push(2)
	if stack.Count() != 2 {
		t.Error("expected stack count to be: 2")
	}
}

func TestStackPop(t *testing.T) {
	stack := stack.NewStack()
	stack.Push(42)
	stack.Push(2)

	var v int
	var empty bool

	v, empty = stack.Pop()
	if e := 2; v != e {
		t.Errorf("expected stack to pop the value: %d, got %d", e, v)
	}
	if empty != false {
		t.Error("expected stack to not be empty.")
	}
	if e := 1; stack.Count() != e {
		t.Errorf("expected stack count to be: %d, got %d", e, v)
	}

	v, empty = stack.Pop()
	if e := 42; v != e {
		t.Errorf("expected stack to pop the value: %d, got %d", e, v)
	}
	if empty != true {
		t.Error("expected stack to be empty.")
	}
	if e := 0; stack.Count() != e {
		t.Errorf("expected stack count to be: %d, got %d", e, v)
	}

	stack.Push(100)
	v, empty = stack.Pop()
	if e := 100; v != e {
		t.Errorf("expected stack to pop the value: %d, got %d", e, v)
	}
	if empty != true {
		t.Error("expected stack to be empty.")
	}
	if e := 0; stack.Count() != e {
		t.Errorf("expected stack count to be: %d, got %d", e, v)
	}

	v, empty = stack.Pop()
	if e := 0; v != e {
		t.Errorf("expected stack to pop the value: %d, got %d", e, v)
	}
	if empty != true {
		t.Error("expected stack to be empty.")
	}
	if e := 0; stack.Count() != e {
		t.Errorf("expected stack count to be: %d, got %d", e, v)
	}
}

func TestManyPushes(t *testing.T) {
	count := 10000
	stack := stack.NewStack()
	for i := 0; i < count; i++ {
		stack.Push(i + 1)
	}
	if a := stack.Count(); a != count {
		t.Errorf("Expected stack count to be %d, got %d", count, a)
	}
	for e := count; e > 0; e-- {
		if a, _ := stack.Pop(); a != e {
			t.Errorf("Expected stack to pop %d, got %d", e, a)
		}
	}
}
