package module01solutions

import (
	"fmt"
	"strconv"
)

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
	res := ""
	// fmt.Print(dec, " to base", base, " = ")
	for {
		if dec == 0 {
			break
		} else {
			q, r := dec/base, dec%base
			str := strconv.Itoa(r)
			if r >= 10 {
				str = DecToHex(r)
			}
			res = str + res
			dec = q
		}
	}
	// fmt.Println(res)
	return res
}

func DecToHex(dec int) string {
	switch dec {
	case 10:
		return "A"
	case 11:
		return "B"
	case 12:
		return "C"
	case 13:
		return "D"
	case 14:
		return "E"
	case 15:
		return "F"
	default:
		return ""
	}
}

// simple example for DecToBinary
func DecToBinary(dec int) string {
	res := ""
	for {
		if dec == 0 {
			break
		} else {
			q, r := dec/2, dec%2
			res = strconv.Itoa(r) + res
			dec = q
			fmt.Println("quotient:", q, ", remainder:", r, ", res:", res)
		}
	}
	return res
}
