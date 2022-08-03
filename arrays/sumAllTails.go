package arrays

func SumAllTails(params ...[]int) []int {
	var result []int

	for _, arr := range params {
		var tail []int

		if len(arr) == 0 {
			tail = []int{0}
		} else {
			tail = arr[1:]
		}
		result = append(result, Sum(tail))
	}

	return result
}
