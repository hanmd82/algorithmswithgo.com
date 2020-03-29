package module01solutions

import "math"

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//   BaseToDec("E", 16) => 14
//   BaseToDec("1110", 2) => 14
//
func BaseToDec(value string, base int) int {
	const charset = "0123456789ABCDEF"
	var res int
	value = Reverse(value)

	for i,v := range(value) {
		res += IndexOf(charset, v) * PowInt(base, i)
	}
	return res
}

func PowInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func IndexOf(word string, c rune) int {
	for i,v := range(word) {
		if c == v {
			return i
		}
	}
	panic ("something went wrong")
}

// // simple example for BinaryToDec
// func BinaryToDec(value string) int {
// 	var res int
// 	value = Reverse(value)

// 	for i,v := range(value) {
// 		factor := int(v-'0')
// 		res += factor * PowInt(2, i)
// 	}
// 	return res
// }
