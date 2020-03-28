package module01solutions

// DecToBase will return a string representing
// the provided decimal number in the provided base.
// This is limited to bases 2-16 for simplicity.
//
// Eg:
//
//   DecToBase(14, 16) => "E"
//   DecToBase(14, 2) => "1110"
//
func DecToBase(dec, base int) string {
	const charset = "0123456789ABCDEF"
	res := ""
	for dec > 0 {
		rem := dec%base
		res = string(charset[rem]) + res
		dec /= base
	}
	return res
}

// // simple example for DecToBinary
// func DecToBinary(dec int) string {
// 	res := ""
// 	for {
// 		if dec == 0 {
// 			break
// 		} else {
// 			q, r := dec/2, dec%2
// 			res = strconv.Itoa(r) + res
// 			dec = q
// 			fmt.Println("quotient:", q, ", remainder:", r, ", res:", res)
// 		}
// 	}
// 	return res
// }
