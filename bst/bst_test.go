package bst

import "testing"

func TestBinarySearchTree(t *testing.T) {
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
		tree := BinarySearchTree{}
		want := []int{15, 10, 20, 6, 12, 25}
		for _, val := range want {
			tree.Insert(val)
		}
		got := tree.BFS()
		assertArray(t, got, want)
	})

	t.Run("Max", func(t *testing.T) {
		tree := BinarySearchTree{}
		values := []int{15, 10, 20, 6, 12, 25}
		for _, val := range values {
			tree.Insert(val)
		}
		got := tree.Max()
		want := 25
		assertTrue(t, got, want)
	})

	t.Run("Min", func(t *testing.T) {
		tree := BinarySearchTree{}
		values := []int{15, 10, 20, 6, 12}
		for _, val := range values {
			tree.Insert(val)
		}
		got := tree.Min()
		want := 6
		assertTrue(t, got, want)
	})
}
