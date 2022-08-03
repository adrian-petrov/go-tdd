package iteration

func Repeat(character string, multiplier int) string {
	var repeated string

	for i := 0; i < multiplier; i++ {
		repeated += character
	}

	return repeated
}
