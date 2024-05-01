package utils

// size of bytes to cover english alphabet
const MAX_CHILDREN = 26

type node struct {
	children [MAX_CHILDREN]*node
	terminal bool
}

type Trie struct {
	root *node
}

func NewTrie() *Trie {
	root := &node{}
	return &Trie{
		root: root,
	}
}

func (t *Trie) Insert(w string) {
	curr := t.root
	for idx := range w {
		chIdx := w[idx] - 'a'
		if curr.children[chIdx] == nil {
			curr.children[chIdx] = &node{}
		}
		curr = curr.children[chIdx]
	}

	curr.terminal = true
}

func (t *Trie) Search(w string) bool {
	curr := t.root
	for idx := range w {
		chIdx := w[idx] - 'a'
		if curr.children[chIdx] == nil {
			return false
		}
		curr = curr.children[chIdx]
	}
	if curr.terminal {
		return true
	}

	return false
}
