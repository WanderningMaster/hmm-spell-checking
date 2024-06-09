package cmd

import (
	"fmt"

	"github.com/WanderningMaster/hmm-spell-checking/services"
	"github.com/WanderningMaster/hmm-spell-checking/utils"
)

func RunCli() {
	lambda := 0.035
	spellChecker := services.NewSpellChecker(30, lambda)
	input := "U've rjally dnjoyed the aovie we wxtched last noght"

	candidates, _, err := spellChecker.CorrectText(input)
	utils.Require(err)

	fmt.Println(input)
	for _, c := range candidates {
		fmt.Println(c)
	}
}
