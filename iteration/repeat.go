package iteration

const repeatTimes = 5

func Repeat(char string) string {
	var repeated string

	for i := 0; i < repeatTimes; i++ {
		repeated += char
	}

	return repeated
}
