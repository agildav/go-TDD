package arraysslices

func sum(numbers []int) int {
	sum := 0

	for _, n := range numbers {
		sum += n
	}
	return sum
}

func sumAll(numbersList ...[]int) []int {
	nSlices := len(numbersList)
	sums := make([]int, nSlices)

	for i, numbers := range numbersList {
		sums[i] = sum(numbers)
	}

	return sums
}
