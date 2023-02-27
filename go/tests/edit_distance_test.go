package tests

import (
	"testing"
	"github.com/charles-co/algorithm_questions/randoms"

)


func TestEditDistance(t *testing.T){
	tt := []struct{
		word1 string
		word2 string
		expected int
	}{
		{"horse", "ros", 3},
		{"intention", "execution", 5},
	}

	for _, tc := range tt{
		t.Run(tc.word1+tc.word2, func(t *testing.T){
			if got := randoms.EditDistance(tc.word1, tc.word2); got != tc.expected{
				t.Errorf("expected %d, got %d", tc.expected, got)
			}
		})
	}
}

func TestEditDistanceWithRecursion(t *testing.T){
	tt := []struct{
		word1 string
		word2 string
		expected int
	}{
		{"horse", "ros", 3},
		{"intention", "execution", 5},
	}

	for _, tc := range tt{
		t.Run(tc.word1+tc.word2, func(t *testing.T){
			if got := randoms.EditDistanceWithRecursion(tc.word1, tc.word2, 0, 0, make(map[string]int)); got != tc.expected{
				t.Errorf("expected %d, got %d", tc.expected, got)
			}
		})
	}
}