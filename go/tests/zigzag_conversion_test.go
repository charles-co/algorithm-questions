// The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

// P   A   H   N
// A P L S I I G
// Y   I   R
// And then read line by line: "PAHNAPLSIIGYIR"

// Write the code that will take a string and make this conversion given a number of rows:

// string convert(string s, int numRows);

package tests

import (
	"github.com/charles-co/algorithm_questions/randoms"
	"testing"
)

func TestZigZagConversion(t *testing.T) {
	tt := []struct {
		name     string
		input    string
		numRows  int
		expected string
	}{
		{name: "Test 1", input: "PAYPALISHIRING", numRows: 3, expected: "PAHNAPLSIIGYIR"},
		{name: "Test 2", input: "PAYPALISHIRING", numRows: 4, expected: "PINALSIGYAHRPI"},
		{name: "Test 3", input: "A", numRows: 1, expected: "A"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			result := randoms.Convert(tc.input, tc.numRows)
			if result != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
		})
	}
}
