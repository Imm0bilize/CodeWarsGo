// https://www.codewars.com/kata/57a1fd2ce298a731b20006a4/

package IsItAPalindrome

import (
	"strings"
)

func IsPalindrome(str string) bool {
	str = strings.ToLower(str)

	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}
	return true
}
