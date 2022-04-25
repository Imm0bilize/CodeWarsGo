package ToSquareOrNotToSquare

import (
	"reflect"
	"testing"
)

func TestSquareOrSquareRoot(t *testing.T) {
	testTable := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{4, 3, 9, 7, 2, 1},
			expected: []int{2, 9, 3, 49, 4, 1},
		},
		{
			input:    []int{100, 101, 5, 5, 1, 1},
			expected: []int{10, 10201, 25, 25, 1, 1},
		},
		{
			input:    []int{1, 2, 3, 4, 5, 6},
			expected: []int{1, 4, 9, 2, 25, 36},
		},
	}

	for i, testCase := range testTable {
		if reflect.DeepEqual(SquareOrSquareRoot(testCase.input), testCase.expected) {
			t.Logf("Test: %v OK", i)
		} else {
			t.Errorf("Test: %v Failed! Expected: %+v Returned: %+v", i, testCase.expected, testCase.input)
		}
	}

}
