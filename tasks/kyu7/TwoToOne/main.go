// https://www.codewars.com/kata/5656b6906de340bd1b0000ac/

package TwoToOne

import (
	"sort"
	"strings"
)

func TwoToOne(s1 string, s2 string) string {
	charMap := make(map[string]bool)

	for _, v := range s1 + s2 {
		convertedValue := string(v)
		if _, ok := charMap[convertedValue]; !ok {
			charMap[convertedValue] = true
		}
	}

	keys := make([]string, 0, len(charMap))
	for k := range charMap {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	return strings.Join(keys, "")
}
