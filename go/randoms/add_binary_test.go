package randoms_test

// Given two binary strings a and b, return their sum as a binary string.

import (
	"testing"
	"fmt"
	"strconv"
)

func addBinary(a string, b string) string {
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
			result = strconv.Itoa(sum % 2) + result
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

func TestAddBinary(t *testing.T){
	tt := []struct{
		name string
		a string
		b string
		expected string
	}{
		{name: "Test 1", a: "11", b: "1", expected: "100"},
		{name: "Test 2", a: "1010", b: "1011", expected: "10101"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T){
			result := addBinary(tc.a, tc.b)
			if result != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
		})
	}
}