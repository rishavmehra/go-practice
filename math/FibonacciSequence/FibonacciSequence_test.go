package fibonaccisequence

import (
	"reflect"
	"testing"
)

func TestFibonacciSequence(t *testing.T) {
	t.Run("Result for 9 Fibonacci recursion", func(t *testing.T) {
		got := FibonacciRecur(9)
		want := 34

		if got != want {
			t.Errorf("Result is wrong, its giving: %x", got)
		}

	})

	t.Run("Result for 9 Fibonacci withtout recursion", func(t *testing.T) {
		got := Fibonaccisequence(9)
		want := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Result is wrong, its giving: %x", got)
		}

	})
}
