package randoms_test

import (
	"testing"
)
func addToArrayForm(A []int, K int) []int {
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


func TestAddToArrayForm(t *testing.T) {
	tt := []struct {
		name     string
		A        []int
		K        int
		expected []int
	}{
		{"test1", []int{1, 2, 0, 0}, 34, []int{1, 2, 3, 4}},
		{"test2", []int{2, 7, 4}, 181, []int{4, 5, 5}},
		{"test3", []int{2, 1, 5}, 806, []int{1, 0, 2, 1}},
		{"test4", []int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9}, 1, []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			result := addToArrayForm(tc.A, tc.K)
			if len(result) != len(tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
			for i := range result {
				if result[i] != tc.expected[i] {
					t.Errorf("expected %v, got %v", tc.expected, result)
				}
			}
		})
	}
}