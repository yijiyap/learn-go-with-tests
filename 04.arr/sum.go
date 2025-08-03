package arr

func Sum(numbers []int) int {
	var ans int
	for _, n := range numbers {
		ans += n
	}
	return ans
}

func SumAll(numbersToSum ...[]int) []int {
	ans := make([]int, len(numbersToSum))
	for i, numbers := range numbersToSum {
		ans[i] = Sum(numbers)
	}
	return ans
}

func SumAllTails(numbersToSum ...[]int) []int {
	ans := make([]int, len(numbersToSum))

	for i, numbers := range numbersToSum {
		if len(numbers) == 0 {
			ans[i] = 0
		} else {
			ans[i] = Sum(numbers[1:])
		}
	}
	return ans
}
