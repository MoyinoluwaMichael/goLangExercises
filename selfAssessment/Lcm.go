package selfAssessment

import (
	"fmt"
	"math"
)

func findLcm(numbers []int) int {
	var result int = 1
	var maxIndex int = math.MinInt64
	for i := 0; i < len(numbers); i++ {
		if numbers[i] > maxIndex {
			maxIndex = numbers[i]
		}
	}
	for divisor := 2; divisor < maxIndex; divisor++ {
		var count int = 0
		for i := 0; i < len(numbers); i++ {
			if numbers[i]%divisor == 0 {
				numbers[i] /= divisor
				count++
			}
		}
		if count > 0 {
			result *= divisor
			divisor--
		}
	}
	return result
}

func main() {
	ints := []int{3, 5, 4, 10}
	fmt.Println(findLcm(ints))
}
