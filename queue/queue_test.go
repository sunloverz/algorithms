package queue

import (
	"testing"
)

func TestArrayList(t *testing.T) {
	assertTrue := func(t *testing.T, got interface{}, want interface{}) {
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("Enqueue", func(t *testing.T){
		queue := Queue{}
		queue.Enqueue("a")
		queue.Enqueue("b")
		queue.Enqueue("c")

		got := queue.Values()
		want := [3]string{"a", "b", "c"}
		
		for i, value := range want {
			assertTrue(t, value, got[i])
		}
		assertTrue(t, queue.Size(), 3)
	})

	t.Run("Dequeu", func(t *testing.T) {
		queue := Queue{}
		queue.Enqueue("a")
		queue.Enqueue("b")
		queue.Enqueue("c")
		want := [3]string{"a", "b", "c"}

		for _, wantValue := range want {
			assertTrue(t, queue.Dequeue(), wantValue)
		}

		got := queue.Dequeue()
		assertTrue(t, got, nil)
	})

	t.Run("Peek", func(t *testing.T) {
		queue := Queue{}
		queue.Enqueue("a")
		queue.Enqueue("b")
		queue.Enqueue("c")

		got := queue.Peek()
		want := "a"
		assertTrue(t, got, want)
	})
}