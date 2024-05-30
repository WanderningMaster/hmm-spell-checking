package utils

import "fmt"

// size of bytes to cover english alphabet + extra '\” symbol (pos 27)
const MAX_CHILDREN = 28

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
		var chIdx byte
		if w[idx] == '\'' {
			chIdx = 27
		} else {
			chIdx = w[idx] - 'a'
		}
		if curr.children[chIdx] == nil {
			curr.children[chIdx] = &node{}
		}
		curr = curr.children[chIdx]
	}

	curr.terminal = true
}

func (t *Trie) Search(w string) (bool, error) {
	curr := t.root
	for idx := range w {
		chIdx := w[idx] - 'a'
		if chIdx >= MAX_CHILDREN {
			if w[idx] == '\'' {
				chIdx = 27
			} else {
				return false, fmt.Errorf("unknown character")
			}
		}
		if curr.children[chIdx] == nil {
			return false, nil
		}
		curr = curr.children[chIdx]
	}
	if curr.terminal {
		return true, nil
	}

	return false, nil
}
