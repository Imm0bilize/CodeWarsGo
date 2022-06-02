package FindTheUniqueNumber

import "testing"

func TestFindUniq(t *testing.T) {
	testTable := []struct {
		input    []float32
		expected float32
	}{
		{
			input:    []float32{1, 1, 1, 2, 1, 1},
			expected: 2,
		},
		{
			input:    []float32{0, 0, 0.55, 0, 0},
			expected: 0.55,
		},
		{
			input:    []float32{0, 0, 0, 0, 0, 0, 0, 0, 9, 0, 0, 0},
			expected: 9,
		},
	}

	for _, testCase := range testTable {
		result := FindUniq(testCase.input)
		if result != testCase.expected {
			t.Errorf("Expected %v, got: %v", testCase.expected, result)
		}
	}
}
