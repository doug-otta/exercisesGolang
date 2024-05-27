package sum

func Sum(numbers []int) int {
	add := 0
	for _, num := range numbers {
		add += num
	}
	return add
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumOthers(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbersToSum) == 0 {
			sums = append(sums, 0)
		} else {
			final := numbers[1:]
			sums = append(sums, Sum(final))
		}
	}

	return sums
}
