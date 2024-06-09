package services

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"
	"unicode"

	"github.com/WanderningMaster/hmm-spell-checking/internal/hmm"
	"github.com/WanderningMaster/hmm-spell-checking/internal/logger"
	"github.com/WanderningMaster/hmm-spell-checking/internal/viterbi"
	"github.com/WanderningMaster/hmm-spell-checking/internal/vocabulary"
	"github.com/WanderningMaster/hmm-spell-checking/utils"
)

type SpellChecker struct {
	maxVariants int
	hmm         *hmm.HMM
	voc         *vocabulary.Vocabulary
	lambda      float64
}
type Candidate struct {
	Valid    bool     `json:"valid"`
	Best     string   `json:"best"`
	Typo     string   `json:"typo"`
	Variants []string `json:"variants"`
}

func (s *SpellChecker) SetMaxVariants(val int) {
	s.maxVariants = val
}

func getPairs() []string {
	rHandle, err := os.Open("data/training_set.txt")
	utils.Require(err)
	defer rHandle.Close()

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
	rHandle, err := os.Open("data/words_clean.txt")
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

func loadModel(withLogs bool, lambda float64) *hmm.HMM {
	logger := logger.GetLogger()

	pairs := getPairs()
	start := time.Now()
	model, err := hmm.New(hmm.WithCache)
	if err != nil {
		logger.Warn(fmt.Sprintf("Err: %v, skipping...", err))
		model, _ = hmm.New()

		model.Load(pairs, lambda)
	}
	logger.Info(
		fmt.Sprintf("Loaded model into memory in: %s", time.Since(start)),
	)

	if withLogs {
		logProbs(model)
	}

	return model
}

func loadVocabulary() *vocabulary.Vocabulary {
	logger := logger.GetLogger()

	logger.Info("Loading vocabulary into memory")

	start := time.Now()

	words := getRawVocabulary()
	voc := vocabulary.New()
	voc.Load(words)

	logger.Info(
		fmt.Sprintf("Finished in: %s", time.Since(start)),
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

func NewSpellChecker(maxVariant int, lambda float64) *SpellChecker {
	hmm := loadModel(true, lambda)
	voc := loadVocabulary()

	return &SpellChecker{
		hmm:         hmm,
		voc:         voc,
		maxVariants: maxVariant,
		lambda:      lambda,
	}
}

func (s *SpellChecker) CorrectAssert(
	word string,
	actual string,
	counter chan struct{},
	sem chan struct{},
	wg *sync.WaitGroup,
) {
	sem <- struct{}{}
	defer wg.Done()
	candidates := viterbi.ViterbiKBest(
		[]rune(word),
		s.hmm,
		s.maxVariants,
	)

	for _, c := range candidates {
		if string(c) == actual {
			counter <- struct{}{}
			break
		}
	}

	<-sem
}

func (s *SpellChecker) CorrectAssertSync(
	word string,
	actual string,
) bool {
	candidates := viterbi.ViterbiKBest(
		[]rune(word),
		s.hmm,
		s.maxVariants,
	)

	for _, c := range candidates {
		if string(c) == actual {
			return true
		}
	}
	return false
}

func (s *SpellChecker) Correct(
	word string,
	dataChan chan Candidate,
	wg *sync.WaitGroup,
	sem chan struct{},
	r []bool,
) {
	sem <- struct{}{}
	defer wg.Done()
	candidates := viterbi.ViterbiKBest(
		[]rune(word),
		s.hmm,
		s.maxVariants,
	)

	var res Candidate
	res.Typo = word
	best := candidates[0]
	exists, _ := s.voc.WordExists(string(best))
	if exists {
		res.Best = string(best)
	}

	for _, c := range candidates[1:] {
		exists, _ = s.voc.WordExists(string(c))
		if exists && res.Best == "" {
			res.Best = string(c)
		} else if exists {
			res.Variants = append(res.Variants, string(c))
		}
	}
	res.Best = applyRegistr(res.Best, r)
	for idx := range res.Variants {
		res.Variants[idx] = applyRegistr(res.Variants[idx], r)
	}

	dataChan <- res
	<-sem
}

func sanitizeInput(text string) string {
	text = strings.ReplaceAll(text, ",", "")
	text = strings.ReplaceAll(text, ".", "")
	text = strings.ReplaceAll(text, ";", "")
	text = strings.ReplaceAll(text, "!", "")
	text = strings.ReplaceAll(text, "?", "")

	return text
}

func tokenize(text string) []string {
	wordRegex := regexp.MustCompile(`\b\w+('\w+)?\b`)
	matches := wordRegex.FindAllString(text, -1)
	return matches
}

func SaveRegistr(word string) []bool {
	r := []bool{}
	for _, ch := range word {
		if unicode.IsUpper(ch) {
			r = append(r, true)
		} else {
			r = append(r, false)
		}
	}

	return r
}

func applyRegistr(word string, r []bool) string {
	w := []rune{}
	for idx, ch := range word {
		if r[idx] {
			w = append(w, unicode.ToUpper(ch))
		} else {
			w = append(w, ch)
		}
	}

	return string(w)
}

func (s *SpellChecker) CorrectText(text string) ([]Candidate, int, error) {
	text = sanitizeInput(text)
	words := tokenize(text)

	wg := sync.WaitGroup{}
	dataChan := make(chan Candidate, len(words))
	maxConcurrency := 200
	sem := make(chan struct{}, maxConcurrency)

	candidates := []Candidate{}
	totalErrors := 0
	for _, word := range words {
		if word == " " || word == "" {
			continue
		}
		r := SaveRegistr(word)
		word = strings.ToLower(word)

		exists, err := s.voc.WordExists(word)
		if err != nil {
			return nil, 0, err
		}
		if exists {
			candidate := Candidate{
				Valid: true,
				Best:  word,
			}
			candidates = append(candidates, candidate)
			continue
		}
		wg.Add(1)
		go s.Correct(word, dataChan, &wg, sem, r)
		totalErrors += 1
	}
	wg.Wait()
	close(dataChan)

	for c := range dataChan {
		candidates = append(candidates, c)
	}

	return candidates, totalErrors, nil
}
