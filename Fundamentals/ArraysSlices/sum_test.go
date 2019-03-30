package arraysslices

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	checkSum := func(t *testing.T, got, want int) {
		t.Helper()

		// Assert the message
		if got != want {
			errMsg := fmt.Sprintf("got -> %d, want -> %d", got, want)
			// Print the error with colors
			t.Errorf(colorError(errMsg))
		}
	}

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := sum(numbers)
		want := 15

		checkSum(t, got, want)
	})

	t.Run("collection with negative numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 5, -4, 5, 9, -1, -2, -18, -2, -3, 5, 1}
		got := sum(numbers)
		want := 1

		checkSum(t, got, want)
	})
}

func TestSumAll(t *testing.T) {

	checkSum := func(t *testing.T, got []int, want []int) {
		t.Helper()

		// Assert the message
		if !reflect.DeepEqual(got, want) {
			errMsg := fmt.Sprintf("got -> %v, want -> %v", got, want)
			// Print the error with colors
			t.Errorf(colorError(errMsg))
		}
	}

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		numbers2 := []int{2, 4, 6, 8, 10}
		numbers3 := []int{3, 5, 7, 9, 11}

		got := sumAll(numbers, numbers2, numbers3)
		want := []int{15, 30, 35}

		checkSum(t, got, want)
	})
}
