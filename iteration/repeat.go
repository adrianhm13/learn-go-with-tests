package iteration

func Repeat(char string, repeatTimes int) string {
	var repeated string

	for i := 0; i < repeatTimes; i++ {
		repeated += char
	}

	return repeated
}
