package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/WanderningMaster/hmm-spell-checking/internal/hmm"
	"github.com/WanderningMaster/hmm-spell-checking/utils"
)

func getPairs() []string {
	rHandle, err := os.Open("data/en_keystrokes_pairs_clean.txt")
	defer rHandle.Close()

	utils.Require(err)

	scanner := bufio.NewScanner(rHandle)
	scanner.Split(bufio.ScanLines)

	pairs := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		pairs = append(pairs, line)
	}

	return pairs
}

func loadModel(withLogs bool) *hmm.HMM {
	pairs := getPairs()
	start := time.Now()
	model, err := hmm.New(hmm.WithCache)
	if err != nil {
		fmt.Printf("Err: %v, skipping...\n", err)
		model, _ = hmm.New()

		model.Load(pairs)
	}
	fmt.Println("Loaded model into memory in:", time.Since(start))

	if withLogs {
		logProbs(model)
	}

	return model
}

func logProbs(model *hmm.HMM) {
	fmt.Println("Writing probs to fs...")

	f1, _ := os.Create("cache/transition_probs.txt")
	f2, _ := os.Create("cache/emission_probs.txt")
	f3, _ := os.Create("cache/init_probs.txt")
	err := model.LogProbs(hmm.LogConfig{
		Outs:       f1,
		ProbMatrix: 1,
	})
	utils.Require(err)
	err = model.LogProbs(hmm.LogConfig{
		Outs:       f2,
		ProbMatrix: 2,
	})
	utils.Require(err)
	err = model.LogProbs(hmm.LogConfig{
		Outs:       f3,
		ProbMatrix: 3,
	})
	utils.Require(err)

	fmt.Println("Finished.")
}

func main() {
	loadModel(true)
}
