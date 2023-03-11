package randoms

// Companies
// The array-form of an integer num is an array representing its digits in left to right order.

// For example, for num = 1321, the array form is [1,3,2,1].
// Given num, the array-form of an integer, and an integer k, return the array-form of the integer num + k.

 

// Example 1:

// Input: num = [1,2,0,0], k = 34
// Output: [1,2,3,4]
// Explanation: 1200 + 34 = 1234
// Example 2:

// Input: num = [2,7,4], k = 181
// Output: [4,5,5]
// Explanation: 274 + 181 = 455
// Example 3:

// Input: num = [2,1,5], k = 806
// Output: [1,0,2,1]
// Explanation: 215 + 806 = 1021

func AddToArrayForm(A []int, K int) []int {
	for i := len(A) - 1; i > -1; i-- {
		A[i] += K % 10
		K /= 10
		if A[i] > 9 {
			A[i] -= 10
			K++
		}
	}
	res := A[:]

	for K > 0 {
		tmp := make([]int, len(res)+1)
		copy(tmp[1:], res)
		tmp[0] = K % 10
		K /= 10
		res = tmp
	}

	return res
}
