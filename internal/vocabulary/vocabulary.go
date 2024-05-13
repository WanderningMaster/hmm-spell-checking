package vocabulary

import (
	"errors"
	"github.com/WanderningMaster/hmm-spell-checking/utils"
)

type Vocabulary struct {
	trie  *utils.Trie
	ready bool
}

var (
	VocabularyIsNotReadyYet = errors.New("vocabulary is not ready yet")
)

func New() *Vocabulary {
	trie := utils.NewTrie()
	voc := &Vocabulary{
		trie: trie,
	}

	return voc
}

/*
	 NOTE: cant use gob encoding because trie may have nil elements
		in children array which causes crushes.
		In terms of server startup its not such big of a deal but
		for cli implementation can be annoying
*/
func (v *Vocabulary) Load(data []string) {
	// runs ~100ms-200ms
	for _, w := range data {
		v.trie.Insert(w)
	}

	v.ready = true
}

func (v *Vocabulary) WordExists(word string) (bool, error) {
	if !v.ready {
		return false, VocabularyIsNotReadyYet
	}

	return v.trie.Search(word)
}
