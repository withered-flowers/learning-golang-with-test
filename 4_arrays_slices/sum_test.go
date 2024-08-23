package array_slices

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	// t.Run("Sum collection of 5 numbers", func(t *testing.T) {
	// 	numbers := []int{1, 2, 3, 4, 5}

	// 	got := Sum(numbers)
	// 	want := 15

	// 	if got != want {
	// 		t.Errorf("got %d want %d given %v", got, want, numbers)
	// 	}
	// })

	t.Run("Sum collection of any sizes", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 3}, []int{4, 5})
	want := []int{4, 9}

	if !slices.Equal(got, want) {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !slices.Equal(got, want) {
			t.Errorf("got %d want %d", got, want)
		}
	}

	t.Run("Make sum of the slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 3}, []int{0, 9})
		want := []int{3, 9}

		checkSums(t, got, want)
	})

	t.Run("Safely sum the slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSums(t, got, want)
	})
}
