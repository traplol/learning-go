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
