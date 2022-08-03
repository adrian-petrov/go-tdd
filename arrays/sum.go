package arrays

func Sum(numbers []int) int {
	var result int

	// basic for loop
	// for i := 0; i < len(numbers); i++ {
	// 	result += numbers[i]
	// }
	for _, number := range numbers {
		result += number
	}

	return result
}
