package singly

import "testing"

func TestLinkedList(t *testing.T) {
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

	t.Run("Insert", func(t *testing.T) {
		list := List{}
		list.Insert(2)
		list.Insert(3)
		list.Insert(4)
		got := list.Values()
		want := []int{4, 3, 2}
		assertArray(t, got, want)
	})

	t.Run("InsertAt", func(t *testing.T) {
		list := List{}
		list.InsertAt(1, 2)
		list.InsertAt(2, 3)
		list.InsertAt(1, 4)
		list.InsertAt(2, 5)
		got := list.Values()
		want := []int{4, 5, 2, 3}
		assertArray(t, got, want)
	})

	t.Run("DeleteAt", func(t *testing.T) {
		list := List{}
		list.BatchInsert([]int{2, 3, 4, 5, 6})
		list.DeleteAt(3)
		list.DeleteAt(1)
		got := list.Values()
		want := []int{5, 3, 2}
		assertArray(t, got, want)
	})

	t.Run("Reverse", func(t *testing.T) {
		list := List{}
		list.BatchInsert([]int{2, 3, 4, 5, 6})
		list.Reverse()
		got := list.Values()
		want := []int{2, 3, 4, 5, 6}
		assertArray(t, got, want)
	})

	t.Run("RecursiveReverse", func(t *testing.T) {
		list := List{}
		list.BatchInsert([]int{2, 3, 4, 5, 6})
		list.RecursiveReverse(list.Head)
		got := list.Values()
		want := []int{2, 3, 4, 5, 6}
		assertArray(t, got, want)
	})
}
