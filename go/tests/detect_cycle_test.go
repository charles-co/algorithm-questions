package tests

import (
	"testing"

	"github.com/charles-co/algorithm_questions/randoms"
)

func TestDetectCycle(t *testing.T) {
	tt := []struct {
		name     string
		head     *randoms.ListNode
		expected *randoms.ListNode
	}{
		{name: "Test 2", head: &randoms.ListNode{Val: 1, Next: &randoms.ListNode{Val: 2, Next: &randoms.ListNode{Val: 3, Next: &randoms.ListNode{Val: 4, Next: &randoms.ListNode{Val: 5, Next: nil}}}}}, expected: nil},
		{name: "Test 3", head: &randoms.ListNode{Val: 1, Next: &randoms.ListNode{Val: 2, Next: nil}}, expected: nil},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if got := randoms.DetectCycle(tc.head); got != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, got)
			}
		})
	}
}
