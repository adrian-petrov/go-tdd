package arrays

import (
	"reflect"
	"testing"
)

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, expected []int) {
		t.Helper()

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %v but got %v", expected, got)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{1, 9})
		expected := []int{2, 9}
		checkSums(t, got, expected)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		expected := []int{0, 9}
		checkSums(t, got, expected)
	})

}
