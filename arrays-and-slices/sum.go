package arrs

func Sum(numbers []int) int {
	var sum int

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var totals []int

	for _, numbers := range numbersToSum {
		totals = append(totals, Sum(numbers))
	}

	return totals
}

func SumAllTails(tailsToSum ...[]int) (tails []int) {

	for _, numbers := range tailsToSum {
		sliceTail := numbers[1:]
		tails = append(tails, Sum(sliceTail))
	}

	return tails
}
