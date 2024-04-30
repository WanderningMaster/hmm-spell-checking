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

	exists, _ := voc.WordExists("zwitterionic")
	fmt.Printf("Word %s is %v\n", "zwitterionic", exists)

	input := "tirture"

	fmt.Printf("Observed: %s\n", input)
	candidates := viterbi.ViterbiNBest([]rune(input), hmm, 10)

	best := candidates[0]
	fmt.Printf("Best: %s\n", string(best))
	fmt.Printf("\nOther Candidates:\n")
	for _, c := range candidates[1:] {
		fmt.Printf("%s\n", string(c))
	}
}
