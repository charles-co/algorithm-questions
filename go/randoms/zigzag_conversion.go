// The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

// P   A   H   N
// A P L S I I G
// Y   I   R
// And then read line by line: "PAHNAPLSIIGYIR"

// Write the code that will take a string and make this conversion given a number of rows:

// string convert(string s, int numRows);

package randoms

func Convert(s string, numRows int) string {

	n := len(s)
	answer := ""

	diff := 2*numRows - 2

	for i := 0; i < numRows; i++ {
		index := i

		for index < n {
			answer += string(s[index])
			if i != 0 && i != numRows-1 {
				diagonal_diff := diff - 2*i
				if index+diagonal_diff < n {
					answer += string(s[index+diagonal_diff])
				}
			}
			index += diff
		}
	}

	return string(answer)

}
