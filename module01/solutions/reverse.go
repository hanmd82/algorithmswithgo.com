package module01solutions

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
func Reverse(word string) string {
	len := len(word)
	result := make([]rune, len)

	for i,v := range(word) {
		result[len-i-1] = v
	}
	return string(result)
}
