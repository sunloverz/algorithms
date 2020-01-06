package arraylist

import (
	"reflect"
	"testing"
)

func TestArrayList(t *testing.T) {

	assertTrue := func(t *testing.T, got interface{}, want interface{}) {
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("Add", func(t *testing.T) {
		arraylist := ArrayList{}
		arraylist.Add(1)

		got := arraylist.Size()
		want := 1
		assertTrue(t, got, want)
	})

	t.Run("Get", func(t *testing.T) {
		arrayList := ArrayList{}
		arrayList.Add(2)

		got := arrayList.Get(0)
		want := 2
		assertTrue(t, got, want)
	})

	t.Run("Set", func(t *testing.T) {
		list := ArrayList{}
		list.Add(2)
		list.Set(0, 3)

		got := list.Get(0)
		want := 3
		assertTrue(t, got, want)
	})

	t.Run("Remove", func(t *testing.T) {
		list := ArrayList{}
		list.Add(1,2,3,4,5)
		list.Remove(2)

		got := list.Values()
		want := [4]int{1, 2, 4, 5}
		if reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
