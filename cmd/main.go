package main

import (
	"log"
	"os"

	"github.com/WanderningMaster/hmm-spell-checking/internal/hmm"
)

func main() {

	pairs := []string{
		"actoss	across",
		"actualll	actually",
		"actuallu	actually",
		"acuarium	aquarium",
		"cay	cat",
	}
	model := hmm.New()
	model.Load(pairs)

	f, err := os.Create("cache/transition_probs.txt")
	if err != nil {
		log.Fatal(err)
	}
	model.LogProbs(hmm.LogConfig{
		Outs:       f,
		ProbMatrix: 1,
	})
}
