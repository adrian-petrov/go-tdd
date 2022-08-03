package arrays

import (
	"reflect"
	"testing"
)

func TestSumAll(t *testing.T) {
	t.Run("collection of one array", func(t *testing.T) {
		collection := []int{1, 2, 3}

		got := SumAll(collection)
		expected := []int{6}

		if len(got) != len(expected) {
			t.Errorf("expected a length of %d but got %d", expected, got)
		}

		// method 1
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %v but got %v", expected, got)
		}

		// method 2
		// for i, v := range got {
		// 	if v != expected[i] {
		// 		t.Errorf("expected %v but got %v", expected[i], v)
		// 	}
		// }
	})

	t.Run("collection of multiple arrays", func(t *testing.T) {
		collection1 := []int{1, 2, 3}
		collection2 := []int{4, 5, 6}
		collection3 := []int{22, 33, 99}

		got := SumAll(collection1, collection2, collection3)
		expected := []int{6, 15, 154}

		if len(got) != len(expected) {
			t.Errorf("expected a length of %d but got %d", expected, got)
		}

		// method 1
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %v but got %v", expected, got)
		}

		// method 2
		// for i, v := range got {
		// 	if v != expected[i] {
		// 		t.Errorf("expected %v but got %v", expected[i], v)
		// 	}
		// }
	})
}
