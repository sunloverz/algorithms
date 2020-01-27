package linkedliststack

import (
	"testing"
)

func TestStack(t *testing.T) {
	assertTrue := func(t *testing.T, got interface{}, want interface{}) {
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	assertArray := func(t *testing.T, got, want []int) {
		for i, val := range got {
			assertTrue(t, val, want[i])
		}
	}

	t.Run("Push", func(t *testing.T) {
		stack := Stack{}
		stack.Push(1)
		stack.Push(2)
		stack.Push(3)

		got := stack.Values()
		want := []int{3, 2, 1}
		assertArray(t, got, want)
	})

	t.Run("Pop", func(t *testing.T) {
		stack := Stack{}
		stack.Push(1)
		stack.Push(2)
		stack.Push(3)

		got := stack.Pop()
		want := 3
		assertTrue(t, got, want)
		assertArray(t, stack.Values(), []int{2, 1})
	})

	t.Run("Peek", func(t *testing.T) {
		stack := Stack{}
		stack.Push(2)
		stack.Push(1)

		got := stack.Peek()
		want := 1
		assertTrue(t, got, want)
	})

	t.Run("IsEmpty", func(t *testing.T) {
		stack := Stack{}
		stack.Push(1)
		stack.Pop()
		got := stack.IsEmpty()
		want := true
		assertTrue(t, got, want)
	})
}
