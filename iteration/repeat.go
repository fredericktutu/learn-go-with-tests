package iteration

//given a character and a number, return a string that repeat the character by the number
func Repeat(character string, times int) string {
	repeated := ""
	for i := 0; i < times; i++ {
		repeated += character
	}
	return repeated
}
