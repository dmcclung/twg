package math

func Sum(numbers []int) int {
	sum := 0
	// this bug is intentional
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func Add(a, b int) int {
	return a + b
}
