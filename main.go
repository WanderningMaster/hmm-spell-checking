package main

import (
	"fmt"
	"strings"

	"github.com/WanderningMaster/hmm-spell-checking/cmd"
	"github.com/WanderningMaster/hmm-spell-checking/internal/viterbi"
)

func getStates() []rune {
	alphabet := []rune{}

	for r := 'a'; r <= 'z'; r++ {
		alphabet = append(alphabet, r)
	}

	return alphabet
}

func getRunes(input string) [][]rune {
	dataStr := strings.Split(input, " ")
	dataRunes := [][]rune{}
	for _, str := range dataStr {
		dataRunes = append(dataRunes, []rune(strings.ToLower(str)))
	}

	return dataRunes
}

func main() {
	hmm := cmd.LoadModel(true)
	voc := cmd.LoadVocabulary()

	input := "wrddinh"

	fmt.Printf("Observed: %s\n", input)
	candidates := viterbi.ViterbiNBest([]rune(input), hmm, 10)

	best := candidates[0]
	exists, _ := voc.WordExists(string(best))
	fmt.Printf("Best: %s, Real word: %v\n", string(best), exists)
	fmt.Printf("\nOther Candidates:\n")
	for _, c := range candidates[1:] {
		exists, _ = voc.WordExists(string(c))
		fmt.Printf("%s, Real word: %v\n", string(c), exists)
	}
}
