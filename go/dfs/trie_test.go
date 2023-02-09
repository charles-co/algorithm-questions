package dfs_test

import (
	"testing"
)

type TrieNode struct {
	children map[rune]*TrieNode
	isWord   bool
}

type Trie struct {
	root *TrieNode
}

func Constructor() *Trie {
	return &Trie{root: &TrieNode{children: make(map[rune]*TrieNode), isWord: false}}
}

func (t *Trie) Insert(word string) {
	node := t.root
	for _, r := range word {
		if _, ok := node.children[r]; !ok {
			node.children[r] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		node = node.children[r]
	}
	node.isWord = true
}

func (t *Trie) Search(word string, charCounts map[rune]int) bool {
	node := t.root
	for _, r := range word {
		if _, ok := node.children[r]; !ok {
			return false
		}
		if charCounts[r] == 0 {
			return false
		}
		charCounts[r]--
		node = node.children[r]
	}
	return node.isWord
}

func longestMatchingWord(dictionary []string, chars string) string {
	trie := Constructor()
	charCounts := make(map[rune]int)
	for _, r := range chars {
		charCounts[r]++
	}

	for _, word := range dictionary {
		trie.Insert(word)
	}

	result := ""
	for _, word := range dictionary {
		count := make(map[rune]int)
		for k, v := range charCounts {
			count[k] = v
		}
		if trie.Search(word, count) && len(word) > len(result) {
			result = word
		}
	}
	return result
}

func TestLongestMatchingWord(t *testing.T) {
	tt := []struct {
		name       string
		dictionary []string
		chars      string
		result     string
	}{
		{"Test 1", []string{"when", "what", "whatthen", "whatnow"}, "whatno", "what"},
		{"Test 2", []string{"when", "what", "whatthen", "whatnow"}, "whatnwo", "whatnow"},
		{"Test 3", []string{"when", "what", "whatthen", "whatnow"}, "wonthaw", "whatnow"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			result := longestMatchingWord(tc.dictionary, tc.chars)
			if result != tc.result {
				t.Errorf("Expected %v, got %v", tc.result, result)
			}
		})
	}

}
