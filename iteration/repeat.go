package iteration

func Repeat(x string) (repeated string) {
	for i := 0; i < 5; i++ {
		repeated += x
	}
	return
}
