package IsItAPalindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	testTable := []struct {
		inputLine string
		expected  bool
	}{
		{
			inputLine: "a",
			expected:  true,
		},
		{
			inputLine: "aba",
			expected:  true,
		},
		{
			inputLine: "Abba",
			expected:  true,
		},
		{
			inputLine: "hello",
			expected:  false,
		},
	}

	for _, testCase := range testTable {
		result := IsPalindrome(testCase.inputLine)
		if result != testCase.expected {
			t.Errorf("Expected %v, returned %v", testCase.expected, result)
		}
	}
}
