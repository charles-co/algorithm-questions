package randoms

// Given two binary strings a and b, return their sum as a binary string.

import (
	"fmt"
	"strconv"
)

func AddBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}
	b = fmt.Sprintf("%0*s", len(a), b)
	carry := 0
	result := ""
	for i := len(a) - 1; i > -1; i-- {
		sum := carry
		sum += int(a[i]-'0') + int(b[i]-'0')
		if sum > 1 {
			carry = 1
			result = strconv.Itoa(sum%2) + result
		} else {
			carry = 0
			result = strconv.Itoa(sum) + result
		}

	}
	if carry == 1 {
		return "1" + string(result)
	}
	return string(result)
}
