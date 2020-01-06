package stack

import (
	"testing"
)

func TestArrayList(t *testing.T) {
	assertTrue := func(t *testing.T, got interface{}, want interface{}) {
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("Push", func(t *testing.T) {
		stack := Stack{}
		stack.Push(1)

		got := stack.Size()
		want := 1
		assertTrue(t, got, want)
	})

	t.Run("Pop", func(t *testing.T) {
		stack := Stack{}
		stack.Push("Hello")
		stack.Push("World")

		got := stack.Pop()
		want := "World"
		assertTrue(t, got, want)
	})

	t.Run("Empty", func(t *testing.T) {
		stack := Stack{}
		got := stack.Empty()
		want := true
		assertTrue(t, got, want)
	})

	t.Run("Peek", func(t *testing.T) {
		stack := Stack{}
		stack.Push("Hello")
		stack.Push("World")

		got := stack.Peek()
		want := "World"
		assertTrue(t, got, want)
	})

	t.Run("Search", func(t *testing.T) {
		stack := Stack{}
		stack.Push("one")
		stack.Push("two")
		stack.Push("three")
		
		got := stack.Search("three")
		want := 2
		assertTrue(t, got, want)

		got = stack.Search("no")
		want = -1
		assertTrue(t, got, want)
	})

}