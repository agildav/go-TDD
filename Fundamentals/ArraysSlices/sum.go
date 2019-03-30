package arraysslices

func sum(numbers []int) int {
	sum := 0

	for _, n := range numbers {
		sum += n
	}
	return sum
}
