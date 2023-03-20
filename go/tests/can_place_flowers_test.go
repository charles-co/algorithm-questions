package tests

import (
	"testing"
	"github.com/charles-co/algorithm_questions/randoms"
)

func TestCanPlaceFlowers(t *testing.T){

	tt := []struct{
		name string
		flowerbed []int
		n int
		expected bool
	}{
		{
			name: "Test 1",
			flowerbed: []int{1,0,0,0,1},
			n: 1,
			expected: true,
		},
		{
			name: "Test 2",
			flowerbed: []int{1,0,0,0,1},
			n: 2,
			expected: false,
		},
		{
			name: "Test 3",
			flowerbed: []int{0,0,1,0,1},
			n: 1,
			expected: true,	
		},
		{
			name: "Test 4",
			flowerbed: []int{1,0,0,0,1,0,0},
			n: 2,
			expected: true,
		},
		{
			name: "Test 5",
			flowerbed: []int{1,0,0,0,1,0,0},
			n: 3,
			expected: false,
		},
		{
			name: "Test 6",
			flowerbed: []int{0,0,1,0,1},
			n: 2,
			expected: false,
		},
		{
			name: "Test 7",
			flowerbed: []int{0,0,0,0,0,1,0,0},
			n: 2,
			expected: true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T){
			res := randoms.CanPlaceFlowers(tc.flowerbed, tc.n)
			if res != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, res)
			}
		})
	}

}