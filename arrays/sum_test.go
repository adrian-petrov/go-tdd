package arrays

import "testing"

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		expected := 15
		assertCorrectMessage(t, got, expected, numbers)
	})
}

func assertCorrectMessage(t testing.TB, got int, expected int, numbers []int) {
	t.Helper()
	if got != expected {
		t.Errorf("expected %d but got %d given, %v", expected, got, numbers)
	}
}
