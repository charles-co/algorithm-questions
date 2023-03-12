package tests

import (
	"github.com/charles-co/algorithm_questions/randoms"
	"testing"
)

func TestMergeKLists(t *testing.T) {
	tt := []struct {
		name     string
		lists    []*randoms.ListNode
		expected *randoms.ListNode
	}{
		{
			name: "Example 1",
			lists: []*randoms.ListNode{
				{Val: 1, Next: &randoms.ListNode{Val: 4, Next: &randoms.ListNode{Val: 5}}},
				{Val: 1, Next: &randoms.ListNode{Val: 3, Next: &randoms.ListNode{Val: 4}}},
				{Val: 2, Next: &randoms.ListNode{Val: 6}},
			},
			expected: &randoms.ListNode{Val: 1, Next: &randoms.ListNode{Val: 1, Next: &randoms.ListNode{Val: 2, Next: &randoms.ListNode{Val: 3, Next: &randoms.ListNode{Val: 4, Next: &randoms.ListNode{Val: 4, Next: &randoms.ListNode{Val: 5, Next: &randoms.ListNode{Val: 6}}}}}}}},
		},
		{
			name:     "Example 2",
			lists:    []*randoms.ListNode{},
			expected: nil,
		},
		{
			name: "Example 3",
			lists: []*randoms.ListNode{
				nil,
			},
			expected: nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := randoms.MergeKList(tc.lists)
			if actual == nil && tc.expected == nil {
				return
			}
			if actual == nil || tc.expected == nil {
				t.Errorf("Expected %v, got %v", tc.expected, actual)
			}
			for actual != nil && tc.expected != nil {
				if actual.Val != tc.expected.Val {
					t.Errorf("Expected %v, got %v", tc.expected, actual)
				}
				actual = actual.Next
				tc.expected = tc.expected.Next
			}
			if actual != nil || tc.expected != nil {
				t.Errorf("Expected %v, got %v", tc.expected, actual)
			}
		})
	}
}
