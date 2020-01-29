package linkedlistqueue

import (
	"testing"
)

func TestQueue(t *testing.T) {
	assertTrue := func(t *testing.T, got interface{}, want interface{}) {
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("Enqueue", func(t *testing.T) {
		queue := Queue{}
		queue.Enqueue(1)
		queue.Enqueue(2)
		queue.Enqueue(3)

		got := queue.Values()
		want := [3]int{1, 2, 3}

		for i, value := range want {
			assertTrue(t, value, got[i])
		}
	})

	t.Run("Dequeue", func(t *testing.T) {
		queue := Queue{}
		queue.Enqueue(1)
		queue.Enqueue(2)
		queue.Enqueue(3)
		want := [3]int{1, 2, 3}

		for _, wantValue := range want {
			got := queue.Peek()
			assertTrue(t, got, wantValue)
			queue.Dequeue()
		}
	})

	t.Run("Peek", func(t *testing.T) {
		queue := Queue{}
		queue.Enqueue(1)
		queue.Enqueue(2)
		queue.Enqueue(3)

		got := queue.Peek()
		want := 1
		assertTrue(t, got, want)
	})
}
