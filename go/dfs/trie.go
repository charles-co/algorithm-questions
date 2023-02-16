package dfs

type TrieNode struct {
	children map[rune]*TrieNode
	isWord   bool
}

type Trie struct {
	root *TrieNode
}

func constructor() *Trie {
	return &Trie{root: &TrieNode{children: make(map[rune]*TrieNode), isWord: false}}
}

func (t *Trie) insert(word string) {
	node := t.root
	for _, r := range word {
		if _, ok := node.children[r]; !ok {
			node.children[r] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		node = node.children[r]
	}
	node.isWord = true
}

func (t *Trie) search(word string, charCounts map[rune]int) bool {
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

func LongestMatchingWord(dictionary []string, chars string) string {
	trie := constructor()
	charCounts := make(map[rune]int)
	for _, r := range chars {
		charCounts[r]++
	}

	for _, word := range dictionary {
		trie.insert(word)
	}

	result := ""
	for _, word := range dictionary {
		count := make(map[rune]int)
		for k, v := range charCounts {
			count[k] = v
		}
		if trie.search(word, count) && len(word) > len(result) {
			result = word
		}
	}
	return result
}
