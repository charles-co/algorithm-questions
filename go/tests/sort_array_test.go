package tests

import (
	"github.com/charles-co/algorithm_questions/randoms"
	"testing"
)

func TestSortArray(t *testing.T) {
	tt := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Example 1",
			input:    []int{5, 2, 3, 1},
			expected: []int{1, 2, 3, 5},
		},
		{
			name:     "Example 2",
			input:    []int{5, 1, 1, 2, 0, 0},
			expected: []int{0, 0, 1, 1, 2, 5},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := randoms.BubbleSort(tc.input)
			if len(actual) != len(tc.expected) {
				t.Errorf("Expected %d, got %d", len(tc.expected), len(actual))
			}
			for i := range actual {
				if actual[i] != tc.expected[i] {
					t.Errorf("Expected %d, got %d", tc.expected[i], actual[i])
				}
			}

		})
		t.Run(tc.name, func(t *testing.T) {
			actual := randoms.InsertionSort(tc.input)
			if len(actual) != len(tc.expected) {
				t.Errorf("Expected %d, got %d", len(tc.expected), len(actual))
			}
			for i := range actual {
				if actual[i] != tc.expected[i] {
					t.Errorf("Expected %d, got %d", tc.expected[i], actual[i])
				}
			}
		})
		t.Run(tc.name, func(t *testing.T) {
			actual := randoms.SelectionSort(tc.input)
			if len(actual) != len(tc.expected) {
				t.Errorf("Expected %d, got %d", len(tc.expected), len(actual))
			}
			for i := range actual {
				if actual[i] != tc.expected[i] {
					t.Errorf("Expected %d, got %d", tc.expected[i], actual[i])
				}
			}
		})
		t.Run(tc.name, func(t *testing.T) {
			actual := randoms.MergeSort(tc.input)
			if len(actual) != len(tc.expected) {
				t.Errorf("Expected %d, got %d", len(tc.expected), len(actual))
			}
			for i := range actual {
				if actual[i] != tc.expected[i] {
					t.Errorf("Expected %d, got %d", tc.expected[i], actual[i])
				}
			}
		})
		t.Run(tc.name, func(t *testing.T) {
			actual := randoms.QuickSort(tc.input)
			if len(actual) != len(tc.expected) {
				t.Errorf("Expected %d, got %d", len(tc.expected), len(actual))
			}
			for i := range actual {
				if actual[i] != tc.expected[i] {
					t.Errorf("Expected %d, got %d", tc.expected[i], actual[i])
				}
			}
		})
	}

}
