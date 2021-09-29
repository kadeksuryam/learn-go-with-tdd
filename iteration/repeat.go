package iteration

func Repeat(s string) string {
	var repStr string
	for i := 0; i < 5; i++ {
		repStr = repStr + "a"
	}
	return repStr
}
