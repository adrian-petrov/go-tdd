package arrays

func SumAll(params ...[]int) []int {
	var result []int

	for _, arr := range params {
		result = append(result, Sum(arr))
	}

	return result
}
