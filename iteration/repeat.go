package iteration

const nRep = 5

func Repeat(s string) string {
	var repStr string
	for i := 0; i < nRep; i++ {
		repStr += "a"
	}
	return repStr
}
