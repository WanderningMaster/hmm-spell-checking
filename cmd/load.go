package cmd

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/WanderningMaster/hmm-spell-checking/internal/hmm"
	"github.com/WanderningMaster/hmm-spell-checking/internal/logger"
	"github.com/WanderningMaster/hmm-spell-checking/internal/vocabulary"
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

func getRawVocabulary() []string {
	rHandle, err := os.Open("data/words_alpha.txt")
	defer rHandle.Close()

	utils.Require(err)

	scanner := bufio.NewScanner(rHandle)
	scanner.Split(bufio.ScanLines)

	words := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		words = append(words, line)
	}

	return words
}

func LoadModel(withLogs bool) *hmm.HMM {
	logger := logger.GetLogger()

	pairs := getPairs()
	start := time.Now()
	model, err := hmm.New(hmm.WithCache)
	if err != nil {
		logger.Warn(fmt.Sprintf("Err: %v, skipping...", err))
		model, _ = hmm.New()

		model.Load(pairs)
	}
	logger.Info(
		fmt.Sprintf("Loaded model into memory in: %s", time.Since(start)),
	)

	if withLogs {
		logProbs(model)
	}

	return model
}

func LoadVocabulary() *vocabulary.Vocabulary {
	logger := logger.GetLogger()

	start := time.Now()

	words := getRawVocabulary()
	voc := vocabulary.New()
	voc.Load(words)

	logger.Info(
		fmt.Sprintf("Loaded vocabulary into memory in: %s", time.Since(start)),
	)

	return voc
}

func logProbs(model *hmm.HMM) {
	logger := logger.GetLogger()

	logger.Info("Writing probs to fs...")

	f1, _ := os.Create("cache/transition_probs.txt")
	f2, _ := os.Create("cache/emission_probs.txt")
	f3, _ := os.Create("cache/init_probs.txt")
	defer f1.Close()
	defer f2.Close()
	defer f3.Close()

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

	logger.Info("Finished.")
}
