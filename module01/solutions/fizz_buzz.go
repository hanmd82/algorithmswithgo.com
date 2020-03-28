package module01solutions

import (
	"fmt"
	"strconv"
)

// FizzBuzz will print out all of the numbers
// from 1 to N replacing any divisible by 3
// with "Fizz", and divisible by 5 with "Buzz",
// and any divisible by both with "Fizz Buzz".
//
// Note: The test for this is a little
// complicated so that you can just use the
// `fmt` package and print to standard out.
// I wouldn't normally recommend this, but did
// it here to make life easier for beginners.
func FizzBuzz(n int) {
	results := ""
	for i := 1; i <= n; i++ {
		if (i%3 == 0) {
			results += "Fizz"
			if (i%5 == 0) {
				results += " Buzz"
			}
		} else if (i%5 == 0) {
			results += "Buzz"
		} else {
			results += strconv.Itoa(i)
		}
		if i != n {
			results += ", "
		}
	}
	fmt.Println(results)
}
