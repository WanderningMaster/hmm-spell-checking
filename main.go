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

func testData() [][]rune {
	dataStr := []string{
		"abd",
		"hrlp",
		"wrdding",
		"vook",
		"nark",
		"traon",
		"alpja",
		"pjone",
		"jite",
	}

	dataRunes := [][]rune{}
	for _, str := range dataStr {
		dataRunes = append(dataRunes, []rune(str))
	}

	return dataRunes
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
