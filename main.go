package main

import (
	"fmt"
	"github.com/WanderningMaster/hmm-spell-checking/services"
)

func main() {
	spellChecker := services.NewSpellChecker(10)

	input := "wrddinh"
	candidate, _ := spellChecker.Correct(input)

	fmt.Printf("Observed: %s\n", input)

	fmt.Printf("Best: %s\n", string(candidate.Best))
	fmt.Printf("\nOther Candidates:\n")
	for _, c := range candidate.Variants {
		fmt.Printf("%s\n", string(c))
	}
}
