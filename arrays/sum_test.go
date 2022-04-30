package arrays

import "testing"

func TestSum(t *testing.T) {

	t.Run("collection of any size", func(t *testing.T) {

		var numbers := []int{1, 2, 3}

		got := SumArray(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}

	})
}
