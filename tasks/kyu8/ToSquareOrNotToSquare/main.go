// https://www.codewars.com/kata/57f6ad55cca6e045d2000627

package ToSquareOrNotToSquare

import "math"

func SquareOrSquareRoot(arr []int) []int {
	outputArray := make([]int, len(arr))

	for i, v := range arr {
		sqrt := math.Sqrt(float64(v))

		if float64(int(sqrt)) == sqrt {
			outputArray[i] = int(sqrt)
		} else {
			outputArray[i] = v * v
		}
	}
	return outputArray
}
