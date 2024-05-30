package cmd

import (
	"fmt"

	"github.com/WanderningMaster/hmm-spell-checking/services"
	"github.com/WanderningMaster/hmm-spell-checking/utils"
)

func RunCli() {
	spellChecker := services.NewSpellChecker(30)
	input := "Thry're"

	candidates, _, err := spellChecker.CorrectText(input)
	utils.Require(err)

	fmt.Println("Best: ", candidates[0].Best)
	for _, c := range candidates[0].Variants {
		fmt.Println(c)
	}
}
