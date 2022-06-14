package hw04lrucache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushBack(25)      // [25]
		l.Remove(l.Front()) // []
		l.PushFront(35)     // [35]
		l.Remove(l.Back())  // []
		l.PushFront(10)     // [10]
		l.PushFront(15)     // [15, 10]
		l.PushFront(20)     // [20, 15, 10]
		l.PushBack(20)      // [20, 15, 10, 20]
		l.PushBack(25)      // [20, 15, 10, 20, 25]
		l.PushBack(30)      // [20, 15, 10, 20, 25, 30]
		require.Equal(t, 6, l.Len())
		l.Remove(l.Front()) // [15, 10, 20, 25, 30]
		l.Remove(l.Back())  // [15, 10, 20, 25]
		require.Equal(t, 4, l.Len())

		l.PushFront(35)          // [35, 15, 10, 20, 25]
		l.Remove(l.Front().Next) // [35, 10, 20, 25]
		l.Remove(l.Back().Prev)  // [35, 10, 25]
		require.Equal(t, 3, l.Len())

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 35, 10, 25, 50, 70]

		require.Equal(t, 8, l.Len())
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		l.MoveToFront(l.Front()) // [80, 60, 40, 35, 10, 25, 50, 70]
		l.MoveToFront(l.Back())  // [70, 80, 60, 40, 35, 10, 25, 50]

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{70, 80, 60, 40, 35, 10, 25, 50}, elems)
	})
}
