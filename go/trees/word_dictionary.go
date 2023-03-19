package trees

type Node struct {
	children []*Node
	isWord   bool
}

func WordDictionary() *Node {
	return &Node{children: make([]*Node, 26)}
}

func (n *Node) AddWord(word string) {
	for _, char := range word {
		if n.children[char-'a'] == nil {
			n.children[char-'a'] = WordDictionary()
		}
		n = n.children[char-'a']
	}
	n.isWord = true
}

func (n *Node) Search(word string) bool {

	var res bool

	var dfs func(node *Node, word string)
	dfs = func(node *Node, word string) {

		if len(word) == 0 {
			if node.isWord {
				res = true
			}
			return
		}

		if word[0] == '.' {
			for _, child := range node.children {
				if child != nil {
					dfs(child, word[1:])
				}
			}
		} else {
			if node.children[word[0]-'a'] != nil {
				dfs(node.children[word[0]-'a'], word[1:])
			}
		}

	}
	dfs(n, word)
	return
}
