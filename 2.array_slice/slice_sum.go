package sum

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	lengthOfNumbers := len(numbersToSum)
	sums = make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = SliceSum(numbers)
	}
	return sums
}
func SliceSum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// 计算面积
