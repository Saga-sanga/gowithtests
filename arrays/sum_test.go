package main

import "testing"

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 3, 4, 6, 2}

		got := Sum(numbers)
		want := 16

		assertErrorMessage(t, got, want, numbers)
	})

}

func assertErrorMessage(t testing.TB, got, want int, numbers []int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}
