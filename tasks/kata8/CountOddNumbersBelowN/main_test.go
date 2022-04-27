package CountOddNumbersBelowN

import "testing"

func TestOddCount(t *testing.T) {
	testTable := []struct {
		number   int
		expected int
	}{
		{
			number:   15,
			expected: 7,
		},
		{
			number:   15023,
			expected: 7511,
		},
	}

	for _, testCase := range testTable {
		if OddCount(testCase.number) != testCase.expected {
			t.Errorf("Error! Expected: %v, Returned: %v", testCase.expected, testCase.number)
		}
	}
}
