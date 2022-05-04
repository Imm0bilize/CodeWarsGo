package TwoToOne

import (
	"reflect"
	"testing"
)

func TestTwoToOne(t *testing.T) {
	testTable := []struct {
		s1       string
		s2       string
		expected string
	}{
		{
			s1:       "aretheyhere",
			s2:       "yestheyarehere",
			expected: "aehrsty",
		},
		{
			s1:       "loopingisfunbutdangerous",
			s2:       "lessdangerousthancoding",
			expected: "abcdefghilnoprstu",
		},
	}

	for _, testCase := range testTable {
		output := TwoToOne(testCase.s1, testCase.s2)
		if !reflect.DeepEqual(output, testCase.expected) {
			t.Errorf("Expected: %v, got: %v", testCase.expected, output)
		}
	}

}
