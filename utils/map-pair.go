package utils

import (
	"fmt"
	"strconv"
	"strings"
)

type Tuple struct {
	Observed rune
	State    rune
}

func (t Tuple) String() string {
	return fmt.Sprintf("(%s, %s) ", strconv.QuoteRune(t.Observed), strconv.QuoteRune(t.State))
}

func bigger(strs []string) int {
	if len(strs[0]) > len(strs[1]) {
		return len(strs[0])
	}
	return len(strs[1])
}

func MapWordPair(line string) ([]Tuple, error) {
	words := strings.Fields(line)
	records := []Tuple{}
	if len(words) > 2 {
		return nil, fmt.Errorf("invalid line format")
	}
	typo, word := words[0], words[1]

	ptr := 0
	maxLen := bigger(words)
	for ; ptr < maxLen; ptr += 1 {
		chWord := ' '
		chTypo := ' '
		if ptr < len(typo) {
			chTypo = rune(typo[ptr])
		}
		if ptr < len(word) {
			chWord = rune(word[ptr])
		}
		records = append(records, Tuple{
			Observed: chTypo,
			State:    chWord,
		})
	}

	return records, nil
}
