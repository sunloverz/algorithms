package doublylinkedlist

import "testing"

func TestDoublyLinkedList(t *testing.T) {
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
	t.Run("InsertAtHead", func(t *testing.T) {
		list := List{}
		list.InsertAtHead(1)
		list.InsertAtHead(2)
		list.InsertAtHead(3)
		got := list.Values()
		want := []int{3, 2, 1}
		assertArray(t, got, want)
	})
}
